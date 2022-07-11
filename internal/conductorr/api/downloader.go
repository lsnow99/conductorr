package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/lsnow99/conductorr/internal/conductorr/app"
	"github.com/lsnow99/conductorr/internal/conductorr/dbstore"
	"github.com/lsnow99/conductorr/internal/conductorr/integration"
)

type DownloaderInput struct {
	Name           string                 `json:"name,omitempty"`
	DownloaderType string                 `json:"downloaderType,omitempty"`
	FileAction     string                 `json:"fileAction,omitempty"`
	Config         map[string]interface{} `json:"config,omitempty"`
}

type DownloaderResponse struct {
	ID             int                    `json:"id,omitempty"`
	Name           string                 `json:"name,omitempty"`
	DownloaderType string                 `json:"downloaderType,omitempty"`
	FileAction     string                 `json:"fileAction,omitempty"`
	Config         map[string]interface{} `json:"config,omitempty"`
}

func TestDownloader(w http.ResponseWriter, r *http.Request) {
	downloaderInput := DownloaderInput{}
	if err := json.NewDecoder(r.Body).Decode(&downloaderInput); err != nil {
		Respond(w, r, err, nil, true)
		return
	}
	downloader, err := integration.NewDownloaderFromConfig(downloaderInput.DownloaderType, downloaderInput.Config)
	if err != nil {
		Respond(w, r, err, nil, true)
		return
	}
	resp := make(map[string]interface{})
	err = downloader.TestConnection()
	resp["success"] = err == nil
	if err != nil {
		resp["msg"] = err.Error()
	}
	Respond(w, r, nil, resp, true)
}

func NewDownloader(w http.ResponseWriter, r *http.Request) {
	downloaderInput := DownloaderInput{}
	if err := json.NewDecoder(r.Body).Decode(&downloaderInput); err != nil {
		Respond(w, r, err, nil, true)
		return
	}
	id, err := dbstore.NewDownloader(downloaderInput.DownloaderType, downloaderInput.Name, downloaderInput.FileAction, downloaderInput.Config)
	if err != nil {
		Respond(w, r, err, nil, true)
		return
	}
	err = app.DM.RegisterDownloader(id, downloaderInput.DownloaderType, downloaderInput.Name, downloaderInput.Config)
	if err != nil {
		Respond(w, r, err, nil, true)
		return
	}
	Respond(w, r, err, nil, true)
}

func GetDownloaders(w http.ResponseWriter, r *http.Request) {
	dbDownloaders, err := dbstore.GetDownloaders()
	if err != nil {
		Respond(w, r, err, nil, true)
		return
	}
	downloaders := make([]DownloaderResponse, len(dbDownloaders))
	for i, dlr := range dbDownloaders {
		downloaders[i] = DownloaderResponse{
			ID:             dlr.ID,
			Name:           dlr.Name,
			DownloaderType: dlr.DownloaderType,
			FileAction:     dlr.FileAction,
			Config:         dlr.Config,
		}
	}
	Respond(w, r, nil, downloaders, true)
}

func UpdateDownloader(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		Respond(w, r, err, nil, true)
		return
	}
	downloaderInput := DownloaderInput{}
	if err := json.NewDecoder(r.Body).Decode(&downloaderInput); err != nil {
		Respond(w, r, err, nil, true)
		return
	}
	err = dbstore.UpdateDownloader(id, downloaderInput.Name, downloaderInput.FileAction, downloaderInput.Config)
	if err != nil {
		Respond(w, r, err, nil, true)
		return
	}
	downloader, err := dbstore.GetDownloader(id)
	if err != nil {
		Respond(w, r, err, nil, true)
		return
	}
	err = app.DM.RegisterDownloader(id, downloader.DownloaderType, downloader.Name, downloader.Config)
	Respond(w, r, err, nil, true)
}

func DeleteDownloader(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		Respond(w, r, err, nil, true)
		return
	}

	err = dbstore.DeleteDownloader(idInt)
	Respond(w, r, err, nil, true)
}
