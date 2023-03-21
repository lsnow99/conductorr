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

	"github.com/lsnow99/conductorr/internal/conductorr/dbstore"
	"github.com/lsnow99/conductorr/internal/conductorr/integration"
	"github.com/lsnow99/conductorr/pkg/constant"
	"github.com/mholt/archiver/v3"
	"github.com/rs/zerolog/log"
)

type ManagedDownload struct {
	ID             int
	TryAgainOnFail bool
	ReleaseID      *string
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
	didFirstRun      bool
	backupReleaseMap map[int][]integration.Release
	downloaders      []ManagedDownloader
	downloads        []ManagedDownload
}

func (md ManagedDownloader) GetName() string {
	return md.Name
}

func (dm *DownloaderManager) DoTask() {
	downloaders := dm.getDownloaders(false)
	downloads := make([]integration.Download, 0)
	for _, dlr := range downloaders {
		downloadsToMonitor := getDownloadsToMonitor(dm.downloads, dlr.ID)
		updatedDownloads, err := dlr.PollDownloads(downloadsToMonitor)
		if err != nil {
			log.Warn().
				Err(err).
				Msg("error polling for downloads")
		}
		downloads = append(downloads, updatedDownloads...)
	}
	dm.processDownloads(downloads)
	dm.Lock()
	dm.didFirstRun = true
	dm.Unlock()
}

func (dm *DownloaderManager) GetTaskName() string {
	return "Downloader Manager"
}

func (dm *DownloaderManager) GetFrequency() time.Duration {
	return time.Second * 5
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
	for index := 0; index < len(releases); index++ {
		log.Info().
			Int("media_id", mediaID).
			Str("release_title", releases[index].Title).
			Str("release_id", releases[index].ID).
			Msgf("attempting to download release, %d remaining", len(releases)-1)
		if err := dm.Download(mediaID, releases[index], false, true, true); err == nil {
			// Return if there are no errors, after adding the remaining releases to the backup map
			dm.backupReleaseMap[mediaID] = releases[index+1:]
			break
		} else {
			log.Warn().
				Err(err).
				Int("media_id", mediaID).
				Str("release_title", releases[index].Title).
				Str("release_id", releases[index].ID).
				Msg("error downloading release")
		}
	}
}

func (dm *DownloaderManager) Download(mediaID int, release integration.Release, highPriority, tryAgainOnFail, haveLock bool) error {

	downloaders := dm.getDownloaders(haveLock)

	var hadError bool
	for _, downloader := range downloaders {
		dlType, ok := constant.DownloaderTypes[downloader.DownloaderType]
    log.Debug().Str(dlType, release.DownloadType)
		if !ok {
			return errors.New("internal error. Downloader has no registered type")
		}
		if dlType == release.DownloadType {
			if identifier, err := downloader.AddRelease(release); err == nil {
				id, err := dbstore.NewDownload(mediaID, downloader.ID, identifier, constant.StatusWaiting, release.Title, release.ID)
				if err != nil {
					log.Error().
						Err(err).
						Msg("could not save download")
				}
				if !haveLock {
					dm.Lock()
				}
				md := ManagedDownload{}
				md.ReleaseID = &release.ID
				md.ID = id
				md.MediaID = mediaID
				md.FriendlyName = release.Title
				md.Identifier = identifier
				md.TryAgainOnFail = tryAgainOnFail
				md.Status = constant.StatusWaiting
				dm.downloads = append(dm.downloads, md)
				if !haveLock {
					dm.Unlock()
				}
				return nil
			} else {
				hadError = true
				log.Error().
					Err(err).
					Msg("could not add release to downloader")
			}
		}
	}

	if hadError {
		return errors.New("could not add release to one or more downloaders. See logs for details")
	}

	return errors.New("no downloaders for this type of release")
}

