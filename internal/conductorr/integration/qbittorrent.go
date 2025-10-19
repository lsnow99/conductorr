package integration

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/lsnow99/conductorr/pkg/constant"
	"github.com/rs/zerolog/log"
	"github.com/zeebo/bencode"
)

type QBittorrent struct {
	client  *http.Client
	baseUrl url.URL
}

// Non exhaustive struct
type QBittorrentTorrent struct {
	AmountLeft  int    `json:"amount_left"`
	TotalSize   int    `json:"total_size"`
	State       string `json:"state"`
	Hash        string `json:"hash"`
	Name        string `json:"name"`
	ContentPath string `json:"content_path"`
}

func NewQBittorrent(username, password, baseUrl string) (*QBittorrent, error) {
	q := new(QBittorrent)
	var err error
	baseUrlPtr, err := url.Parse(baseUrl)
	if err != nil {
		return nil, fmt.Errorf("bad url %s", baseUrl)
	}
	q.baseUrl = *baseUrlPtr
	q.client = &http.Client{Timeout: time.Duration(6) * time.Second}

	vals := url.Values{
		"username": []string{username},
		"password": []string{password},
	}

	loginUrl := q.baseUrl
	loginUrl.Path = "api/v2/auth/login"

	resp, err := q.client.PostForm(loginUrl.String(), vals)

	if err != nil {
		return nil, err
	}

	resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("bad status code %d %s", resp.StatusCode, loginUrl.String())
	}

	cookies := resp.Cookies()

	if len(cookies) == 0 {
		return nil, fmt.Errorf("received no cookies")
	}

	q.client.Jar, err = cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	q.client.Jar.SetCookies(&loginUrl, cookies)

	return q, err
}

func NewQBittorrentFromConfig(configuration map[string]interface{}) (*QBittorrent, error) {
	username, uOK := configuration["username"].(string)
	password, pOK := configuration["password"].(string)
	baseUrl, bOK := configuration["baseUrl"].(string)
	if !uOK || !pOK || !bOK {
		return nil, fmt.Errorf("failed to parse qbittorrent configuration")
	}

	return NewQBittorrent(username, password, baseUrl)
}

type TorrentMetadata struct {
	Info bencode.RawMessage `bencode:"info"`
}

/*
 * Populate the form to queue a release based on the downloadURL.
 * We expect that the downloadURL contains a torrent file or a redirect to a
 * magnet link.
 */
func writeQueueDownloadForm(release Release, formWriter *multipart.Writer) (string, error) {
	nextURL, err := url.Parse(release.DownloadURL)
	if err != nil {
		return "", err
	}

	// override default http redirect follow behavior
	client := &http.Client{}
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	// perform redirect following manually. stop on magnet url.
	var resp *http.Response
	maxFollows := 100
	for i := 0; i < maxFollows; i++ {
		if nextURL.Scheme == "magnet" {
			break
		}

		if nextURL.Scheme != "http" && nextURL.Scheme != "https" {
			return "", fmt.Errorf("Unknown protocol '%s' while following redirect", nextURL.Scheme)
		}

		resp, err = client.Get(nextURL.String())
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()
		if resp.StatusCode != 301 {
			break
		}

		nextURL, err = url.Parse(resp.Header.Get("Location"))
		if err != nil {
			return "", err
		}
	}

	if nextURL.Scheme == "magnet" {
		fieldWriter, err := formWriter.CreateFormField("urls")
		if err != nil {
			return "", err
		}
		_, err = fieldWriter.Write([]byte(nextURL.String()))
		if err != nil {
			return "", err
		}

		exactTopic := nextURL.Query().Get("xt")
		parts := strings.Split(exactTopic, ":")
		if len(parts) != 3 || parts[0] != "urn" || parts[1] != "btih" {
			return "", fmt.Errorf("Unexpected xt format '%s'", exactTopic)
		}
		btih := parts[2]
		return btih, nil
	} else {
		tm := TorrentMetadata{}
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}

		// Followed http(s) to completion
		err = bencode.NewDecoder(bytes.NewBuffer(data)).Decode(&tm)
		if err != nil {
			return "", err
		}
		fieldWriter, err := formWriter.CreateFormFile("torrents", fmt.Sprintf("%s.torrent", release.Title))
		fieldWriter.Write(data)

		h := sha1.Sum(tm.Info)
		hashStr := hex.EncodeToString(h[:])
		return hashStr, nil
	}
}

