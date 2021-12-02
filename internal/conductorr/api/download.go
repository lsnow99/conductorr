package api

import (
	"net/http"
	"time"

	"github.com/lsnow99/conductorr/internal/conductorr/app"
	"github.com/lsnow99/conductorr/internal/conductorr/dbstore"
)

type DownloadResponse struct {
	ID           int       `json:"id"`
	MediaID      int       `json:"media_id"`
	FriendlyName string    `json:"friendly_name"`
	Identifier   string    `json:"identifier"`
	FinalDir     string    `json:"final_dir"`
	Status       string    `json:"status"`
	Started      time.Time `json:"started"`
	BytesLeft    uint64    `json:"bytes_left"`
	FullSize     uint64    `json:"full_size"`
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
	downloads := app.DM.GetDownloads()
	downloadsResponse := make([]DownloadResponse, len(downloads))
	for i, download := range downloads {
		downloadsResponse[i] = NewDownloadResponseFromManagedDownload(download)
	}
	Respond(w, r, nil, downloadsResponse, true)
}

func GetDoneDownloads(w http.ResponseWriter, r *http.Request) {
	downloads, err := dbstore.GetFinishedDownloads()
	if err != nil {
		Respond(w, r, err, nil, true)
	}
	downloadsResponse := make([]DownloadResponse, len(downloads))
	for i, download := range downloads {
		downloadsResponse[i] = NewDownloadResponseFromDBDownload(download)
	}
	Respond(w, r, nil, downloadsResponse, true)
}
