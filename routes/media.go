package routes

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/lsnow99/conductorr/app"
	"github.com/lsnow99/conductorr/dbstore"
	"github.com/lsnow99/conductorr/integration"
	"github.com/lsnow99/conductorr/services/omdb"
)

type MediaInput struct {
	ID            int        `json:"id,omitempty"`
	Title         string     `json:"title,omitempty"`
	Description   string     `json:"description,omitempty"`
	ReleasedAt    time.Time  `json:"released_at,omitempty"`
	EndedAt       *time.Time `json:"ended_at,omitempty"`
	ContentType   string     `json:"content_type,omitempty"`
	Poster        string     `json:"poster,omitempty"`
	ParentMediaID int        `json:"parent_media_id,omitempty"`
	TmdbID        int        `json:"tmdb_id,omitempty"`
	ImdbID        string     `json:"imdb_id,omitempty"`
	TmdbRating    int        `json:"tmdb_rating,omitempty"`
	ImdbRating    int        `json:"imdb_rating,omitempty"`
	Runtime       int        `json:"runtime,omitempty"`
	ProfileID     int        `json:"profile_id,omitempty"`
}

func NewIntegrationMediaFromDBMedia(media dbstore.Media) (m integration.Media) {
	m.ID = media.ID
	if media.Title.Valid {
		m.Title = media.Title.String
	}
	if media.Description.Valid {
		m.Description = media.Description.String
	}
	return m
}

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

	id, err := dbstore.AddMedia(&result.Title, &result.Plot, &result.ReleasedAt, &result.EndedAt,
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

func GetMedia(w http.ResponseWriter, r *http.Request) {
	mediaIDStr := mux.Vars(r)["id"]
	mediaID, err := strconv.Atoi(mediaIDStr)
	if err != nil {
		Respond(w, r.Host, err, nil, true)
		return
	}
	media, err := dbstore.GetMediaByID(mediaID)
	Respond(w, r.Host, err, NewMediaResponseFromDB(media), true)
}

func UpdateMedia(w http.ResponseWriter, r *http.Request) {
	media := MediaInput{}
	if err := json.NewDecoder(r.Body).Decode(&media); err != nil {
		Respond(w, r.Host, err, nil, true)
		return
	}
	idStr := mux.Vars(r)["id"]
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		Respond(w, r.Host, err, nil, true)
		return
	}

	err = dbstore.UpdateMedia(idInt, media.ProfileID)
	Respond(w, r.Host, err, nil, true)
}

func DownloadMediaRelease(w http.ResponseWriter, r *http.Request) {
	mediaIDStr := mux.Vars(r)["id"]
	mediaID, err := strconv.Atoi(mediaIDStr)
	if err != nil {
		Respond(w, r.Host, err, nil, true)
		return
	}

	release := integration.Release{}
	if err := json.NewDecoder(r.Body).Decode(&release); err != nil {
		Respond(w, r.Host, err, nil, true)
		return
	}

	media, err := dbstore.GetMediaByID(mediaID)
	if err != nil {
		Respond(w, r.Host, err, nil, true)
		return
	}

	err = app.DM.Download(NewIntegrationMediaFromDBMedia(media), release, true)
	Respond(w, r.Host, err, nil, true)
}