func (q *QBittorrent) AddRelease(release Release) (string, error) {
	addUrl := q.baseUrl
	addUrl.Path = "api/v2/torrents/add"

	var buf bytes.Buffer
	formWriter := multipart.NewWriter(&buf)
	defer formWriter.Close()
	hashStr, err := writeQueueDownloadForm(release, formWriter)
	if err != nil {
		return "", nil
	}

	req, err := http.NewRequest(http.MethodPost, addUrl.String(), &buf)

	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", formWriter.FormDataContentType())

	resp, err := q.client.Do(req)
	if err != nil {
		return "", err
	}

	resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("non 200 status code %d", resp.StatusCode)
	}

	return hashStr, err
}

func (q *QBittorrent) DeleteDownload(identifier string) error {
	deleteUrl := q.baseUrl
	deleteUrl.Path = "api/v2/torrents/delete"

	vals := url.Values{
		"hashes":      []string{identifier},
		"deleteFiles": []string{strconv.FormatBool(false)},
	}

	resp, err := q.client.PostForm(deleteUrl.String(), vals)
	if err != nil {
		return err
	}
	resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("qbittorrent could not delete download with id %s", identifier)
	}
	return nil
}

func (q *QBittorrent) PollDownloads(identifiers []string) ([]Download, error) {
	listUrl := q.baseUrl
	listUrl.Path = "api/v2/torrents/info"

	vals := url.Values{
		"hashes": []string{strings.Join(identifiers, "|")},
	}

	listUrl.RawQuery = vals.Encode()

	resp, err := q.client.Get(listUrl.String())
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	torrents := make([]QBittorrentTorrent, 0, len(identifiers))

	if err := json.NewDecoder(resp.Body).Decode(&torrents); err != nil {
		return nil, err
	}

	log.Debug().Msg(fmt.Sprintf("got torrents %v", torrents))

	downloads := make([]Download, 0, len(torrents))

	for _, torrent := range torrents {
		d := Download{}

		if torrent.State == "error" || torrent.State == "missingFiles" || torrent.State == "unknown" {
			d.Status = constant.StatusError
		}

		if torrent.State == "stalledDL" || torrent.State == "queuedDL" || torrent.State == "allocating" || torrent.State == "metaDL" || torrent.State == "checkingResumeData" {
			d.Status = constant.StatusWaiting
		}

		if torrent.State == "moving" || torrent.State == "checkingUP" {
			d.Status = constant.StatusProcessing
		}

		if torrent.State == "forcedUP" || torrent.State == "stalledUP" || torrent.State == "pausedUP" || torrent.State == "queuedUP" || torrent.State == "uploading" {
			d.Status = constant.StatusComplete
		}

		if torrent.State == "pausedDL" {
			d.Status = constant.StatusPaused
		}

		if torrent.State == "downloading" {
			d.Status = constant.StatusDownloading
		}

		if d.Status == constant.StatusComplete {
			d.FinalDir = torrent.ContentPath
		}

		if d.Status == "" {
			return nil, fmt.Errorf("unknown torrent state '%s'", torrent.State)
		}

		d.BytesLeft = uint64(torrent.AmountLeft)
		d.FullSize = uint64(torrent.TotalSize)
		d.FriendlyName = torrent.Name
		d.Identifier = torrent.Hash
		downloads = append(downloads, d)
	}

	log.Debug().Msg(fmt.Sprintf("downloads %v", downloads))

	return downloads, nil
}

func (q *QBittorrent) TestConnection() error {
	infoUrl := q.baseUrl
	infoUrl.Path = "api/v2/transfer/info"

	resp, err := q.client.Get(infoUrl.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	if !json.Valid(body) {
		return fmt.Errorf("error parsing info response")
	}

	return nil
}