func (dm *DownloaderManager) RegisterDownload(id int, mediaID int, friendlyName, status, identifier string, releaseID *string) {
	dm.Lock()
	defer dm.Unlock()

	md := ManagedDownload{}
	md.ReleaseID = releaseID
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
						log.Warn().
							Str("download_identifier", curStateDL.Identifier).
							Msgf("failed to download release %s", curStateDL.FriendlyName)
						if prevStateDL.ReleaseID != nil {
							if err := dbstore.UpdateReleaseHistory(prevStateDL.MediaID, *prevStateDL.ReleaseID); err != nil {
								log.Error().
									Stack().
									Err(err).
									Msg("failed to update release history")
							}
						}
					case constant.StatusComplete:
						if prevStateDL.Status == constant.StatusCProcessing || prevStateDL.Status == constant.StatusDone {
							continue
						}
						// Trigger conductorr post processing
						dm.downloads[i].Status = constant.StatusCProcessing
						go dm.handleCompletedDownload(dm.downloads[i])
					case constant.StatusCError:
						if prevStateDL.TryAgainOnFail {
							dm.doAutoDownload(dm.downloads[i].MediaID, dm.backupReleaseMap[dm.downloads[i].MediaID])
						}
						log.Error().
							Int("download_id", curStateDL.ID).
							Msgf("conductorr had an error processing download id %d", curStateDL.ID)
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
			log.Error().
				Stack().
				Err(err).
				Str("download_identifier", identifier).
				Msg("failed to update download status in database")
		}
	}(identifier, status)
}

