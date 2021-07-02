package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/lsnow99/conductorr/app"
	"github.com/lsnow99/conductorr/dbstore"
	"github.com/lsnow99/conductorr/integration"
)

type IndexerInput struct {
	Name         string `json:"name"`
	UserID       int   `json:"user_id,omitempty"`
	BaseUrl      string `json:"base_url,omitempty"`
	ApiKey       string `json:"api_key,omitempty"`
	ForMovies    bool   `json:"for_movies,omitempty"`
	ForSeries    bool   `json:"for_series,omitempty"`
	DownloadType string `json:"download_type,omitempty"`
}

type IndexerResponse struct {
	IndexerInput
	ID int `json:"id,omitempty"`
}

func NewIndexerResponseFromDB(dbIndexer dbstore.Indexer) IndexerResponse {
	indexer := IndexerResponse{}
	indexer.ID = dbIndexer.ID
	indexer.Name = dbIndexer.Name
	indexer.BaseUrl = dbIndexer.BaseUrl
	indexer.ApiKey = dbIndexer.ApiKey
	indexer.ForMovies = dbIndexer.ForMovies
	indexer.ForSeries = dbIndexer.ForSeries
	indexer.DownloadType = dbIndexer.DownloadType
	return indexer
}

func GetIndexers(w http.ResponseWriter, r *http.Request) {
	dbIndexers, err := dbstore.GetIndexers()
	if err != nil {
		Respond(w, r.Host, err, nil, true)
		return
	}

	indexers := make([]IndexerResponse, len(dbIndexers))
	for i, dbIndexer := range dbIndexers {
		indexers[i] = NewIndexerResponseFromDB(dbIndexer)
	}

	Respond(w, r.Host, nil, indexers, true)
}

func GetIndexer(w http.ResponseWriter, r *http.Request) {
	indexerIDStr := mux.Vars(r)["id"]
	indexerID, err := strconv.Atoi(indexerIDStr)
	if err != nil {
		Respond(w, r.Host, err, nil, true)
		return
	}
	indexer, err := dbstore.GetIndexerByID(indexerID)
	Respond(w, r.Host, err, NewIndexerResponseFromDB(indexer), true)
}

func CreateIndexer(w http.ResponseWriter, r *http.Request) {
	indexer := IndexerInput{}
	if err := json.NewDecoder(r.Body).Decode(&indexer); err != nil {
		Respond(w, r.Host, err, nil, true)
		return
	}

	err := dbstore.CreateIndexer(indexer.Name, indexer.UserID, indexer.BaseUrl, indexer.ApiKey, indexer.ForMovies, indexer.ForSeries, indexer.DownloadType)
	Respond(w, r.Host, err, nil, true)
}

func UpdateIndexer(w http.ResponseWriter, r *http.Request) {
	indexer := IndexerInput{}
	if err := json.NewDecoder(r.Body).Decode(&indexer); err != nil {
		Respond(w, r.Host, err, nil, true)
		return
	}
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		Respond(w, r.Host, err, nil, true)
		return
	}

	err = dbstore.UpdateIndexer(id, indexer.Name, indexer.UserID, indexer.BaseUrl, indexer.ApiKey, indexer.ForMovies, indexer.ForSeries)
	if err != nil {
		Respond(w, r.Host, err, nil, true)
		return
	}
	app.IM.RegisterIndexer(id, indexer.DownloadType, indexer.UserID, indexer.Name, indexer.ApiKey, indexer.BaseUrl, indexer.ForMovies, indexer.ForSeries)
}

func DeleteIndexer(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		Respond(w, r.Host, err, nil, true)
		return
	}

	app.IM.DeleteIndexer(id)

	err = dbstore.DeleteIndexer(id)
	Respond(w, r.Host, err, nil, true)
}

func TestIndexer(w http.ResponseWriter, r *http.Request) {
	indexer := IndexerInput{}
	if err := json.NewDecoder(r.Body).Decode(&indexer); err != nil {
		Respond(w, r.Host, err, nil, true)
		return
	}
	xnab := integration.NewXnab(indexer.UserID, indexer.ApiKey, indexer.BaseUrl, indexer.Name, indexer.DownloadType)
	err := xnab.TestConnection()
	resp := make(map[string]interface{})
	resp["success"] = err == nil
	if err != nil {
		resp["msg"] = err.Error()
	}
	Respond(w, r.Host, nil, resp, true)
}
