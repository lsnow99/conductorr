package integration

import (
	"errors"
	"net/url"
	"strconv"

	"github.com/hekmon/transmissionrpc"
)

type Transmission struct {
	username  string
	password  string
	baseUrl   string
	client    *transmissionrpc.Client
	downloads []Download
}

func NewTransmission(username, password, baseUrl string) *Transmission {
	return &Transmission{
		username: username,
		password: password,
		baseUrl:  baseUrl,
	}
}

func NewTransmissionFromConfig(configuration map[string]interface{}) (*Transmission, error) {
	username, uOK := configuration["username"].(string)
	password, pOK := configuration["password"].(string)
	baseUrl, bOK := configuration["base_url"].(string)
	if !uOK || !pOK || !bOK {
		return nil, errors.New("failed to parse transmission configuration")
	}

	return NewTransmission(username, password, baseUrl), nil
}

func (t *Transmission) Init() error {
	endpoint, _ := url.Parse(t.baseUrl)

	port := endpoint.Port()
	if port == "" && endpoint.Scheme == "http" {
		port = "80"
	} else if port == "" && endpoint.Scheme == "https" {
		port = "443"
	} else if port == "" {
		return errors.New("invalid URL")
	}

	portInt, err := strconv.ParseUint(port, 10, 16)
	if err != nil {
		return err
	}

	client, err := transmissionrpc.New(endpoint.Hostname(), t.username, t.password, &transmissionrpc.AdvancedConfig{
		HTTPS: endpoint.Scheme == "https",
		Port:  uint16(portInt),
	})

	if err != nil {
		return err
	}
	t.client = client

	return nil
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

func (t *Transmission) AddRelease(release *Release) error {
	if release == nil {
		return errors.New("release passed to transmission was nil")
	}
	falseVal := false
	addPayload := transmissionrpc.TorrentAddPayload{
		Filename: &release.DownloadURL,
		Paused:   &falseVal,
	}
	torrent, err := t.client.TorrentAdd(&addPayload)
	if err != nil {
		return err
	}

	if torrent.HashString == nil {
		return errors.New("transmission did not return a hash string")
	}
	release.Identifier = *torrent.HashString

	return nil
}

func (t *Transmission) PollDownloads(hashes []string) error {
	torrents, err := t.client.TorrentGetAllForHashes(hashes)
	if err != nil {
		return err
	}

	downloads := make([]Download, 0)
	for _, torrent := range torrents {
		d := Download{}

		if torrent.Status == nil {
			// Warning
			continue
		}
		if torrent.LeftUntilDone == nil {
			// Warning
			continue
		}
		if torrent.Name == nil {
			// Warning
			continue
		}

		switch *torrent.Status {
		case transmissionrpc.TorrentStatusStopped:
			d.Status = Paused
		case transmissionrpc.TorrentStatusCheckWait:
			d.Status = Waiting
		case transmissionrpc.TorrentStatusCheck:
			d.Status = Processing
		case transmissionrpc.TorrentStatusDownloadWait:
			d.Status = Waiting
		case transmissionrpc.TorrentStatusDownload:
			d.Status = Downloading
		case transmissionrpc.TorrentStatusSeedWait:
			d.Status = Done
		case transmissionrpc.TorrentStatusSeed:
			d.Status = Done
		case transmissionrpc.TorrentStatusIsolated:
			d.Status = Error
		}
		d.BytesLeft = uint64(*torrent.LeftUntilDone)
		d.FullSize = uint64(*torrent.TotalSize)
		d.FriendlyName = *torrent.Name
		downloads = append(downloads, d)
	}

	t.downloads = downloads
	return nil
}

func (t *Transmission) GetDownloads() []Download {
	return t.downloads
}
