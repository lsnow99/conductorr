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

type ManagedDownloader struct {
	ID             int
	Name           string
	DownloaderType string
	integration.Downloader
}

type DownloaderManager struct {
	*sync.RWMutex
	didFirstRun bool
	downloaders []ManagedDownloader
	downloads   []integration.Download
}

func (md ManagedDownloader) GetName() string {
	return md.Name
}

func (dm *DownloaderManager) DoTask() {
	downloads := make([]integration.Download, 0)
	for _, dlr := range dm.downloaders {
		downloadsToMonitor := dm.getDownloadsToMonitor()
		updatedDownloads, err := dlr.PollDownloads(downloadsToMonitor)
		if err != nil {
			logger.LogWarn(err)
		}
		downloads = append(downloads, updatedDownloads...)
	}
	dm.processDownloads(downloads)
	dm.didFirstRun = true
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

func (dm *DownloaderManager) Download(mediaID int, release integration.Release, highPriority bool) error {

	var hadError bool
	for _, downloader := range dm.downloaders {
		dlType, ok := constant.DownloaderTypes[downloader.DownloaderType]
		if !ok {
			return errors.New("internal error. Downloader has no registered type")
		}
		if dlType == release.DownloadType {
			if identifier, err := downloader.AddRelease(release); err == nil {
				dm.downloads = append(dm.downloads, integration.Download{
					MediaID:      mediaID,
					FriendlyName: release.Title,
					Identifier:   identifier,
					Status:       constant.StatusWaiting,
				})
				_, err := dbstore.NewDownload(mediaID, downloader.ID, identifier, constant.StatusWaiting, release.Title)
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
		return errors.New("could not add release to one or more downloaders. See logs for details")
	}

	return errors.New("no downloaders for this type of release")
}

func (dm *DownloaderManager) RegisterDownload(mediaID int, friendlyName, status, identifier string) {
	dm.downloads = append(dm.downloads, integration.Download{
		MediaID:      mediaID,
		FriendlyName: friendlyName,
		Identifier:   identifier,
		Status:       status,
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
		for i, prevStateDL := range dm.downloads {
			if curStateDL.Identifier == prevStateDL.Identifier {
				dm.downloads[i].FinalDir = curStateDL.FinalDir
				dm.downloads[i].BytesLeft = curStateDL.BytesLeft
				dm.downloads[i].FullSize = curStateDL.FullSize
				dm.downloads[i].Started = curStateDL.Started
				if curStateDL.Status != prevStateDL.Status || !dm.didFirstRun {
					err := dbstore.UpdateDownloadStatusByIdentifier(curStateDL.Identifier, curStateDL.Status)
					if err != nil {
						logger.LogDanger(err)
						continue
					}
					dm.downloads[i].Status = curStateDL.Status
					switch curStateDL.Status {
					case constant.StatusError:
						// TODO Trigger re-download
						logger.LogWarn(fmt.Errorf("release failed to download for %s", curStateDL.FriendlyName))
					case constant.StatusComplete:
						if prevStateDL.Status == constant.StatusCProcessing || prevStateDL.Status == constant.StatusDone {
							continue
						}
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
	}
}

func handleCompletedDownload(download integration.Download) {
	var outputFile string
	var dbPath dbstore.Path
	media, err := dbstore.GetMediaByID(download.MediaID)
	if err != nil {
		logger.LogDanger(err)
		return
	}
	if media.ParentMediaID.Valid {
		season, err := dbstore.GetMediaByID(int(media.ParentMediaID.Int32))
		if err != nil {
			logger.LogDanger(err)
			return
		}
		if season.ParentMediaID.Valid {
			series, err := dbstore.GetMediaByID(int(season.ParentMediaID.Int32))
			if err != nil {
				logger.LogDanger(err)
				return
			}
			dbPath, err = dbstore.GetPath(int(series.PathID.Int32))
			if err != nil {
				logger.LogDanger(err)
				return
			}
			outputFile = fmt.Sprintf("%s (%d)", series.Title.String, series.ReleasedAt.Time.Year())
			outputFile = filepath.Join(outputFile, fmt.Sprintf("Season %02d", season.Number.Int32))
			outputFile = filepath.Join(outputFile, fmt.Sprintf("%s S%02dE%02d", media.Title.String, season.Number.Int32,media.Number.Int32))
		} else {
			logger.LogDanger(fmt.Errorf("bad hierarchy"))
			return
		}
	} else {
		dbPath, err = dbstore.GetPath(int(media.PathID.Int32))
		if err != nil {
			logger.LogDanger(err)
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
		return
	}

	destFilepath := filepath.Join(dbPath.Path, outputFile) + filepath.Ext(videoPath)
	destFiledir := filepath.Dir(destFilepath)
	if err := os.MkdirAll(destFiledir, 0777); err != nil {
		logger.LogDanger(err)
		return
	}
	destFile, err := os.OpenFile(destFilepath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		logger.LogDanger(err)
		return
	}
	srcFile, err := os.OpenFile(videoPath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		logger.LogDanger(err)
		return
	}
	n, err := io.Copy(destFile, srcFile)
	if err != nil {
		logger.LogDanger(fmt.Errorf("got error after copying %d bytes: %v", n, err))
		return
	}
	logger.LogInfo(fmt.Errorf("successfully copied %s to %s", videoPath, destFilepath))
	err = dbstore.UpdateDownloadStatusByIdentifier(download.Identifier, constant.StatusDone)
	if err != nil {
		logger.LogDanger(err)
	}
	err = MSM.ImportMedia(destFiledir)
	if err != nil {
		logger.LogDanger(err)
	}
}

func (dm *DownloaderManager) getDownloadsToMonitor() (monitoring []string) {
	for _, dl := range dm.downloads {
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
