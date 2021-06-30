package main

import (
	"github.com/lsnow99/conductorr/integration"
	"github.com/lsnow99/conductorr/logger"
)

type SyncDownloads struct {
	downloaders        []integration.Downloader
	downloadsToMonitor []string
	downloads          []integration.Download
}

func (sd SyncDownloads) DoTask() {
	for _, downloader := range sd.downloaders {
		if err := downloader.PollDownloads(sd.downloadsToMonitor); err != nil {
			logger.LogWarn(err)
		}
		sd.downloads = append(sd.downloads, downloader.GetDownloads()...)
	}
}

func (sd SyncDownloads) GetDownloads() []integration.Download {
	return sd.downloads
}
