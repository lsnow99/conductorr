package search

import (
	"net/http"
	"time"

	"github.com/lsnow99/conductorr/settings"
)

type SearchAPI interface {
	SearchByTitle(string, string, int) (*SearchResults, error)
	SearchByID(string) (*IndividualResult, error)
	SearchByImdbID(string) (*IndividualResult, error)
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var Client HTTPClient
var searchAPI SearchAPI

func init() {
	Client = &http.Client{}
}

func init() {
	// To check whether to use OMDB or TMDB, we do the following:
	// - Prefer OMDB over TMDB if the user supplied only the OMDB API key
	// - Prefer TMDB over OMDB in all other cases
	if settings.OmdbApiKey != settings.DefaultOMDBAPIKey && settings.TmdbAPIKey == settings.DefaultTmdbAPIKey {
		// TODO: set OMDB api instead
		searchAPI = &TmdbAPI{}
	} else {
		searchAPI = &TmdbAPI{}
	}
}

func SearchByTitle(title, contentType string, page int) (*SearchResults, error) {
	return searchAPI.SearchByTitle(title, contentType, page)
}

func GetResultByID(id string) (*IndividualResult, error) {
	return searchAPI.SearchByID(id)
}

func GetResultByImdbID(imdbID string) (*IndividualResult, error) {
	return searchAPI.SearchByImdbID(imdbID)
}

type SearchResult struct {
	Title       string
	Poster      string
	Plot        string
	Rating      float32
	ContentType string
	ReleasedAt  time.Time
	// The api agent used for the search
	Agent string
	// API agent's internal ID for the item
	ID string
}

type IndividualResult struct {
	SearchResult
	Runtime                int
	PhysicalDigitalRelease time.Time
	EndedAt                time.Time
	Genres                 []string
	ImdbID                 string
	TvdbID                 int
}

type SearchResults struct {
	Results      []SearchResult
	PerPage      int
	TotalResults int
}