func (dm *DownloaderManager) handleCompletedDownload(download ManagedDownload) {
	var outputFile string
	var dbPath dbstore.Path
	var season, series dbstore.Media
	media, err := dbstore.GetMediaByID(download.MediaID)
	if err != nil {
		log.Error().
			Stack().
			Err(err).
			Int("media_id", download.MediaID).
			Int("download_id", download.ID).
			Msg("failed to get media from database")

		updateDBStatus(download.Identifier, constant.StatusCError)
		return
	}
	if media.ParentMediaID.Valid {
		season, err = dbstore.GetMediaByID(int(media.ParentMediaID.Int32))
		if err != nil {
			log.Error().
				Stack().
				Err(err).
				Int("media_id", download.MediaID).
				Int("parent_media_id", int(media.ParentMediaID.Int32)).
				Int("download_id", download.ID).
				Msg("failed to get media from database")

			updateDBStatus(download.Identifier, constant.StatusCError)
			return
		}
		if season.ParentMediaID.Valid {
			series, err = dbstore.GetMediaByID(int(season.ParentMediaID.Int32))
			if err != nil {
				log.Error().
					Stack().
					Err(err).
					Int("media_id", download.MediaID).
					Int("parent_media_id", int(season.ParentMediaID.Int32)).
					Int("parent_parent_media_id", int(season.ParentMediaID.Int32)).
					Int("download_id", download.ID).
					Msg("failed to get media from database")

				updateDBStatus(download.Identifier, constant.StatusCError)
				return
			}
			dbPath, err = dbstore.GetPath(int(series.PathID.Int32))
			if err != nil {
				log.Error().
					Stack().
					Err(err).
					Int("media_id", series.ID).
					Int("path_id", int(series.PathID.Int32)).
					Int("download_id", download.ID).
					Msg("failed to get path from database")

				updateDBStatus(download.Identifier, constant.StatusCError)
				return
			}
			// Rename to /path/Show Name (Year)/Season 0X/Episode Title - S0XE0Y
			// (add extension later)
			outputFile = fmt.Sprintf("%s (%d)", series.Title.String, series.ReleasedAt.Time.Year())
			outputFile = filepath.Join(outputFile, fmt.Sprintf("Season %02d", season.ItemNumber.Int32))
			outputFile = filepath.Join(outputFile, fmt.Sprintf("%s - S%02dE%02d", media.Title.String, season.ItemNumber.Int32, media.ItemNumber.Int32))
		} else {
			log.Error().
				Int("media_id", download.MediaID).
				Int("parent_media_id", int(season.ParentMediaID.Int32)).
				Int("parent_parent_media_id", int(season.ParentMediaID.Int32)).
				Int("download_id", download.ID).
				Bool("internal", true).
				Msg("bad media hierarchy")

			updateDBStatus(download.Identifier, constant.StatusCError)
			return
		}
	} else {
		dbPath, err = dbstore.GetPath(int(media.PathID.Int32))
		if err != nil {
			log.Error().
				Stack().
				Err(err).
				Int("media_id", media.ID).
				Int("path_id", int(series.PathID.Int32)).
				Int("download_id", download.ID).
				Msg("failed to get path from database")

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
		log.Error().
			Stack().
			Err(err).
			Str("search_dir", dlPath).
			Int("media_id", download.MediaID).
			Int("download_id", download.ID).
			Msg("error walking output directory")

		updateDBStatus(download.Identifier, constant.StatusCError)
		return
	}

	if videoPath == "" {
		log.Error().
			Err(err).
			Str("search_dir", dlPath).
			Int("media_id", download.MediaID).
			Int("download_id", download.ID).
			Msgf("could not find output file for %s", download.FriendlyName)

		updateDBStatus(download.Identifier, constant.StatusCError)
		return
	}

	destFilepath := filepath.Join(dbPath.Path, outputFile) + filepath.Ext(videoPath)
	destFiledir := filepath.Dir(destFilepath)
	if err := os.MkdirAll(destFiledir, 0777); err != nil {
		log.Error().
			Err(err).
			Str("dest_filedir", destFiledir).
			Int("download_id", download.ID).
			Msg("could not build output directories")

		updateDBStatus(download.Identifier, constant.StatusCError)
		return
	}
	destFile, err := os.OpenFile(destFilepath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		log.Error().
			Err(err).
			Str("dest_filepath", destFilepath).
			Str("dest_filedir", destFiledir).
			Int("download_id", download.ID).
			Msg("could not open output file")

		updateDBStatus(download.Identifier, constant.StatusCError)
		return
	}
	srcFile, err := os.OpenFile(videoPath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Error().
			Err(err).
			Str("video_path", videoPath).
			Int("download_id", download.ID).
			Msg("could not open source file")

		updateDBStatus(download.Identifier, constant.StatusCError)
		return
	}
	n, err := io.Copy(destFile, srcFile)
	if err != nil {
		log.Error().
			Err(err).
			Int("download_id", download.ID).
			Msgf("got error after copying %d bytes", n)

		updateDBStatus(download.Identifier, constant.StatusCError)
		return
	}

	log.Info().
		Int("download_id", download.ID).
		Msgf("successfully copied %s to %s", videoPath, destFilepath)

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
		log.Error().
			Stack().
			Err(err).
			Int("download_id", download.ID).
			Msg("error updating download status")
	}

	err = MSM.ImportMedia(destFiledir)
	if err != nil {
		log.Error().
			Stack().
			Err(err).
			Int("download_id", download.ID).
			Msg("error importing completed download into media server")
	}

	if series.ID > 0 {
		err = dbstore.SetMediaPath(series.ID, filepath.Dir(destFiledir))
		if err != nil {
			log.Error().
				Stack().
				Err(err).
				Int("download_id", download.ID).
				Str("path", destFiledir).
				Int("media_id", series.ID).
				Msg("failed to set series media path")
		}
		err = dbstore.SetMediaPath(season.ID, destFiledir)
		if err != nil {
			log.Error().
				Stack().
				Err(err).
				Int("download_id", download.ID).
				Str("path", destFiledir).
				Int("media_id", season.ID).
				Msg("failed to set season media path")
		}
	}
	err = dbstore.SetMediaPath(media.ID, destFilepath)
	if err != nil {
		log.Error().
			Stack().
			Err(err).
			Int("download_id", download.ID).
			Str("path", destFiledir).
			Int("media_id", media.ID).
			Msg("failed to set media path")
	}
}

// getDownloadsToMonitor convert a slice of downloads to a slice of identifiers
func getDownloadsToMonitor(downloads []ManagedDownload, downloaderID int) (monitoring []string) {
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
			return e
		}
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
		tDir, err := integration.MkdirTemp("conductorr")
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
