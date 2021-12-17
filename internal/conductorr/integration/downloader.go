package integration

import (
	"fmt"
	"time"
)

type Downloader interface {
	// AddRelease takes a Release struct and adds it to the downloader, returning a string
	// identifier that can later be used in the argument passed to PollDownloads to specify
	// which downloads to poll for. Returns non-nil error on failure.
	AddRelease(Release) (string, error)
	
	// DeleteDownload takes a string download identifier and deletes it from the downloader.
	// Returns non-nil error on failure.
	DeleteDownload(string) error

	// PollDownloads takes a string of download identifiers which were returned by earlier calls
	// to AddRelease, and performs a refresh for the status of all the downloads identified in the
	// string array. Returns the updated downloads and any error.
	PollDownloads([]string) ([]Download, error)

	// TestConnection checks the connection to the downloader, and returns a non-nil error on failure.
	TestConnection() error
}

type Download struct {
	ID int
	MediaID      int
	FriendlyName string
	Identifier   string
	FinalDir     string
	// Status of the media
	Status string
	// Started time
	Started time.Time
	// BytesLeft number of bytes remaining to download
	BytesLeft uint64
	// FullSize full size of the download in bytes
	FullSize uint64
}

func NewDownloaderFromConfig(downloaderType string, config map[string]interface{}) (Downloader, error) {
	switch downloaderType {
	case "transmission":
		return NewTransmissionFromConfig(config)
	case "nzbget":
		return NewNZBGetFromConfig(config)
	case "qbittorrent":
		return NewQBittorrentFromConfig(config)
	}
	return nil, fmt.Errorf("unrecognized downloader type: %s", downloaderType)
}
