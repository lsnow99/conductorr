package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/lsnow99/conductorr/internal/conductorr/backup"
)

func CreateNewBackup(w http.ResponseWriter, r *http.Request) {
	id, err := backup.CreateBackup(r.Context())
	if err != nil {
		Respond(w, r, err, nil, true)
		return
	}
	url := fmt.Sprintf("/api/v1/backup/%d", id)
	data := make(map[string]interface{})
	data["url"] = url
	Respond(w, r, nil, data, true)
}

func GetBackupFile(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	bkupFile, timestamp, ok := backup.GetBackupFile(id)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	filename := fmt.Sprintf("backup_%s.zip", timestamp.Format("20060102T150405"))
	filename = strings.ReplaceAll(filename, " ", "_")
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	http.ServeFile(w, r, bkupFile)
}

func RestoreBackupFromFile(w http.ResponseWriter, r *http.Request) {
	
}
