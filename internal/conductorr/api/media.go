package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/lsnow99/conductorr/internal/conductorr/app"
	"github.com/lsnow99/conductorr/internal/conductorr/dbstore"
	"github.com/lsnow99/conductorr/internal/conductorr/integration"
	"github.com/lsnow99/conductorr/internal/conductorr/services/search"
	"github.com/lsnow99/conductorr/internal/conductorr/services/series"
	"golang.org/x/net/html"
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
	PathID        int        `json:"path_id,omitempty"`
}

type MonitoringInput struct {
	Monitoring bool `json:"monitoring"`
}

type AddMediaInput struct {
	ProfileID  int  `json:"profile_id,omitempty"`
	PathID     int  `json:"path_id,omitempty"`
	Monitoring bool `json:"monitoring"`
}

func AddMedia(w http.ResponseWriter, r *http.Request) {
	searchID := mux.Vars(r)["search_id"]

	mi := AddMediaInput{}
	if err := json.NewDecoder(r.Body).Decode(&mi); err != nil {
		Respond(w, r, err, nil, true)
		return
	}

	result, err := search.GetResultByID(searchID)
	if err != nil {
		Respond(w, r, err, nil, true)
		return
	}

	id, err := doAddMedia(result, &mi.ProfileID, &mi.PathID, mi.Monitoring)
	if err != nil {
		Respond(w, r, err, nil, true)
		return
	}
	Respond(w, r, nil, id, true)
}

func RefreshMediaMetadata(w http.ResponseWriter, r *http.Request) {
	mediaIDStr := mux.Vars(r)["id"]

	mediaID, err := strconv.Atoi(mediaIDStr)
	if err != nil {
		Respond(w, r, err, nil, true)
		return
	}

	media, err := dbstore.GetMediaByID(mediaID)
	if err != nil {
		Respond(w, r, err, nil, true)
		return
	}
	if !media.ImdbID.Valid {
		Respond(w, r, errors.New("imdb id is nil for media"), nil, true)
		return
	}

	result, err := search.GetResultByImdbID(media.ImdbID.String)
	if err != nil {
		Respond(w, r, err, nil, true)
		return
	}

	id, err := doAddMedia(result, nil, nil, media.Monitoring)
	if err != nil {
		Respond(w, r, err, nil, true)
		return
	}
	Respond(w, r, nil, id, true)
}

func doAddMedia(result *search.IndividualResult, profileID *int, pathID *int, monitor bool) (int, error) {
	if result == nil {
		return 0, fmt.Errorf("nil search result passed to add media")
	}

	// Scale imdb rating
	imdbRating := int(result.Rating * 10)

	// Fetch the poster
	posterResp, err := http.Get(result.Poster)
	if err != nil {
		return 0, err
	}
	poster, err := ioutil.ReadAll(posterResp.Body)
	if err != nil {
		return 0, err
	}

	var id int

	// If tv show, add all the episodes and seasons
	if result.ContentType == "series" {
		episodes, err := series.GetEpisodes(result.ImdbID)
		if err != nil {
			return id, err
		}
		id, err = dbstore.UpsertMedia(&result.Title, &result.Plot, &result.ReleasedAt, &result.EndedAt,
			&result.ContentType, nil, nil, &result.ImdbID, result.TvdbID, nil, &imdbRating, &result.Runtime,
			&poster, result.Genres, profileID, pathID, nil, monitor)
		if err != nil {
			return id, err
		}

		insertedSeasons := make(map[int]int)
		for _, episode := range episodes {
			seasonID, ok := insertedSeasons[episode.Season]
			if !ok {
				seasonStr := "Season " + strconv.Itoa(episode.Season)
				contentType := "season"
				seasonID, err = dbstore.UpsertMedia(&seasonStr, nil, &episode.Aired, nil, &contentType, &id, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &episode.Season, monitor)
				if err != nil {
					return id, err
				}
				insertedSeasons[episode.Season] = seasonID
			}

			description := getTextFromHTML(episode.Description)
			contentType := "episode"
			_, err = dbstore.UpsertMedia(&episode.Title, &description, &episode.Aired, nil, &contentType, &seasonID, nil, nil, nil, nil, nil, &episode.Runtime, nil, nil, nil, nil, &episode.Episode, monitor)
			if err != nil {
				return id, err
			}
		}
	} else {
		id, err = dbstore.UpsertMedia(&result.Title, &result.Plot, &result.ReleasedAt, &result.EndedAt,
			&result.ContentType, nil, nil, &result.ImdbID, result.TvdbID, nil, &imdbRating, &result.Runtime,
			&poster, result.Genres, profileID, pathID, nil, monitor)
		if err != nil {
			return id, err
		}
	}
	return id, nil
}

