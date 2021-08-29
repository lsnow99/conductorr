package app

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"github.com/lsnow99/conductorr/constant"
	"github.com/lsnow99/conductorr/dbstore"
	"github.com/lsnow99/conductorr/integration"
	"github.com/lsnow99/conductorr/logger"
	"github.com/mholt/archiver/v3"
)

type ManagedDownload struct {
	ID int
	TryAgainOnFail bool
	integration.Download
}

type ManagedDownloader struct {
	ID             int
	Name           string
	DownloaderType string
	integration.Downloader
}

type DownloaderManager struct {
	sync.RWMutex
	didFirstRun bool
	backupReleaseMap map[int][]integration.Release
	downloaders []ManagedDownloader
	downloads   []ManagedDownload
}

func (md ManagedDownloader) GetName() string {
	return md.Name
}

func (dm *DownloaderManager) DoTask() {
	downloaders := dm.getDownloaders(false)
	downloads := make([]integration.Download, 0)
	for _, dlr := range downloaders {
		downloadsToMonitor := getDownloadsToMonitor(dm.downloads)
		updatedDownloads, _ := dlr.PollDownloads(downloadsToMonitor)
		// if err != nil {
		// 	logger.LogWarn(err)
		// }
		downloads = append(downloads, updatedDownloads...)
	}
	dm.processDownloads(downloads)
	dm.didFirstRun = true
}

func (dm *DownloaderManager) GetFrequency() time.Duration {
	return time.Second * 7
}

func (dm *DownloaderManager) GetDownloads() []ManagedDownload {
	dm.RLock()
	defer dm.RUnlock()
	return dm.downloads
}

func (dm *DownloaderManager) getDownloaders(haveLock bool) []ManagedDownloader {
	if haveLock {
		return dm.downloaders
	}
	dm.RLock()
	defer dm.RUnlock()
	return dm.downloaders
}

