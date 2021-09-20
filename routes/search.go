package routes

import (
	"errors"
	"net/http"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/lsnow99/conductorr/dbstore"
	"github.com/lsnow99/conductorr/services/search"
)

type MediaResponse struct {
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
	Number        int        `json:"number,omitempty"`
	Monitoring    bool       `json:"monitoring"`
	Path          string     `json:"path"`
	PathOK        bool       `json:"path_ok"`
	Size          int64      `json:"size"`

	SearchID string `json:"search_id,omitempty"`

	Children []MediaResponse `json:"children,omitempty"`

	// to prevent an infinite loop in Expand()
	visited map[int]bool `json:"-"`
}

// Expand a media response with all the recursively referenced children
func (mr *MediaResponse) Expand() error {
	if mr.visited == nil {
		mr.visited = make(map[int]bool)
	}
	ok, size := CheckMediaPath(mr.Path)
	mr.PathOK = ok
	mr.Size = size
	mr.visited[mr.ID] = true
	children, err := dbstore.GetMediaReferencing(mr.ID)
	if err != nil {
		return err
	}
	medias := make([]MediaResponse, 0)
	for _, child := range children {
		if _, ok := mr.visited[child.ID]; ok {
			return errors.New("self referential media loop")
		}
		if child == nil {
			continue
		}
		media := NewMediaResponseFromDB(*child)
		medias = append(medias, media)
	}
	for _, media := range medias {
		media.visited = mr.visited
		err = media.Expand()
		if err != nil {
			return err
		}
		mr.Children = append(mr.Children, media)
	}
	sort.Slice(mr.Children, func(i, j int) bool {
		return mr.Children[i].Number < mr.Children[j].Number
	})
	return nil
}

func CheckMediaPath(path string) (bool, int64) {
	fi, err := os.Stat(path)
	if err != nil {
		return false, 0
	}
	return true, fi.Size()
}

type SearchResponse struct {
	TotalResults int             `json:"total_results"`
	PerPage      int             `json:"per_page"`
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
	if media.ProfileID.Valid {
		m.ProfileID = int(media.ProfileID.Int32)
	}
	if media.PathID.Valid {
		m.PathID = int(media.PathID.Int32)
	}
	if media.Number.Valid {
		m.Number = int(media.Number.Int32)
	}
	if media.Path.Valid {
		m.Path = media.Path.String
	}
	m.Monitoring = media.Monitoring
	return m
}

func NewMediaResponseFromIndividual(media search.IndividualResult) (m MediaResponse) {
	m.Title = media.Title
	m.Description = media.Plot
	m.ReleasedAt = media.ReleasedAt
	m.EndedAt = &media.EndedAt
	m.ContentType = media.ContentType
	m.Poster = media.Poster
	m.ImdbID = media.ImdbID
	m.ImdbRating = int(media.Rating * 10)
	return m
}

func NewMediaResponseFromSearch(media search.SearchResult) (m MediaResponse) {
	m.Title = media.Title
	m.ReleasedAt = media.ReleasedAt
	m.SearchID = media.ID
	m.ContentType = media.ContentType
	m.Poster = media.Poster
	return m
}

func SearchNewByTitle(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	query := r.URL.Query().Get("q")
	contentType := r.URL.Query().Get("type")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		Respond(w, r.Header.Get("hostname"), err, nil, true)
		return
	}

	searchResults, err := search.SearchByTitle(query, contentType, page)
	if err != nil {
		Respond(w, r.Header.Get("hostname"), err, nil, true)
		return
	}

	medias := make([]MediaResponse, 0)
	for _, searchResult := range searchResults.Results {
		medias = append(medias, NewMediaResponseFromSearch(searchResult))
	}

	sr := SearchResponse{
		TotalResults: searchResults.TotalResults,
		PerPage:      searchResults.PerPage,
		Results:      medias,
	}

	Respond(w, r.Header.Get("hostname"), err, sr, true)
}

func SearchLibraryByTitle(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	query := r.URL.Query().Get("q")
	contentType := r.URL.Query().Get("type")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		Respond(w, r.Header.Get("hostname"), err, nil, true)
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
		PerPage:      10,
		Results:      results,
	}

	Respond(w, r.Header.Get("hostname"), err, resp, true)
}
