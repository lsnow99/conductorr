package integration

import (
	"encoding/base64"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"path/filepath"
	"strconv"

	"github.com/hekmon/transmissionrpc"
	"github.com/lsnow99/conductorr/pkg/constant"
)

type Transmission struct {
	username  string
	password  string
	baseUrl   string
	client    *transmissionrpc.Client
}

func NewTransmission(username, password, baseUrl string) (*Transmission, error) {
	endpoint, err := url.Parse(baseUrl)
	if err != nil {
		return nil, err
	}

	port := endpoint.Port()
	if port == "" && endpoint.Scheme == "http" {
		port = "80"
	} else if port == "" && endpoint.Scheme == "https" {
		port = "443"
	} else if port == "" {
		return nil, errors.New("invalid URL")
	}

	portInt, err := strconv.ParseUint(port, 10, 16)
	if err != nil {
		return nil, err
	}

	client, err := transmissionrpc.New(endpoint.Hostname(), username, password, &transmissionrpc.AdvancedConfig{
		HTTPS: endpoint.Scheme == "https",
		Port:  uint16(portInt),
	})

	if err != nil {
		return nil, err
	}

	return &Transmission{
		username: username,
		password: password,
		baseUrl:  baseUrl,
		client:   client,
	}, nil
}

func NewTransmissionFromConfig(configuration map[string]interface{}) (*Transmission, error) {
	username, uOK := configuration["username"].(string)
	password, pOK := configuration["password"].(string)
	baseUrl, bOK := configuration["baseUrl"].(string)
	if !uOK || !pOK || !bOK {
		return nil, errors.New("failed to parse transmission configuration")
	}

	return NewTransmission(username, password, baseUrl)
}

func (t *Transmission) TestConnection() error {
	ok, _, _, err := t.client.RPCVersion()
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("transmission failed health check, possibly conductorr or transmission is outdated")
	}
	return nil
}

func (t *Transmission) AddRelease(release Release) (string, error) {
	falseVal := false

	resp, err := http.Get(release.DownloadURL)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	torrentContent := base64.StdEncoding.EncodeToString(body)

	addPayload := transmissionrpc.TorrentAddPayload{
		MetaInfo: &torrentContent,
		Paused:   &falseVal,
	}
	torrent, err := t.client.TorrentAdd(&addPayload)
	if err != nil {
		return "", err
	}

	if torrent.ID == nil {
		return "", errors.New("transmission did not return an id for the torrent")
	}

	return strconv.FormatInt(*torrent.ID, 10), nil
}

func (t *Transmission) DeleteDownload(identifier string) error {
	id, err := strconv.ParseInt(identifier, 10, 64)
	if err != nil {
		return err
	}

	payload := transmissionrpc.TorrentRemovePayload{
		IDs: []int64{id},
		DeleteLocalData: true,
	}
	return t.client.TorrentRemove(&payload)
}

func (t *Transmission) PollDownloads(identifiers []string) ([]Download, error) {
	ids := []int64{}
	for _, identifier := range identifiers {
		id, err := strconv.ParseInt(identifier, 10, 64)
		if err != nil {
			// Bad identifier
			continue
			// return nil, err
		}
		ids = append(ids, id)
	}

	torrents, err := t.client.TorrentGetAllFor(ids)
	if err != nil {
		return nil, err
	}

	downloads := make([]Download, 0)
	for _, torrent := range torrents {
		d := Download{}

		if torrent.Status == nil || torrent.LeftUntilDone == nil || torrent.Name == nil || torrent.HashString == nil ||
			torrent.DownloadDir == nil || torrent.TotalSize == nil {
			// Warning
			continue
		}

		switch *torrent.Status {
		case transmissionrpc.TorrentStatusStopped:
			if torrent.Error != nil && *torrent.Error != 0 {
				d.Status = constant.StatusError
			} else {
				d.Status = constant.StatusPaused
			}
		case transmissionrpc.TorrentStatusCheckWait:
			d.Status = constant.StatusWaiting
		case transmissionrpc.TorrentStatusCheck:
			d.Status = constant.StatusProcessing
		case transmissionrpc.TorrentStatusDownloadWait:
			d.Status = constant.StatusWaiting
		case transmissionrpc.TorrentStatusDownload:
			d.Status = constant.StatusDownloading
		case transmissionrpc.TorrentStatusSeedWait:
			d.Status = constant.StatusComplete
		case transmissionrpc.TorrentStatusSeed:
			d.Status = constant.StatusComplete
		case transmissionrpc.TorrentStatusIsolated:
			d.Status = constant.StatusError
		}

		if d.Status == constant.StatusComplete {
			d.FinalDir = filepath.Join(*torrent.DownloadDir, *torrent.Name)
		}

		d.BytesLeft = uint64(*torrent.LeftUntilDone)
		d.FullSize = uint64(*torrent.TotalSize)/8
		d.FriendlyName = *torrent.Name
		d.Identifier = strconv.FormatInt(*torrent.ID, 10)
		downloads = append(downloads, d)
	}

	return downloads, nil
}
