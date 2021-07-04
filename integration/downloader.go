package integration

import (
	"errors"
	"time"
)

type Downloader interface {
	AddRelease(*Release) error
	/*
		PollDownloads takes a string of download ID's (typically hashes for torrents or nzb ids
		for nzbget) and performs a refresh for the status of all the downloads identified in the
		string array. A future call to GetDownloads will
	*/
	PollDownloads([]string) error
	GetDownloads() []Download
	TestConnection() error
}

type Download struct {
	Media *Media
	FriendlyName string
	Identifier string
	FinalDir string
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
	}
	return nil, errors.New("no such downloader registered")
}