func (dm *DownloaderManager) RegisterDownloader(id int, downloaderType, name string, configuration map[string]interface{}) error {
	dm.Lock()
	defer dm.Unlock()

	dlr, err := integration.NewDownloaderFromConfig(downloaderType, configuration)
	if err != nil {
		return err
	}
	md := ManagedDownloader{
		ID:             id,
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

func (dm *DownloaderManager) AutoDownload(mediaID int, releases []integration.Release) {
	dm.Lock()
	defer dm.Unlock()
	dm.doAutoDownload(mediaID, releases)
}

func (dm *DownloaderManager) doAutoDownload(mediaID int, releases []integration.Release) {
	if len(releases) < 1 {
		return
	}
	dm.backupReleaseMap[mediaID] = releases[1:]
	if err := dm.Download(mediaID, releases[0], false, true, true); err != nil {
		logger.LogWarn(err)
	}
}

func (dm *DownloaderManager) Download(mediaID int, release integration.Release, highPriority, tryAgainOnFail, haveLock bool) error {

	downloaders := dm.getDownloaders(haveLock)

	var hadError bool
	for _, downloader := range downloaders {
		dlType, ok := constant.DownloaderTypes[downloader.DownloaderType]
		if !ok {
			return errors.New("internal error. Downloader has no registered type")
		}
		if dlType == release.DownloadType {
			if identifier, err := downloader.AddRelease(release); err == nil {
				id, err := dbstore.NewDownload(mediaID, downloader.ID, identifier, constant.StatusWaiting, release.Title)
				if err != nil {
					logger.LogDanger(fmt.Errorf("database error! could not save download: %v", err))
				}
				dm.Lock()
				md := ManagedDownload{}
				md.ID = id
				md.MediaID = mediaID
				md.FriendlyName = release.Title
				md.Identifier = identifier
				md.TryAgainOnFail = tryAgainOnFail
				md.Status = constant.StatusWaiting
				dm.downloads = append(dm.downloads, md)
				dm.Unlock()
				return err
			} else {
				hadError = true
				logger.LogDanger(err)
			}
		}
	}

	if hadError {
		return errors.New("could not add release to one or more downloaders. See logs for details")
	}

	return errors.New("no downloaders for this type of release")
}

func (dm *DownloaderManager) RegisterDownload(id int, mediaID int, friendlyName, status, identifier string) {
	dm.Lock()
	defer dm.Unlock()

	md := ManagedDownload{}
	md.ID = id
	md.MediaID = mediaID
	md.FriendlyName = friendlyName
	md.Identifier = identifier
	md.Status = status
	dm.downloads = append(dm.downloads, md)
}

func (dm *DownloaderManager) DeleteDownloader(id int) {
	dm.Lock()
	defer dm.Unlock()
	for i, downloader := range dm.downloaders {
		if downloader.ID == id {
			dm.downloaders = append(dm.downloaders[:i], dm.downloaders[i+1:]...)
			return
		}
	}
}

// processDownloads
func (dm *DownloaderManager) processDownloads(curState []integration.Download) {

	dm.Lock()
	defer dm.Unlock()

	for _, curStateDL := range curState {
		for i, prevStateDL := range dm.downloads {
			if curStateDL.Identifier == prevStateDL.Identifier {
				dm.downloads[i].FinalDir = curStateDL.FinalDir
				dm.downloads[i].BytesLeft = curStateDL.BytesLeft
				dm.downloads[i].FullSize = curStateDL.FullSize
				dm.downloads[i].Started = curStateDL.Started
				if prevStateDL.Status == constant.StatusCProcessing {
					continue
				}
				if curStateDL.Status != prevStateDL.Status || !dm.didFirstRun {
					updateDBStatus(curStateDL.Identifier, curStateDL.Status)
					dm.downloads[i].Status = curStateDL.Status
					switch curStateDL.Status {
					case constant.StatusError:
						if prevStateDL.TryAgainOnFail {
							dm.doAutoDownload(dm.downloads[i].MediaID, dm.backupReleaseMap[dm.downloads[i].MediaID])
						}
						logger.LogWarn(fmt.Errorf("release failed to download for %s", curStateDL.FriendlyName))
					case constant.StatusComplete:
						if prevStateDL.Status == constant.StatusCProcessing || prevStateDL.Status == constant.StatusDone {
							continue
						}
						// Trigger conductorr post processing
						dm.downloads[i].Status = constant.StatusCProcessing
						go dm.handleCompletedDownload(dm.downloads[i])
					// Do nothing for these statuses
					case constant.StatusCError:
						if prevStateDL.TryAgainOnFail {
							dm.doAutoDownload(dm.downloads[i].MediaID, dm.backupReleaseMap[dm.downloads[i].MediaID])
						}
						logger.LogDanger(fmt.Errorf("conductorr had an error processing %v", curStateDL))
						// case constant.StatusCProcessing:
					}
				}
				break
			}
		}
	}
}

func updateDBStatus(identifier, status string) {
	go func(identifier, status string) {
		err := dbstore.UpdateDownloadStatusByIdentifier(identifier, status)
		if err != nil {
			logger.LogDanger(err)
		}
	}(identifier, status)
}

func (dm *DownloaderManager) handleCompletedDownload(download ManagedDownload) {
	var outputFile string
	var dbPath dbstore.Path
	var season, series dbstore.Media
	media, err := dbstore.GetMediaByID(download.MediaID)
	if err != nil {
		logger.LogDanger(err)
		updateDBStatus(download.Identifier, constant.StatusCError)
		return
	}
	if media.ParentMediaID.Valid {
		season, err = dbstore.GetMediaByID(int(media.ParentMediaID.Int32))
		if err != nil {
			logger.LogDanger(err)
			updateDBStatus(download.Identifier, constant.StatusCError)
			return
		}
		if season.ParentMediaID.Valid {
			series, err = dbstore.GetMediaByID(int(season.ParentMediaID.Int32))
			if err != nil {
				logger.LogDanger(err)
				updateDBStatus(download.Identifier, constant.StatusCError)
				return
			}
			dbPath, err = dbstore.GetPath(int(series.PathID.Int32))
			if err != nil {
				logger.LogDanger(err)
				updateDBStatus(download.Identifier, constant.StatusCError)
				return
			}
			// Rename to /path/Show Name (Year)/Season 0X/Episode Title - S0XE0Y
			// (add extension later)
			outputFile = fmt.Sprintf("%s (%d)", series.Title.String, series.ReleasedAt.Time.Year())
			outputFile = filepath.Join(outputFile, fmt.Sprintf("Season %02d", season.Number.Int32))
			outputFile = filepath.Join(outputFile, fmt.Sprintf("%s - S%02dE%02d", media.Title.String, season.Number.Int32, media.Number.Int32))
		} else {
			logger.LogDanger(fmt.Errorf("bad media hierarchy"))
			updateDBStatus(download.Identifier, constant.StatusCError)
			return
		}
	} else {
		dbPath, err = dbstore.GetPath(int(media.PathID.Int32))
		if err != nil {
			logger.LogDanger(err)
			updateDBStatus(download.Identifier, constant.StatusCError)
			return
		}
		outputFile = media.Title.String + " (" + strconv.Itoa(media.ReleasedAt.Time.Year()) + ")"
		outputFile = filepath.Join(outputFile, outputFile)
	}
	// Find completed download file from directory
	dlPath := download.FinalDir
	videoPath, err := getVideoPath(dlPath)
	if err != nil {
		logger.LogDanger(err)
	}

	if videoPath == "" {
		logger.LogDanger(fmt.Errorf("could not find output file for %v", download.FriendlyName))
		updateDBStatus(download.Identifier, constant.StatusCError)
		return
	}

	destFilepath := filepath.Join(dbPath.Path, outputFile) + filepath.Ext(videoPath)
	destFiledir := filepath.Dir(destFilepath)
	if err := os.MkdirAll(destFiledir, 0777); err != nil {
		logger.LogDanger(err)
		updateDBStatus(download.Identifier, constant.StatusCError)
		return
	}
	destFile, err := os.OpenFile(destFilepath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		logger.LogDanger(err)
		updateDBStatus(download.Identifier, constant.StatusCError)
		return
	}
	srcFile, err := os.OpenFile(videoPath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		logger.LogDanger(err)
		updateDBStatus(download.Identifier, constant.StatusCError)
		return
	}
	n, err := io.Copy(destFile, srcFile)
	if err != nil {
		logger.LogDanger(fmt.Errorf("got error after copying %d bytes: %v", n, err))
		updateDBStatus(download.Identifier, constant.StatusCError)
		return
	}
	logger.LogInfo(fmt.Errorf("successfully copied %s to %s", videoPath, destFilepath))

	dm.Lock()

	for i, dmdl := range dm.downloads {
		if download.Identifier == dmdl.Identifier {
			dm.downloads = append(dm.downloads[:i], dm.downloads[i+1:]...)
			break
		}
	}

	dm.Unlock()

	err = dbstore.UpdateDownloadStatusByIdentifier(download.Identifier, constant.StatusDone)
	if err != nil {
		logger.LogDanger(err)
	}

	err = MSM.ImportMedia(destFiledir)
	if err != nil {
		logger.LogDanger(err)
	}

	if series.ID > 0 {
		err = dbstore.SetMediaPath(series.ID, filepath.Dir(destFiledir))
		if err != nil {
			logger.LogDanger(err)
		}
		err = dbstore.SetMediaPath(season.ID, destFiledir)
		if err != nil {
			logger.LogDanger(err)
		}
	}
	err = dbstore.SetMediaPath(media.ID, destFilepath)
	if err != nil {
		logger.LogDanger(err)
	}
}

// getDownloadsToMonitor convert a slice of downloads to a slice of identifiers
func getDownloadsToMonitor(downloads []ManagedDownload) (monitoring []string) {
	for _, dl := range downloads {
		monitoring = append(monitoring, dl.Identifier)
	}
	return
}

func getVideoPath(dlPath string) (string, error) {
	fmt.Println("getting videos in ", dlPath)
	var rarfile string
	err := filepath.WalkDir(dlPath, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			fmt.Println(e)
			return e
		}
		fmt.Println(s, d.Name(), filepath.Ext(d.Name()))
		if filepath.Ext(d.Name()) == ".rar" {
			rarfile = s
		}
		return nil
	})
	if err != nil {
		return "", err
	}

	candidates, err := getVideoCandidates(dlPath)
	if err != nil {
		return "", err
	}

	if rarfile != "" {
		tDir, err := os.MkdirTemp("", "conductorr")
		if err != nil {
			return "", err
		}
		err = archiver.Unarchive(rarfile, tDir)
		if err != nil {
			return "", err
		}
		unarchivedCandidates, err := getVideoCandidates(tDir)
		if err != nil {
			return "", err
		}
		candidates = append(candidates, unarchivedCandidates...)
	}

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
	return bestCandidate, nil
}

func getVideoCandidates(dlPath string) ([]string, error) {
	candidates := make([]string, 0)
	err := filepath.WalkDir(dlPath, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if filepath.Ext(d.Name()) == ".mkv" || filepath.Ext(d.Name()) == ".mp4" {
			candidates = append(candidates, s)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return candidates, nil
}
