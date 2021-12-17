package integration

import (
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/l3uddz/go-qbittorrent/qbt"
	"github.com/lsnow99/conductorr/pkg/torrentfile"
)

type QBittorrent struct {
	client *qbt.Client
}

func NewQBittorrent(username, password, baseUrl string) (*QBittorrent, error) {
	q := new(QBittorrent)
	q.client = qbt.NewClient(baseUrl)
	err := q.client.Login(qbt.LoginOptions{
		Username: username,
		Password: password,
	})

	return q, err
}

func NewQBittorrentFromConfig(configuration map[string]interface{}) (*QBittorrent, error) {
	username, uOK := configuration["username"].(string)
	password, pOK := configuration["password"].(string)
	baseUrl, bOK := configuration["base_url"].(string)
	if !uOK || !pOK || !bOK {
		return nil, fmt.Errorf("failed to parse transmission configuration")
	}

	return NewQBittorrent(username, password, baseUrl)
}

type TorrentMetadata struct {
	Announce string `bencode:"announce"`
	Info     []byte `bencode:"info"`
}

func (q *QBittorrent) AddRelease(release Release) (string, error) {

	resp, err := http.Get(release.DownloadURL)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	torrent, err := torrentfile.Unmarshal(resp.Body)
	if err != nil {
		return "", err
	}
	hash, err := torrent.Info.Hash()
	hashStr := hex.EncodeToString(hash)
	return hashStr, err
}

func (q *QBittorrent) DeleteDownload(identifier string) error {
	// Delete the torrent and associated files
	ok, err := q.client.Delete([]string{identifier}, true)
	if !ok {
		return fmt.Errorf("qbittorrent could not delete download with id %s", identifier)
	}
	return err
}

func (q *QBittorrent) PollDownloads(identifiers []string) ([]Download, error) {
	return nil, nil
}

func (q *QBittorrent) TestConnection() error {
	_, err := q.client.Info()
	return err
}
