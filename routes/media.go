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
	"github.com/lsnow99/conductorr/services/series"
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
	PathID     int        `json:"path_id,omitempty"`
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

type AddMediaInput struct {
	ProfileID int `json:"profile_id,omitempty"`
	PathID    int `json:"path_id,omitempty"`
}

func AddMedia(w http.ResponseWriter, r *http.Request) {
	imdbID := mux.Vars(r)["imdb_id"]

	mi := AddMediaInput{}
	if err := json.NewDecoder(r.Body).Decode(&mi); err != nil {
		Respond(w, r.Header.Get("hostname"), err, nil, true)
		return
	}

	result, err := omdb.SearchByImdbID(imdbID)
	if err != nil {
		Respond(w, r.Header.Get("hostname"), err, nil, true)
		return
	}

	// Scale imdb rating
	imdbRating := int(result.ImdbRating * 10)

	// Fetch the poster
	posterResp, err := http.Get(result.Poster)
	if err != nil {
		Respond(w, r.Header.Get("hostname"), err, nil, true)
		return
	}
	poster, err := ioutil.ReadAll(posterResp.Body)
	if err != nil {
		Respond(w, r.Header.Get("hostname"), err, nil, true)
		return
	}

	genres := strings.Split(result.Genre, ", ")

	var id int
	// If tv show, add all the episodes and seasons
	if result.Type == "series" {
		episodes, err := series.GetEpisodes(result.ImdbID)
		if err != nil {
			Respond(w, r.Header.Get("hostname"), err, 0, true)
			return
		}
		id, err = dbstore.AddMedia(&result.Title, &result.Plot, &result.ReleasedAt, &result.EndedAt,
			&result.Type, nil, nil, &result.ImdbID, nil, &imdbRating, &result.Runtime,
			&poster, genres, &mi.ProfileID, &mi.PathID)
		if err != nil {
			Respond(w, r.Header.Get("hostname"), err, id, true)
			return
		}
	
		insertedSeasons := make(map[int]int)
		for _, episode := range episodes {
			seasonID, ok := insertedSeasons[episode.Season]
			if !ok {
				seasonStr := "Season " + strconv.Itoa(episode.Season)
				contentType := "season"
				seasonID, err = dbstore.AddMedia(&seasonStr, nil, &episode.Aired, nil, &contentType, &id, nil, nil, nil, nil, nil, nil, nil, nil, nil)
				if err != nil {
					Respond(w, r.Header.Get("hostname"), err, 0, true)
					return
				}
				insertedSeasons[episode.Season] = seasonID
			}
			contentType := "episode"
			_, err = dbstore.AddMedia(&episode.Title, &episode.Description, &episode.Aired, nil, &contentType, &seasonID, nil, nil, nil, nil, &episode.Runtime, nil, nil, nil, nil)
			if err != nil {
				Respond(w, r.Header.Get("hostname"), err, 0, true)
				return
			}
		}
	} else {
		id, err = dbstore.AddMedia(&result.Title, &result.Plot, &result.ReleasedAt, &result.EndedAt,
			&result.Type, nil, nil, &result.ImdbID, nil, &imdbRating, &result.Runtime,
			&poster, genres, &mi.ProfileID, &mi.PathID)
		if err != nil {
			Respond(w, r.Header.Get("hostname"), err, id, true)
			return
		}
	}

	Respond(w, r.Header.Get("hostname"), nil, id, true)
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
		Respond(w, r.Header.Get("hostname"), err, nil, true)
		return
	}
	media, err := dbstore.GetMediaByID(mediaID)
	Respond(w, r.Header.Get("hostname"), err, NewMediaResponseFromDB(media), true)
}

func DeleteMedia(w http.ResponseWriter, r *http.Request) {
	mediaIDStr := mux.Vars(r)["id"]
	mediaID, err := strconv.Atoi(mediaIDStr)
	if err != nil {
		Respond(w, r.Header.Get("hostname"), err, nil, true)
		return
	}
	err = dbstore.DeleteMedia(mediaID)
	Respond(w, r.Header.Get("hostname"), err, nil, true)
}

func UpdateMedia(w http.ResponseWriter, r *http.Request) {
	media := MediaInput{}
	if err := json.NewDecoder(r.Body).Decode(&media); err != nil {
		Respond(w, r.Header.Get("hostname"), err, nil, true)
		return
	}
	idStr := mux.Vars(r)["id"]
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		Respond(w, r.Header.Get("hostname"), err, nil, true)
		return
	}

	err = dbstore.UpdateMedia(idInt, media.ProfileID, media.PathID)
	Respond(w, r.Header.Get("hostname"), err, nil, true)
}

func DownloadMediaRelease(w http.ResponseWriter, r *http.Request) {
	mediaIDStr := mux.Vars(r)["id"]
	mediaID, err := strconv.Atoi(mediaIDStr)
	if err != nil {
		Respond(w, r.Header.Get("hostname"), err, nil, true)
		return
	}

	release := integration.Release{}
	if err := json.NewDecoder(r.Body).Decode(&release); err != nil {
		Respond(w, r.Header.Get("hostname"), err, nil, true)
		return
	}

	media, err := dbstore.GetMediaByID(mediaID)
	if err != nil {
		Respond(w, r.Header.Get("hostname"), err, nil, true)
		return
	}

	err = app.DM.Download(media.ID, release, true)
	Respond(w, r.Header.Get("hostname"), err, nil, true)
}
