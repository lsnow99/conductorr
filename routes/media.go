package routes

import (
	"database/sql"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/lsnow99/conductorr/dbstore"
	"github.com/lsnow99/conductorr/services/omdb"
)

func AddMedia(w http.ResponseWriter, r *http.Request) {
	imdbID := mux.Vars(r)["imdb_id"]

	result, err := omdb.SearchByImdbID(imdbID)
	if err != nil {
		Respond(w, r.Host, err, nil, true)
		return
	}

	// Scale imdb rating
	imdbRating := int(result.ImdbRating * 10)

	// Fetch the poster
	posterResp, err := http.Get(result.Poster)
	if err != nil {
		Respond(w, r.Host, err, nil, true)
		return
	}
	poster, err := ioutil.ReadAll(posterResp.Body)
	if err != nil {
		Respond(w, r.Host, err, nil, true)
		return
	}

	genres := strings.Split(result.Genre, ", ")

	id, err := dbstore.AddMedia(&result.Title, &result.Plot, nil, nil,
		&result.Type, nil, nil, &result.ImdbID, nil, &imdbRating, &result.Runtime,
		&poster, genres)
	Respond(w, r.Host, err, id, true)
}

func GetPoster(w http.ResponseWriter, r *http.Request) {
	mediaIDStr := mux.Vars(r)["id"]
	mediaID, err := strconv.Atoi(mediaIDStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	poster, err := dbstore.GetPoster(mediaID)
	if err != nil && err != sql.ErrNoRows {
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(poster)
}
