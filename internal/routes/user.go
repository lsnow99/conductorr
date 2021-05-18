package routes

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/lsnow99/conductorr/internal/settings"
)

func SignIn(w http.ResponseWriter, r *http.Request) {

	Respond(w, nil, nil)
}

type SignUpInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	sui := SignUpInput{}
	if err := json.NewDecoder(r.Body).Decode(&sui); err != nil {
		Respond(w, err, nil)
		return
	}

}

func FirstTime(w http.ResponseWriter, r *http.Request) {
	if settings.ResetUser {
		Respond(w, nil, nil)
		return
	}
	Respond(w, errors.New("not first time"), nil)
}
