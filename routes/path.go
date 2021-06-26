package routes

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/lsnow99/conductorr/dbstore"
)

type Path struct {
	Path          string `json:"path,omitempty"`
	MoviesDefault bool   `json:"movies_default,omitempty"`
	SeriesDefault bool   `json:"series_default,omitempty"`
}

func NewDBPathFromPath(path Path) dbstore.Path {
	return dbstore.Path{
		Path:          path.Path,
		MoviesDefault: path.MoviesDefault,
		SeriesDefault: path.SeriesDefault,
	}
}

func NewPathFromDBPath(p dbstore.Path) Path {
	return Path{
		Path:          p.Path,
		MoviesDefault: p.MoviesDefault,
		SeriesDefault: p.SeriesDefault,
	}
}

func SetPaths(w http.ResponseWriter, r *http.Request) {
	paths := make([]Path, 0)
	if err := json.NewDecoder(r.Body).Decode(&paths); err != nil {
		Respond(w, r.Host, err, nil, true)
		return
	}
	dbPaths := make([]dbstore.Path, len(paths))
	for i, path := range paths {
		dbPaths[i] = NewDBPathFromPath(path)
	}
	err := dbstore.UpdatePaths(r.Context(), dbPaths)
	Respond(w, r.Host, err, nil, true)
}

func GetPaths(w http.ResponseWriter, r *http.Request) {
	dbPaths, err := dbstore.GetPaths()
	if err != nil && err != sql.ErrNoRows {
		Respond(w, r.Host, err, nil, true)
		return
	}
	paths := make([]Path, len(dbPaths))
	for i, path := range dbPaths {
		paths[i] = NewPathFromDBPath(path)
	}
	Respond(w, r.Host, nil, paths, true)
}

func DeletePath(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		Respond(w, r.Host, err, nil, true)
		return
	}

	err = dbstore.DeletePath(idInt)
	Respond(w, r.Host, err, nil, true)
}
