package routes

import (
	"net/http"
	"time"

	"github.com/lsnow99/conductorr/app"
	"github.com/lsnow99/conductorr/integration"
)

type DownloadResponse struct {
	MediaID      int       `json:"media_id,omitempty"`
	FriendlyName string    `json:"friendly_name,omitempty"`
	Identifier   string    `json:"identifier,omitempty"`
	FinalDir     string    `json:"final_dir,omitempty"`
	Status       string    `json:"status,omitempty"`
	Started      time.Time `json:"started,omitempty"`
	BytesLeft    uint64    `json:"bytes_left,omitempty"`
	FullSize     uint64    `json:"full_size,omitempty"`
}

func NewDownloadResponseFromIntegrationDownload(dl integration.Download) (dlr DownloadResponse) {
	dlr.MediaID = dl.MediaID
	dlr.Identifier = dl.Identifier
	dlr.FriendlyName = dl.FriendlyName
	dlr.Status = dl.Status
	dlr.Started = dl.Started
	dlr.BytesLeft = dl.BytesLeft
	dlr.FullSize = dl.FullSize
	return
}

func GetDownloads(w http.ResponseWriter, r *http.Request) {
	downloads := app.DM.GetDownloads()
	downloadsResponse := make([]DownloadResponse, len(downloads))
	for i, download := range downloads {
		downloadsResponse[i] = NewDownloadResponseFromIntegrationDownload(download)
	}
	Respond(w, r.Header.Get("hostname"), nil, downloadsResponse, true)
}
