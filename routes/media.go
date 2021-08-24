package routes

import (
	"database/sql"
	"encoding/json"
	"errors"
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
	ProfileID  int  `json:"profile_id,omitempty"`
	PathID     int  `json:"path_id,omitempty"`
	Monitoring bool `json:"monitoring"`
}

func AddMedia(w http.ResponseWriter, r *http.Request) {
	imdbID := mux.Vars(r)["imdb_id"]

	mi := AddMediaInput{}
	if err := json.NewDecoder(r.Body).Decode(&mi); err != nil {
		Respond(w, r.Header.Get("hostname"), err, nil, true)
		return
	}

	id, err := doAddMedia(imdbID, &mi.ProfileID, &mi.PathID, mi.Monitoring)
	if err != nil {
		Respond(w, r.Header.Get("hostname"), err, nil, true)
		return
	}
	Respond(w, r.Header.Get("hostname"), nil, id, true)
}

func RefreshMediaMetadata(w http.ResponseWriter, r *http.Request) {
	mediaIDStr := mux.Vars(r)["id"]

	mediaID, err := strconv.Atoi(mediaIDStr)
	if err != nil {
		Respond(w, r.Header.Get("hostname"), err, nil, true)
		return
	}

	media, err := dbstore.GetMediaByID(mediaID)
	if err != nil {
		Respond(w, r.Header.Get("hostname"), err, nil, true)
		return
	}
	if !media.ImdbID.Valid {
		Respond(w, r.Header.Get("hostname"), errors.New("imdb id is nil for media"), nil, true)
		return
	}

	id, err := doAddMedia(media.ImdbID.String, nil, nil, media.Monitoring)
	if err != nil {
		Respond(w, r.Header.Get("hostname"), err, nil, true)
		return
	}
	Respond(w, r.Header.Get("hostname"), nil, id, true)
}

func doAddMedia(imdbID string, profileID *int, pathID *int, monitor bool) (int, error) {
	result, err := omdb.SearchByImdbID(imdbID)
	if err != nil {
		return 0, err
	}

	// Scale imdb rating
	imdbRating := int(result.ImdbRating * 10)

	// Fetch the poster
	posterResp, err := http.Get(result.Poster)
	if err != nil {
		return 0, err
	}
	poster, err := ioutil.ReadAll(posterResp.Body)
	if err != nil {
		return 0, err
	}

	genres := strings.Split(result.Genre, ", ")

	var id int

	// If tv show, add all the episodes and seasons
	if result.Type == "series" {
		episodes, err := series.GetEpisodes(result.ImdbID)
		if err != nil {
			return id, err
		}
		id, err = dbstore.UpsertMedia(&result.Title, &result.Plot, &result.ReleasedAt, &result.EndedAt,
			&result.Type, nil, nil, &result.ImdbID, nil, &imdbRating, &result.Runtime,
			&poster, genres, profileID, pathID, nil, monitor)
		if err != nil {
			return id, err
		}

		insertedSeasons := make(map[int]int)
		for _, episode := range episodes {
			seasonID, ok := insertedSeasons[episode.Season]
			if !ok {
				seasonStr := "Season " + strconv.Itoa(episode.Season)
				contentType := "season"
				seasonID, err = dbstore.UpsertMedia(&seasonStr, nil, &episode.Aired, nil, &contentType, &id, nil, nil, nil, nil, nil, nil, nil, nil, nil, &episode.Season, monitor)
				if err != nil {
					return id, err
				}
				insertedSeasons[episode.Season] = seasonID
			}

			description := getTextFromHTML(episode.Description)
			contentType := "episode"
			_, err = dbstore.UpsertMedia(&episode.Title, &description, &episode.Aired, nil, &contentType, &seasonID, nil, nil, nil, nil, &episode.Runtime, nil, nil, nil, nil, &episode.Episode, monitor)
			if err != nil {
				return id, err
			}
		}
	} else {
		id, err = dbstore.UpsertMedia(&result.Title, &result.Plot, &result.ReleasedAt, &result.EndedAt,
			&result.Type, nil, nil, &result.ImdbID, nil, &imdbRating, &result.Runtime,
			&poster, genres, profileID, pathID, nil, monitor)
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
		Respond(w, r.Header.Get("hostname"), err, nil, true)
		return
	}
	media, err := dbstore.GetMediaByID(mediaID)
	if err != nil {
		Respond(w, r.Header.Get("hostname"), err, nil, true)
		return
	}
	mr := NewMediaResponseFromDB(media)
	err = mr.Expand()
	if err != nil {
		Respond(w, r.Header.Get("hostname"), err, nil, true)
		return
	}
	Respond(w, r.Header.Get("hostname"), nil, mr, true)
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

func SetMonitoringMedia(w http.ResponseWriter, r *http.Request) {
	mediaIDStr := mux.Vars(r)["id"]
	mediaID, err := strconv.Atoi(mediaIDStr)
	if err != nil {
		Respond(w, r.Header.Get("hostname"), err, nil, true)
		return
	}

	mi := MonitoringInput{}
	if err := json.NewDecoder(r.Body).Decode(&mi); err != nil {
		Respond(w, r.Header.Get("hostname"), err, nil, true)
		return
	}

	err = dbstore.UpdateMediaMonitoring(mediaID, mi.Monitoring)
	Respond(w, r.Header.Get("hostname"), err, nil, true)
}

// func RefreshMediaMetadata(w http.ResponseWriter, r *http.Request) {
// 	mediaIDStr := mux.Vars(r)["id"]
// 	mediaID, err := strconv.Atoi(mediaIDStr)
// 	if err != nil {
// 		Respond(w, r.Header.Get("hostname"), err, nil, true)
// 		return
// 	}

// }
