package app

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/lsnow99/conductorr/constant"
	"github.com/lsnow99/conductorr/dbstore"
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
	downloads          []integration.Download
}

func (md ManagedDownloader) GetName() string {
	return md.Name
}

func (dm *DownloaderManager) DoTask() {
	downloads := make([]integration.Download, 0)
	for _, dlr := range dm.downloaders {
		downloadsToMonitor := dm.getDownloadsToMonitor()
		if err := dlr.PollDownloads(downloadsToMonitor); err != nil {
			logger.LogWarn(err)
		}
		downloads = append(downloads, dlr.GetDownloads()...)
	}
	dm.processDownloads(downloads)
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

func (dm *DownloaderManager) Download(mediaID int, release integration.Release, highPriority bool) error {

	var hadError bool
	for _, downloader := range dm.downloaders {
		dlType, ok := constant.DownloaderTypes[downloader.DownloaderType]
		if !ok {
			return errors.New("internal error. Downloader has no registered type")
		}
		if dlType == release.DownloadType {
			if err := downloader.AddRelease(&release); err == nil {
				dm.downloads = append(dm.downloads, integration.Download{
					MediaID: mediaID,
					FriendlyName: release.FriendlyName,
					Identifier: release.Identifier,
					Status: constant.StatusWaiting,
				})
				_, err := dbstore.NewDownload(mediaID, downloader.ID, release.Identifier, constant.StatusWaiting, release.FriendlyName)
				if err != nil {
					logger.LogDanger(fmt.Errorf("database error! could not save download: %v", err))
				}
				return err
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

func (dm *DownloaderManager) RegisterDownload(mediaID int, friendlyName, status, identifier string) {
	dm.downloads = append(dm.downloads, integration.Download{
		MediaID: mediaID,
		FriendlyName: friendlyName,
		Identifier: identifier,
		Status: status,
	})
}

func (dm *DownloaderManager) DeleteDownloader(id int) {
	for i, downloader := range dm.downloaders {
		if downloader.ID == id {
			dm.downloaders = append(dm.downloaders[:i], dm.downloaders[i+1:]...)
			return
		}
	}
}

func (dm *DownloaderManager) processDownloads(curState []integration.Download) {
	for _, curStateDL := range curState {
		var found bool
		for i, prevStateDL := range dm.downloads {
			if curStateDL.Identifier == prevStateDL.Identifier {
				found = true
				if curStateDL.Status != prevStateDL.Status {
					err := dbstore.UpdateDownloadStatusByIdentifier(curStateDL.Identifier, curStateDL.Status)
					if err != nil {
						logger.LogDanger(err)
						continue
					}
					dm.downloads[i].Status = curStateDL.Status
					switch curStateDL.Status {
					case constant.StatusError:
						// TODO Trigger re-download
					case constant.StatusComplete:
						// Trigger conductorr post processing
						dm.downloads[i].Status = constant.StatusCProcessing
						go handleCompletedDownload(dm.downloads[i])
					// Do nothing for these statuses
					case constant.StatusCError:
						logger.LogDanger(fmt.Errorf("conductorr had an error processing %v", curStateDL))
					// case constant.StatusCProcessing:
					case constant.StatusDone:
						// Remove download from the list of monitored downloads
						dm.downloads = append(dm.downloads[:i], dm.downloads[i+1:]...)
					}
				}
				break
			}
		}
		if !found {
			logger.LogDanger(fmt.Errorf("did not find a previous state for download %v", curStateDL))
		}
	}
}

func handleCompletedDownload(download integration.Download) {
	media, err := dbstore.GetMediaByID(download.MediaID)
	if err != nil {
		logger.LogDanger(err)
		return
	}
	if !media.PathID.Valid {
		logger.LogDanger(fmt.Errorf("no path assigned for %v", media))
		return
	}
	dbPath, err := dbstore.GetPath(int(media.PathID.Int32))
	if err != nil {
		logger.LogDanger(err)
		return
	}
	// Find completed download file from directory
	dlPath := download.FinalDir
	var candidates []string
	filepath.WalkDir(dlPath, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if filepath.Ext(d.Name()) == "mkv" || filepath.Ext(d.Name()) == "mp4" {
			candidates = append(candidates, s)
		}
		return nil
	})
	var bestCandidate string
	var largestSize int64 = -1
	for _, candidate := range candidates {
		fi, err := os.Stat(candidate)
		if err != nil {
			continue
		}
		size := fi.Size()
		if size > largestSize {
			bestCandidate = candidate
			largestSize = size
		}
	}
	if bestCandidate == "" {
		logger.LogDanger(fmt.Errorf("could not find output file for %v", download.FriendlyName))
		return
	}
	destFile, err := os.OpenFile(filepath.Join(dbPath.Path, "testout"), os.O_CREATE | os.O_TRUNC | os.O_WRONLY, os.ModePerm)
	if err != nil {
		logger.LogDanger(err)
		return
	}
	srcFile, err := os.OpenFile(bestCandidate, os.O_RDONLY, os.ModePerm)
	if err != nil {
		logger.LogDanger(err)
		return
	}
	n, err := io.Copy(destFile, srcFile)
	if err != nil {
		logger.LogDanger(fmt.Errorf("got error after copying %d bytes: %v", n, err))
		return
	}
	logger.LogInfo(fmt.Errorf("successfully copied %s", download.FriendlyName))
}

func (dm *DownloaderManager) getDownloadsToMonitor() (monitoring []string) {
	for _, dl := range dm.downloads {
		monitoring = append(monitoring, dl.Identifier)
	}
	return
}