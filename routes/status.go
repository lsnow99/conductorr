package routes

import (
	"net/http"

	"github.com/lsnow99/conductorr/app"
)

func GetStatus(w http.ResponseWriter, r *http.Request) {
	Respond(w, r.Header.Get("hostname"), nil, app.Monitor.GetStatus(), true)
}
