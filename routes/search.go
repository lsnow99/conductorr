package routes

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/lsnow99/conductorr/dbstore"
	"github.com/lsnow99/conductorr/services/omdb"
)

type MediaResponse struct {
	ID            int       `json:"id,omitempty"`
	Title         string    `json:"title,omitempty"`
	Description   string    `json:"description,omitempty"`
	ReleasedAt    time.Time `json:"released_at,omitempty"`
	EndedAt       *time.Time `json:"ended_at,omitempty"`
	ContentType   string    `json:"content_type,omitempty"`
	Poster        string    `json:"poster,omitempty"`
	ParentMediaID int       `json:"parent_media_id,omitempty"`
	TmdbID        int       `json:"tmdb_id,omitempty"`
	ImdbID        string    `json:"imdb_id,omitempty"`
	TmdbRating    int       `json:"tmdb_rating,omitempty"`
	ImdbRating    int       `json:"imdb_rating,omitempty"`
	Runtime       int       `json:"runtime,omitempty"`

	InLibrary bool `json:"in_library,omitempty"`
}

type SearchResponse struct {
	TotalResults int              `json:"total_results"`
	Results      []MediaResponse `json:"results"`
}

func NewMediaResponseFromDB(media dbstore.Media) (m MediaResponse) {
	m.ID = media.ID
	if media.Title.Valid {
		m.Title = media.Title.String
	}
	if media.Description.Valid {
		m.Description = media.Description.String
	}
	if media.ReleasedAt.Valid {
		m.ReleasedAt = media.ReleasedAt.Time
	}
	if media.EndedAt.Valid {
		m.EndedAt = &media.EndedAt.Time
	}
	if media.ContentType.Valid {
		m.ContentType = media.ContentType.String
	}
	m.Poster = "/api/v1/poster/" + strconv.Itoa(m.ID)
	if media.ParentMediaID.Valid {
		m.ParentMediaID = int(media.ParentMediaID.Int32)
	}
	if media.TmdbID.Valid {
		m.TmdbID = int(media.TmdbID.Int32)
	}
	if media.ImdbID.Valid {
		m.ImdbID = media.ImdbID.String
	}
	if media.TmdbRating.Valid {
		m.TmdbRating = int(media.TmdbRating.Int32)
	}
	if media.ImdbRating.Valid {
		m.ImdbRating = int(media.ImdbRating.Int32)
	}
	return m
}

func NewMediaResponseFromOmdbIndividual(media omdb.IndividualResult) (m MediaResponse) {
	m.Title = media.Title
	m.Description = media.Plot
	m.ReleasedAt = media.ReleasedAt
	m.EndedAt = &media.EndedAt
	m.ContentType = media.Type
	m.Poster = media.Poster
	m.ImdbID = media.ImdbID
	m.ImdbRating = int(media.ImdbRating * 10)
	return m
}

func NewMediaResponseFromOmdbSearch(media omdb.SearchResult) (m MediaResponse) {
	m.Title = media.Title
	m.ReleasedAt, _ = time.Parse("2006", strconv.Itoa(media.YearStart))
	if media.YearEnd != -1 {
		ended, _ := time.Parse("2006", strconv.Itoa(media.YearEnd))
		m.EndedAt = &ended
	}
	m.ImdbID = media.ImdbID
	m.ContentType = media.Type
	m.Poster = media.Poster
	return m
}

func SearchNewByTitle(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	query := r.URL.Query().Get("q")
	contentType := r.URL.Query().Get("type")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		Respond(w, r.Host, err, nil, true)
		return
	}

	omdbResults, err := omdb.SearchByTitle(query, contentType, page)
	if err != nil {
		Respond(w, r.Host, err, nil, true)
		return
	}

	medias := make([]MediaResponse, 0)
	for _, omdbResult := range omdbResults.Search {
		var media MediaResponse
		dbMedia, err := dbstore.GetMediaByImdbID(omdbResult.ImdbID)
		if err != nil && err != sql.ErrNoRows {
			Respond(w, r.Host, err, nil, true)
			return
		} else if err == sql.ErrNoRows {
			media = NewMediaResponseFromOmdbSearch(omdbResult)
			media.InLibrary = false
		} else {
			media = NewMediaResponseFromDB(dbMedia)
			media.InLibrary = true
		}
		medias = append(medias, media)
	}

	sr := SearchResponse{
		TotalResults: omdbResults.TotalResults,
		Results: medias,
	}

	Respond(w, r.Host, err, sr, true)
}

func SearchLibraryByTitle(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	query := r.URL.Query().Get("q")
	contentType := r.URL.Query().Get("type")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		Respond(w, r.Host, err, nil, true)
		return
	}

	// Search our own library
	medias, total, err := dbstore.SearchMedia(query, contentType, page)
	
	results := make([]MediaResponse, 0)
	for _, media := range medias {
		if media == nil {
			continue
		}
		results = append(results, NewMediaResponseFromDB(*media))
	}

	resp := SearchResponse{
		TotalResults: total,
		Results:      results,
	}

	Respond(w, r.Host, err, resp, true)
}