func getTextFromHTML(markup string) string {
	domDocTest := html.NewTokenizer(strings.NewReader(markup))
	previousStartTokenTest := domDocTest.Token()
	text := ""
loopDomTest:
	for {
		tt := domDocTest.Next()
		switch {
		case tt == html.ErrorToken:
			break loopDomTest // End of the document,  done
		case tt == html.StartTagToken:
			previousStartTokenTest = domDocTest.Token()
		case tt == html.TextToken:
			if previousStartTokenTest.Data == "script" {
				continue
			}
			text += strings.TrimSpace(html.UnescapeString(string(domDocTest.Text())))
		}
	}
	return text
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
		Respond(w, r, err, nil, true)
		return
	}
	media, err := dbstore.GetMediaByID(mediaID)
	if err != nil {
		Respond(w, r, err, nil, true)
		return
	}
	mr := NewMediaResponseFromDB(media)
	err = mr.Expand()
	if err != nil {
		Respond(w, r, err, nil, true)
		return
	}
	Respond(w, r, nil, mr, true)
}

func DeleteMedia(w http.ResponseWriter, r *http.Request) {
	mediaIDStr := mux.Vars(r)["id"]
	mediaID, err := strconv.Atoi(mediaIDStr)
	if err != nil {
		Respond(w, r, err, nil, true)
		return
	}
	err = dbstore.DeleteMedia(mediaID)
	Respond(w, r, err, nil, true)
}

func UpdateMedia(w http.ResponseWriter, r *http.Request) {
	media := MediaInput{}
	if err := json.NewDecoder(r.Body).Decode(&media); err != nil {
		Respond(w, r, err, nil, true)
		return
	}
	idStr := mux.Vars(r)["id"]
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		Respond(w, r, err, nil, true)
		return
	}

	err = dbstore.UpdateMedia(idInt, media.ProfileID, media.PathID)
	Respond(w, r, err, nil, true)
}

func DownloadMediaRelease(w http.ResponseWriter, r *http.Request) {
	mediaIDStr := mux.Vars(r)["id"]
	mediaID, err := strconv.Atoi(mediaIDStr)
	if err != nil {
		Respond(w, r, err, nil, true)
		return
	}

	release := integration.Release{}
	if err := json.NewDecoder(r.Body).Decode(&release); err != nil {
		Respond(w, r, err, nil, true)
		return
	}

	media, err := dbstore.GetMediaByID(mediaID)
	if err != nil {
		Respond(w, r, err, nil, true)
		return
	}

	err = app.DM.Download(media.ID, release, true, false, false)
	Respond(w, r, err, nil, true)
}

func SetMonitoringMedia(w http.ResponseWriter, r *http.Request) {
	mediaIDStr := mux.Vars(r)["id"]
	mediaID, err := strconv.Atoi(mediaIDStr)
	if err != nil {
		Respond(w, r, err, nil, true)
		return
	}

	mi := MonitoringInput{}
	if err := json.NewDecoder(r.Body).Decode(&mi); err != nil {
		Respond(w, r, err, nil, true)
		return
	}

	err = dbstore.UpdateMediaMonitoring(mediaID, mi.Monitoring)
	Respond(w, r, err, nil, true)
}

func GetRecentMedia(w http.ResponseWriter, r *http.Request) {
	medias, err := dbstore.GetRecentlyAddedMedia(5)
	if err != nil {
		Respond(w, r, err, nil, true)
		return
	}

	mediasResponse := make([]MediaResponse, 0, len(medias))
	for _, media := range medias {
		mr := NewMediaResponseFromDB(media)
		mediasResponse = append(mediasResponse, mr)
	}

	Respond(w, r, nil, mediasResponse, true)
}
