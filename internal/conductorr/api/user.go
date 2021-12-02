package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/lsnow99/conductorr/internal/conductorr/dbstore"
	"github.com/lsnow99/conductorr/internal/conductorr/settings"
)

type AuthInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	creds := AuthInput{}
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		Respond(w, r, err, nil, false)
		return
	}
	if err := dbstore.CheckUser(r.Context(), creds.Username, creds.Password); err != nil {
		Respond(w, r, err, nil, false)
		return
	}
	Respond(w, r, nil, nil, true)
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	creds := AuthInput{}
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		Respond(w, r, err, nil, false)
		return
	}
	if err := dbstore.SetUser(r.Context(), creds.Username, creds.Password); err != nil {
		Respond(w, r, err, nil, false)
		return
	}
	// Don't allow further registrations for the duration of this server instance
	settings.ResetUser = false
	Respond(w, r, nil, nil, true)
}

func FirstTime(w http.ResponseWriter, r *http.Request) {
	if settings.ResetUser {
		Respond(w, r, nil, nil, false)
		return
	}
	Respond(w, r, errors.New("not first time"), nil, false)
}
