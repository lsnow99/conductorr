package app

import (
	"errors"
	"sync"
	"time"

	"github.com/lsnow99/conductorr/constant"
	"github.com/lsnow99/conductorr/integration"
	"github.com/lsnow99/conductorr/logger"
)

type ManagedDownloader struct {
	ID             int
	Name           string
	DownloaderType string
	integration.Downloader
}

type DownloaderManager struct {
	*sync.RWMutex
	downloaders        []ManagedDownloader
	downloadsToMonitor []string
	downloads          []integration.Download
}

func (md ManagedDownloader) GetName() string {
	return md.Name
}

func (dm *DownloaderManager) DoTask() {
	for _, dlr := range dm.downloaders {
		if err := dlr.PollDownloads(dm.downloadsToMonitor); err != nil {
			logger.LogWarn(err)
		}
		dm.downloads = dlr.GetDownloads()
	}
}

func (dm *DownloaderManager) GetFrequency() time.Duration {
	return time.Second * 30
}

func (dm *DownloaderManager) GetDownloads() []integration.Download {
	return dm.downloads
}

func (dm *DownloaderManager) getDownloaders() []ManagedDownloader {
	return dm.downloaders
}

func (dm *DownloaderManager) RegisterDownloader(id int, downloaderType, name string, configuration map[string]interface{}) error {
	dlr, err := integration.NewDownloaderFromConfig(downloaderType, configuration)
	if err != nil {
		return err
	}
	md := ManagedDownloader{
		Name:           name,
		DownloaderType: downloaderType,
		Downloader:     dlr,
	}
	var added bool
	for i, downloader := range dm.downloaders {
		if downloader.ID == id {
			dm.downloaders[i] = md
			added = true
		}
	}
	if !added {
		dm.downloaders = append(dm.downloaders, md)
	}
	return nil
}

func (dm *DownloaderManager) Download(media integration.Media, release integration.Release, highPriority bool) error {

	var hadError bool
	for _, downloader := range dm.downloaders {
		dlType, ok := constant.DownloaderTypes[downloader.DownloaderType]
		if !ok {
			return errors.New("internal error. Downloader has no registered type")
		}
		if dlType == release.DownloadType {
			if err := downloader.AddRelease(&release); err == nil {
				dm.downloadsToMonitor = append(dm.downloadsToMonitor, release.Identifier)
				return nil
			} else {
				hadError = true
				logger.LogDanger(err)
			}
		}
	}

	if hadError {
		return errors.New("could not add release to downloader. See logs for details")
	}

	return errors.New("no downloaders for this type of release")
}

func (dm *DownloaderManager) DeleteDownloader(id int) {
	for i, downloader := range dm.downloaders {
		if downloader.ID == id {
			dm.downloaders = append(dm.downloaders[:i], dm.downloaders[i+1:]...)
			return
		}
	}
}
