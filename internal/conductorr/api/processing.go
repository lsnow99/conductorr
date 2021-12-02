package api

import (
	"encoding/json"
	"net/http"

	"github.com/lsnow99/conductorr/internal/conductorr/integration"
)

type PathInput struct {
	Path string `json:"path,omitempty"`
}

func TestPath(w http.ResponseWriter, r *http.Request) {
	pathInput := PathInput{}
	if err := json.NewDecoder(r.Body).Decode(&pathInput); err != nil {
		Respond(w, r, err, nil, true)
		return
	}
	err := integration.CheckPath(pathInput.Path)
	resp := make(map[string]interface{})
	resp["success"] = err == nil
	if err != nil {
		resp["msg"] = err.Error()
	}
	Respond(w, r, nil, resp, true)
}
