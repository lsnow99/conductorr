package integration

import "time"

type Downloader interface {
	AddMedia(*Media) error
	/*
		PollDownloads takes a string of download ID's (typically hashes for torrents or nzb ids
		for nzbget) and performs a refresh for the status of all the downloads identified in the
		string array. A future call to GetDownloads will
	*/
	PollDownloads([]string) error
	GetDownloads() []Download
}

type Download struct {
	// Media embedded struct
	Media
	// Status of the media
	Status int
	// Started time
	Started time.Time
	// BytesLeft number of bytes remaining to download
	BytesLeft uint64
	// FullSize full size of the download in bytes
	FullSize uint64
}
