package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/lsnow99/conductorr/internal/conductorr/app"
	"github.com/lsnow99/conductorr/internal/conductorr/dbstore"
	"github.com/lsnow99/conductorr/internal/conductorr/helpers"
)

type DownloadResponse struct {
	ID           int       `json:"id"`
	MediaID      int       `json:"mediaId"`
	FriendlyName string    `json:"friendlyName"`
	Identifier   string    `json:"identifier"`
	FinalDir     string    `json:"finalDir"`
	Status       string    `json:"status"`
	Started      time.Time `json:"started"`
	BytesLeft    uint64    `json:"bytesLeft"`
	FullSize     uint64    `json:"fullSize"`
}

func NewDownloadResponseFromManagedDownload(dl app.ManagedDownload) (dlr DownloadResponse) {
	dlr.ID = dl.ID
	dlr.MediaID = dl.MediaID
	dlr.Identifier = dl.Identifier
	dlr.FriendlyName = dl.FriendlyName
	dlr.Status = dl.Status
	dlr.Started = dl.Started
	dlr.BytesLeft = dl.BytesLeft
	dlr.FullSize = dl.FullSize
	return
}

func NewDownloadResponseFromDBDownload(dl dbstore.Download) (dlr DownloadResponse) {
	dlr.ID = dl.ID
	dlr.MediaID = int(dl.MediaID.Int32)
	dlr.Identifier = dl.Identifier
	dlr.FriendlyName = dl.FriendlyName
	dlr.Status = dl.Status
	return
}

func GetActiveDownloads(w http.ResponseWriter, r *http.Request) {
	mediaIDStr := r.URL.Query().Get("media_id")
	ids, err := getWatchingMediaIDs(mediaIDStr)
	if err != nil {
		Respond(w, r, err, nil, true)
	}

	downloads := app.DM.GetDownloads()
	downloadsResponse := make([]DownloadResponse, 0, len(downloads))
	for _, download := range downloads[:helpers.Min(len(downloads), 100)] {
		if shouldInclude(download.MediaID, ids) {
			downloadsResponse = append(downloadsResponse, NewDownloadResponseFromManagedDownload(download))
		}
	}
	Respond(w, r, nil, downloadsResponse, true)
}

func GetDoneDownloads(w http.ResponseWriter, r *http.Request) {
	mediaIDStr := r.URL.Query().Get("media_id")
	ids, err := getWatchingMediaIDs(mediaIDStr)
	if err != nil {
		Respond(w, r, err, nil, true)
	}

	downloads, err := dbstore.GetFinishedDownloads()
	if err != nil {
		Respond(w, r, err, nil, true)
	}
	downloadsResponse := make([]DownloadResponse, 0, len(downloads))
	for _, download := range downloads[:helpers.Min(len(downloads), 100)] {
		if shouldInclude(int(download.MediaID.Int32), ids) {
			downloadsResponse = append(downloadsResponse, NewDownloadResponseFromDBDownload(download))
		}
	}
	Respond(w, r, nil, downloadsResponse, true)
}

func shouldInclude(downloadMediaID int, ids []int) bool {
	// If the ids filter is empty, then let the download through
	if len(ids) == 0 {
		return true
	}
	for _, id := range ids {
		if downloadMediaID == id {
			// Return true if this download's media ID matches any id in the filter
			return true
		}
	}
	// Else return false
	return false
}

func getWatchingMediaIDs(mediaIDStr string) ([]int, error) {
	mediaID, _ := strconv.Atoi(mediaIDStr)
	if mediaID == 0 {
		return []int{}, nil
	}
	media, err := dbstore.GetMediaByID(mediaID)
	if err != nil {
		return nil, err
	}
	mr := NewMediaResponseFromDB(media)
	err = mr.Expand()
	if err != nil {
		return nil, err
	}

	return append([]int{mr.ID}, mr.getAllChildIDs()...), nil
}

func (mr MediaResponse) getAllChildIDs() (ids []int) {
	for _, child := range mr.Children {
		ids = append(ids, child.ID)
		ids = append(ids, child.getAllChildIDs()...)
	}
	return
}
