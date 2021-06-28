package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/lsnow99/conductorr/dbstore"
	"github.com/lsnow99/conductorr/integration"
)

type DownloaderInput struct {
	DownloaderType string                 `json:"downloader_type,omitempty"`
	Config         map[string]interface{} `json:"config,omitempty"`
}

type DownloaderResponse struct {
	ID             int                    `json:"id,omitempty"`
	DownloaderType string                 `json:"downloader_type,omitempty"`
	Config         map[string]interface{} `json:"config,omitempty"`
}

func TestDownloader(w http.ResponseWriter, r *http.Request) {
	downloaderInput := DownloaderInput{}
	if err := json.NewDecoder(r.Body).Decode(&downloaderInput); err != nil {
		Respond(w, r.Host, err, nil, true)
		return
	}
	downloader, err := integration.NewDownloaderFromConfig(downloaderInput.DownloaderType, downloaderInput.Config)
	if err != nil {
		Respond(w, r.Host, err, nil, true)
		return
	}
	resp := make(map[string]interface{})
	if err := downloader.Init(); err != nil {
		resp["success"] = err == nil
		resp["msg"] = err.Error()
		Respond(w, r.Host, nil, resp, true)
		return
	}
	err = downloader.TestConnection()
	resp["success"] = err == nil
	if err != nil {
		resp["msg"] = err.Error()
	}
	Respond(w, r.Host, nil, resp, true)
}

func NewDownloader(w http.ResponseWriter, r *http.Request) {
	downloaderInput := DownloaderInput{}
	if err := json.NewDecoder(r.Body).Decode(&downloaderInput); err != nil {
		Respond(w, r.Host, err, nil, true)
		return
	}
	err := dbstore.NewDownloader(downloaderInput.DownloaderType, downloaderInput.Config)
	Respond(w, r.Host, err, nil, true)
}

func GetDownloaders(w http.ResponseWriter, r *http.Request) {
	dbDownloaders, err := dbstore.GetDownloaders()
	if err != nil {
		Respond(w, r.Host, err, nil, true)
		return
	}
	downloaders := make([]DownloaderResponse, len(dbDownloaders))
	for i, dlr := range dbDownloaders {
		downloaders[i] = DownloaderResponse{
			ID: dlr.ID,
			DownloaderType: dlr.DownloaderType,
			Config:         dlr.Config,
		}
	}
	Respond(w, r.Host, nil, downloaders, true)
}

func UpdateDownloader(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		Respond(w, r.Host, err, nil, true)
		return
	}
	downloaderInput := DownloaderInput{}
	if err := json.NewDecoder(r.Body).Decode(&downloaderInput); err != nil {
		Respond(w, r.Host, err, nil, true)
		return
	}
	err = dbstore.UpdateDownloader(idInt, downloaderInput.Config)
	Respond(w, r.Host, err, nil, true)
}
