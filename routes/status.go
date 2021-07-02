package routes

import (
	"net/http"

	"github.com/lsnow99/conductorr/app"
)

func GetStatus(w http.ResponseWriter, r *http.Request) {
	Respond(w, r.Host, nil, app.Monitor.GetStatus(), true)
}
