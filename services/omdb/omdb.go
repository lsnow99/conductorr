package omdb

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"time"

	"github.com/lsnow99/conductorr/settings"
)

const apiHost = "www.omdbapi.com"

var yearPattern = regexp.MustCompile(`^(\d{4})(\-|–)?(\d{4})?`)
var runtimePattern = regexp.MustCompile(`^([0-9]+) min`)

type SearchResults struct {
	Search          []SearchResult
	TotalResults    int
	TotalResultsStr string `json:"totalResults"`
	Response        string
	Error           string
}

type SearchResult struct {
	Title     string
	YearStart int
	YearEnd   int
	YearStr   string `json:"year"`
	ImdbID    string `json:"imdbID"`
	Type      string
	Poster    string
}

type IndividualResult struct {
	SearchResult
	Rated      string
	Released   string
	ReleasedAt time.Time
	EndedAt    time.Time
	Runtime    int
	RuntimeStr string `json:"Runtime"`
	Genre      string
	Director   string
	Writer     string
	Actors     string
	Plot       string
	Language   string
	Country    string
	Awards     string
	Ratings    []struct {
		Source string
		Value  string
	}
	Metascore     int
	MetascoreStr  string `json:"Metascore"`
	ImdbRating    float32
	ImdbRatingStr string `json:"imdbRating"`
	ImdbVotes     string `json:"imdbVotes"`
	DVD           string
	BoxOffice     string
	Production    string
	Website       string
	Response      string
	Error         string
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var Client HTTPClient

func init() {
	Client = &http.Client{}
}

func SearchByTitle(title string, contentType string, page int) (*SearchResults, error) {
	if title == "" {
		return &SearchResults{}, nil
	}

	u := url.URL{}
	u.Scheme = "https"
	u.Host = apiHost
	u.Path = "/"
	q := u.Query()
	q.Set("s", title)
	q.Set("page", strconv.Itoa(page))
	q.Set("apikey", settings.OmdbApiKey)
	q.Set("plot", "full")
	q.Set("r", "json")
	if contentType == "movie" {
		q.Set("type", "movie")
	} else if contentType == "series" {
		q.Set("type", "series")
	} // else include all content types
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := Client.Do(req)
	if err != nil {
		return nil, err
	}

	results := &SearchResults{}
	err = json.NewDecoder(resp.Body).Decode(results)
	if err != nil {
		return nil, err
	}

	if results.Response != "True" {
		return nil, errors.New(results.Error)
	}

	results.TotalResults, err = strconv.Atoi(results.TotalResultsStr)
	if err != nil {
		return nil, err
	}

	for i := range results.Search {
		if err := results.Search[i].Sanitize(); err != nil {
			return nil, err
		}
	}
	return results, nil
}

func SearchByImdbID(imdbID string) (*IndividualResult, error) {
	u := url.URL{}
	u.Scheme = "https"
	u.Host = apiHost
	u.Path = "/"
	q := u.Query()
	q.Set("i", imdbID)
	q.Set("apikey", settings.OmdbApiKey)
	q.Set("r", "json")
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := Client.Do(req)
	if err != nil {
		return nil, err
	}

	result := &IndividualResult{}
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return nil, err
	}

	if result.Response != "True" {
		return nil, errors.New(result.Error)
	}

	return result, result.Sanitize()
}

func (ir *IndividualResult) Sanitize() error {
	// Sanitize the underlying SearchResult
	if err := ir.SearchResult.Sanitize(); err != nil {
		return err
	}

	/*
		Parse the runtime into a number of minutes
	*/
	if ir.RuntimeStr != "N/A" {
		matches := runtimePattern.FindAllStringSubmatch(ir.RuntimeStr, -1)
		if len(matches) != 1 {
			return fmt.Errorf("got %d matches expected 1", len(matches))
		}
		if len(matches[0]) != 2 {
			return fmt.Errorf("got %d groups expected 2", len(matches[0]))
		}
		minsStr := matches[0][1]
		mins, err := strconv.Atoi(minsStr)
		if err != nil {
			return err
		}
		ir.Runtime = mins
	}

	/*
		Parse the released string into a date
	*/
	releasedTime, err := time.Parse("02 Jan 2006", ir.Released)
	if err != nil {
		return err
	}
	ir.ReleasedAt = releasedTime

	if ir.YearEnd != -1 {
		if ir.YearEnd == ir.YearStart {
			ir.EndedAt = ir.ReleasedAt
		} else {
			ir.EndedAt, _ = time.Parse("2006", strconv.Itoa(ir.YearEnd))
		}
	}

	/*
		Parse the imdb rating into a float
	*/
	imdbRating, err := strconv.ParseFloat(ir.ImdbRatingStr, 32)
	if err != nil {
		imdbRating = 0
		// return err
	}
	ir.ImdbRating = float32(imdbRating)

	return nil
}

func (sr *SearchResult) Sanitize() error {
	/*
		Parse the year start and year end
	*/
	var yearStart, yearEnd int
	matches := yearPattern.FindAllStringSubmatch(sr.YearStr, -1)
	if len(matches) != 1 {
		return fmt.Errorf("got matches length of %d expected 1", len(matches))
	}
	if len(matches[0]) != 4 {
		return fmt.Errorf("got group length of %d expected 4", len(matches[0]))
	}

	yearStartMatch := matches[0][1]
	yearDashMatch := matches[0][2]
	yearEndMatch := matches[0][3]
	if len(yearStartMatch) != 4 {
		return fmt.Errorf("got %d digits in year expected 4", len(yearStartMatch))
	}

	yearStart, err := strconv.Atoi(yearStartMatch)
	if err != nil {
		return err
	}

	if len(yearEndMatch) == 4 {
		yearEnd, err = strconv.Atoi(yearEndMatch)
		if err != nil {
			return err
		}
	} else {
		if yearDashMatch == "-" || yearDashMatch == "–" {
			yearEnd = -1
		} else {
			yearEnd = yearStart
		}
	}

	sr.YearStart = yearStart
	sr.YearEnd = yearEnd

	return nil
}
