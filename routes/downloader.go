package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/lsnow99/conductorr/dbstore"
	"github.com/lsnow99/conductorr/integration"
	"github.com/lsnow99/conductorr/app"
)

type DownloaderInput struct {
	Name           string                 `json:"name,omitempty"`
	DownloaderType string                 `json:"downloader_type,omitempty"`
	Config         map[string]interface{} `json:"config,omitempty"`
}

type DownloaderResponse struct {
	ID             int                    `json:"id,omitempty"`
	Name           string                 `json:"name,omitempty"`
	DownloaderType string                 `json:"downloader_type,omitempty"`
	Config         map[string]interface{} `json:"config,omitempty"`
}

func TestDownloader(w http.ResponseWriter, r *http.Request) {
	downloaderInput := DownloaderInput{}
	if err := json.NewDecoder(r.Body).Decode(&downloaderInput); err != nil {
		Respond(w, r.Header.Get("hostname"), err, nil, true)
		return
	}
	downloader, err := integration.NewDownloaderFromConfig(downloaderInput.DownloaderType, downloaderInput.Config)
	if err != nil {
		Respond(w, r.Header.Get("hostname"), err, nil, true)
		return
	}
	resp := make(map[string]interface{})
	err = downloader.TestConnection()
	resp["success"] = err == nil
	if err != nil {
		resp["msg"] = err.Error()
	}
	Respond(w, r.Header.Get("hostname"), nil, resp, true)
}

func NewDownloader(w http.ResponseWriter, r *http.Request) {
	downloaderInput := DownloaderInput{}
	if err := json.NewDecoder(r.Body).Decode(&downloaderInput); err != nil {
		Respond(w, r.Header.Get("hostname"), err, nil, true)
		return
	}
	id, err := dbstore.NewDownloader(downloaderInput.DownloaderType, downloaderInput.Name, downloaderInput.Config)
	if err != nil {
		Respond(w, r.Header.Get("hostname"), err, nil, true)
		return
	}
	err = app.DM.RegisterDownloader(id, downloaderInput.DownloaderType, downloaderInput.Name, downloaderInput.Config)
	if err != nil {
		Respond(w, r.Header.Get("hostname"), err, nil, true)
		return
	}
	Respond(w, r.Header.Get("hostname"), err, nil, true)
}

func GetDownloaders(w http.ResponseWriter, r *http.Request) {
	dbDownloaders, err := dbstore.GetDownloaders()
	if err != nil {
		Respond(w, r.Header.Get("hostname"), err, nil, true)
		return
	}
	downloaders := make([]DownloaderResponse, len(dbDownloaders))
	for i, dlr := range dbDownloaders {
		downloaders[i] = DownloaderResponse{
			ID:             dlr.ID,
			Name:           dlr.Name,
			DownloaderType: dlr.DownloaderType,
			Config:         dlr.Config,
		}
	}
	Respond(w, r.Header.Get("hostname"), nil, downloaders, true)
}

func UpdateDownloader(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		Respond(w, r.Header.Get("hostname"), err, nil, true)
		return
	}
	downloaderInput := DownloaderInput{}
	if err := json.NewDecoder(r.Body).Decode(&downloaderInput); err != nil {
		Respond(w, r.Header.Get("hostname"), err, nil, true)
		return
	}
	err = dbstore.UpdateDownloader(id, downloaderInput.DownloaderType, downloaderInput.Name, downloaderInput.Config)
	if err != nil {
		Respond(w, r.Header.Get("hostname"), err, nil, true)
		return
	}
	err = app.DM.RegisterDownloader(id, downloaderInput.DownloaderType, downloaderInput.Name, downloaderInput.Config)
	Respond(w, r.Header.Get("hostname"), err, nil, true)
}

func DeleteDownloader(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		Respond(w, r.Header.Get("hostname"), err, nil, true)
		return
	}

	err = dbstore.DeleteDownloader(idInt)
	Respond(w, r.Header.Get("hostname"), err, nil, true)
}
