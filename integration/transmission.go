package integration

import (
	"encoding/json"
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

type TransmissionConfig struct {
	Username string
	Password string
	URL      string
}

func NewTransmission(username, password, baseUrl string) *Transmission {
	return &Transmission{
		username: username,
		password: password,
		baseUrl:  baseUrl,
	}
}

func NewTransmissionFromConfig(configuration string) (*Transmission, error) {
	cfg := TransmissionConfig{}
	if err := json.Unmarshal([]byte(configuration), &cfg); err != nil {
		return nil, err
	}

	return NewTransmission(cfg.Username, cfg.Password, cfg.URL), nil
}

func (t *Transmission) Init() error {
	endpoint, _ := url.Parse(t.baseUrl)

	port := endpoint.Port()
	if port == "" && endpoint.Scheme == "http" {
		port = "80"
	} else if port == "" && endpoint.Scheme == "https" {
		port = "443"
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

func (t *Transmission) AddMedia(media *Media) error {
	falseVal := false
	addPayload := transmissionrpc.TorrentAddPayload{
		Filename: &media.URL,
		Paused: &falseVal,
	}
	torrent, err := t.client.TorrentAdd(&addPayload)
	if err != nil {
		return err
	}
	
	if torrent.HashString == nil {
		return errors.New("transmission did not return a hash string")
	}
	media.Identifier = *torrent.HashString

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
			break
		case transmissionrpc.TorrentStatusCheckWait:
			d.Status = Waiting
			break
		case transmissionrpc.TorrentStatusCheck:
			d.Status = Processing
			break
		case transmissionrpc.TorrentStatusDownloadWait:
			d.Status = Waiting
			break
		case transmissionrpc.TorrentStatusDownload:
			d.Status = Downloading
			break
		case transmissionrpc.TorrentStatusSeedWait:
			d.Status = Done
			break
		case transmissionrpc.TorrentStatusSeed:
			d.Status = Done
			break
		case transmissionrpc.TorrentStatusIsolated:
			d.Status = Error
			break
		}
		d.BytesLeft = uint64(*torrent.LeftUntilDone)
		d.FullSize = uint64(*torrent.TotalSize)
		d.FriendlyName = *torrent.Name
		downloads = append(downloads, d)
	}
	return nil
}

func (t *Transmission) GetDownloads() []Download {
	return t.downloads
}