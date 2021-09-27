package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/lsnow99/conductorr/app"
	"github.com/lsnow99/conductorr/dbstore"
	"github.com/lsnow99/conductorr/integration"
	"github.com/lsnow99/conductorr/services/plex"
)

type FetchPlexAuthTokenInput struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type AuthTokenResponse struct {
	Token string `json:"token"`
}

type MediaServerResponse struct {
	ID              int                    `json:"id,omitempty"`
	Name            string                 `json:"name,omitempty"`
	MediaServerType string                 `json:"media_server_type,omitempty"`
	Config          map[string]interface{} `json:"config,omitempty"`
}

func FetchPlexAuthToken(w http.ResponseWriter, r *http.Request) {
	fetchInput := FetchPlexAuthTokenInput{}

	if err := json.NewDecoder(r.Body).Decode(&fetchInput); err != nil {
		Respond(w, r.Header.Get("hostname"), err, nil, true)
		return
	}

	token, err := plex.RetrievePlexToken(fetchInput.Username, fetchInput.Password)
	if err != nil {
		Respond(w, r.Header.Get("hostname"), err, nil, true)
		return
	}

	re := AuthTokenResponse{
		Token: token,
	}

	Respond(w, r.Header.Get("hostname"), err, re, true)
}

type MediaServerInput struct {
	Name            string                 `json:"name,omitempty"`
	MediaServerType string                 `json:"media_server_type,omitempty"`
	Config          map[string]interface{} `json:"config,omitempty"`
}

func TestMediaServer(w http.ResponseWriter, r *http.Request) {
	mediaServerInput := MediaServerInput{}
	if err := json.NewDecoder(r.Body).Decode(&mediaServerInput); err != nil {
		Respond(w, r.Header.Get("hostname"), err, nil, true)
		return
	}
	mediaServer, err := integration.NewMediaServerFromConfig(mediaServerInput.MediaServerType, mediaServerInput.Config)
	if err != nil {
		Respond(w, r.Header.Get("hostname"), err, nil, true)
		return
	}
	err = mediaServer.TestConnection()
	Respond(w, r.Header.Get("hostname"), err, nil, true)
}

func NewMediaServer(w http.ResponseWriter, r *http.Request) {
	mediaServerInput := MediaServerInput{}
	if err := json.NewDecoder(r.Body).Decode(&mediaServerInput); err != nil {
		Respond(w, r.Header.Get("hostname"), err, nil, true)
		return
	}
	id, err := dbstore.NewMediaServer(mediaServerInput.MediaServerType, mediaServerInput.Name, mediaServerInput.Config)
	if err != nil {
		Respond(w, r.Header.Get("hostname"), err, nil, true)
		return
	}
	err = app.MSM.RegisterMediaServer(id, mediaServerInput.MediaServerType, mediaServerInput.Name, mediaServerInput.Config)
	if err != nil {
		Respond(w, r.Header.Get("hostname"), err, nil, true)
		return
	}
	Respond(w, r.Header.Get("hostname"), err, nil, true)
}

func GetMediaServers(w http.ResponseWriter, r *http.Request) {
	dbMediaServers, err := dbstore.GetMediaServers()
	if err != nil {
		Respond(w, r.Header.Get("hostname"), err, nil, true)
		return
	}
	mediaServers := make([]MediaServerResponse, len(dbMediaServers))
	for i, msvr := range dbMediaServers {
		mediaServers[i] = MediaServerResponse{
			ID:              msvr.ID,
			Name:            msvr.Name,
			MediaServerType: msvr.MediaServerType,
			Config:          msvr.Config,
		}
	}
	Respond(w, r.Header.Get("hostname"), nil, mediaServers, true)
}

func UpdateMediaServer(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		Respond(w, r.Header.Get("hostname"), err, nil, true)
		return
	}
	mediaServerInput := MediaServerInput{}
	if err := json.NewDecoder(r.Body).Decode(&mediaServerInput); err != nil {
		Respond(w, r.Header.Get("hostname"), err, nil, true)
		return
	}
	err = dbstore.UpdateMediaServer(id, mediaServerInput.Name, mediaServerInput.Config)
	if err != nil {
		Respond(w, r.Header.Get("hostname"), err, nil, true)
		return
	}
	mediaServer, err := dbstore.GetMediaServer(id)
	if err != nil {
		Respond(w, r.Header.Get("hostname"), err, nil, true)
		return
	}
	err = app.MSM.RegisterMediaServer(id, mediaServer.MediaServerType, mediaServer.Name, mediaServer.Config)
	Respond(w, r.Header.Get("hostname"), err, nil, true)
}

func DeleteMediaServer(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		Respond(w, r.Header.Get("hostname"), err, nil, true)
		return
	}

	err = dbstore.DeleteMediaServer(idInt)
	Respond(w, r.Header.Get("hostname"), err, nil, true)
}
