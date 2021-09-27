package search

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/lsnow99/conductorr/logger"
	"github.com/lsnow99/conductorr/settings"
)

type TmdbAPI struct{}

type TmdbTitleSearch struct {
	Page    int `json:"page"`
	Results []struct {
		Adult            bool     `json:"adult,omitempty"`
		BackdropPath     string   `json:"backdrop_path"`
		GenreIds         []int    `json:"genre_ids"`
		ID               int      `json:"id"`
		MediaType        string   `json:"media_type"`
		OriginalLanguage string   `json:"original_language"`
		OriginalTitle    string   `json:"original_title,omitempty"`
		Overview         string   `json:"overview"`
		Popularity       float64  `json:"popularity"`
		PosterPath       string   `json:"poster_path"`
		ReleaseDate      string   `json:"release_date,omitempty"`
		Title            string   `json:"title,omitempty"`
		Video            bool     `json:"video,omitempty"`
		VoteAverage      float64  `json:"vote_average"`
		VoteCount        int      `json:"vote_count"`
		FirstAirDate     string   `json:"first_air_date,omitempty"`
		Name             string   `json:"name,omitempty"`
		OriginCountry    []string `json:"origin_country,omitempty"`
		OriginalName     string   `json:"original_name,omitempty"`
	} `json:"results"`
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
}

type TmdbIndividualSearch struct {
	Adult               bool   `json:"adult,omitempty"`
	BackdropPath        string `json:"backdrop_path"`
	BelongsToCollection struct {
		ID           int    `json:"id"`
		Name         string `json:"name"`
		PosterPath   string `json:"poster_path"`
		BackdropPath string `json:"backdrop_path"`
	} `json:"belongs_to_collection,omitempty"`
	Budget int `json:"budget,omitempty"`
	Genres []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"genres"`
	Homepage            string  `json:"homepage"`
	ID                  int     `json:"id"`
	ImdbID              string  `json:"imdb_id,omitempty"`
	OriginalLanguage    string  `json:"original_language"`
	OriginalTitle       string  `json:"original_title,omitempty"`
	Overview            string  `json:"overview"`
	Popularity          float64 `json:"popularity"`
	PosterPath          string  `json:"poster_path"`
	ProductionCompanies []struct {
		ID            int    `json:"id"`
		LogoPath      string `json:"logo_path"`
		Name          string `json:"name"`
		OriginCountry string `json:"origin_country"`
	} `json:"production_companies"`
	ProductionCountries []struct {
		Iso31661 string `json:"iso_3166_1"`
		Name     string `json:"name"`
	} `json:"production_countries"`
	ReleaseDate     string `json:"release_date,omitempty"`
	Revenue         int    `json:"revenue,omitempty"`
	Runtime         int    `json:"runtime,omitempty"`
	SpokenLanguages []struct {
		EnglishName string `json:"english_name"`
		Iso6391     string `json:"iso_639_1"`
		Name        string `json:"name"`
	} `json:"spoken_languages"`
	Status           string        `json:"status"`
	Tagline          string        `json:"tagline"`
	Title            string        `json:"title,omitempty"`
	Video            bool          `json:"video,omitempty"`
	VoteAverage      float64       `json:"vote_average"`
	VoteCount        int           `json:"vote_count"`
	CreatedBy        []interface{} `json:"created_by,omitempty"`
	EpisodeRunTime   []int         `json:"episode_run_time,omitempty"`
	FirstAirDate     string        `json:"first_air_date,omitempty"`
	InProduction     bool          `json:"in_production,omitempty"`
	Languages        []string      `json:"languages,omitempty"`
	LastAirDate      string        `json:"last_air_date,omitempty"`
	LastEpisodeToAir struct {
		AirDate        string  `json:"air_date"`
		EpisodeNumber  int     `json:"episode_number"`
		ID             int     `json:"id"`
		Name           string  `json:"name"`
		Overview       string  `json:"overview"`
		ProductionCode string  `json:"production_code"`
		SeasonNumber   int     `json:"season_number"`
		StillPath      string  `json:"still_path"`
		VoteAverage    float64 `json:"vote_average"`
		VoteCount      int     `json:"vote_count"`
	} `json:"last_episode_to_air,omitempty"`
	Name             string      `json:"name,omitempty"`
	NextEpisodeToAir interface{} `json:"next_episode_to_air,omitempty"`
	Networks         []struct {
		Name          string `json:"name"`
		ID            int    `json:"id"`
		LogoPath      string `json:"logo_path"`
		OriginCountry string `json:"origin_country"`
	} `json:"networks,omitempty"`
	NumberOfEpisodes int      `json:"number_of_episodes,omitempty"`
	NumberOfSeasons  int      `json:"number_of_seasons,omitempty"`
	OriginCountry    []string `json:"origin_country,omitempty"`
	OriginalName     string   `json:"original_name,omitempty"`
	Seasons          []struct {
		AirDate      string `json:"air_date"`
		EpisodeCount int    `json:"episode_count"`
		ID           int    `json:"id"`
		Name         string `json:"name"`
		Overview     string `json:"overview"`
		PosterPath   string `json:"poster_path"`
		SeasonNumber int    `json:"season_number"`
	} `json:"seasons,omitempty"`
	Type         string `json:"type,omitempty"`
	ReleaseDates struct {
		Results []struct {
			Iso31661     string `json:"iso_3166_1"`
			ReleaseDates []struct {
				Certification string    `json:"certification"`
				Iso6391       string    `json:"iso_639_1"`
				Note          string    `json:"note"`
				ReleaseDate   time.Time `json:"release_date"`
				Type          int       `json:"type"`
			} `json:"release_dates"`
		} `json:"results"`
	} `json:"release_dates,omitempty"`
	ExternalIds struct {
		ImdbID      string      `json:"imdb_id"`
		FreebaseMid interface{} `json:"freebase_mid"`
		FreebaseID  interface{} `json:"freebase_id"`
		TvdbID      int         `json:"tvdb_id"`
		TvrageID    interface{} `json:"tvrage_id"`
		FacebookID  string      `json:"facebook_id"`
		InstagramID string      `json:"instagram_id"`
		TwitterID   string      `json:"twitter_id"`
	} `json:"external_ids"`
	Success       bool   `json:"success"`
	StatusCode    int    `json:"status_code,omitempty"`
	StatusMessage string `json:"status_message,omitempty"`
}

type TmdbImdbResult struct {
	MovieResults []struct {
		GenreIds         []int   `json:"genre_ids"`
		OriginalLanguage string  `json:"original_language"`
		OriginalTitle    string  `json:"original_title"`
		PosterPath       string  `json:"poster_path"`
		Video            bool    `json:"video"`
		ID               int     `json:"id"`
		VoteCount        int     `json:"vote_count"`
		Overview         string  `json:"overview"`
		ReleaseDate      string  `json:"release_date"`
		VoteAverage      float64 `json:"vote_average"`
		Title            string  `json:"title"`
		Adult            bool    `json:"adult"`
		BackdropPath     string  `json:"backdrop_path"`
		Popularity       float64 `json:"popularity"`
	} `json:"movie_results"`
	TvResults []struct {
		OriginalName     string   `json:"original_name"`
		GenreIds         []int    `json:"genre_ids"`
		OriginalLanguage string   `json:"original_language"`
		ID               int      `json:"id"`
		PosterPath       string   `json:"poster_path"`
		VoteAverage      float64  `json:"vote_average"`
		Overview         string   `json:"overview"`
		VoteCount        int      `json:"vote_count"`
		Name             string   `json:"name"`
		BackdropPath     string   `json:"backdrop_path"`
		OriginCountry    []string `json:"origin_country"`
		FirstAirDate     string   `json:"first_air_date"`
		Popularity       float64  `json:"popularity"`
	} `json:"tv_results"`
}

type tmdbConfigObj struct {
	Images struct {
		BaseURL       string   `json:"base_url"`
		SecureBaseURL string   `json:"secure_base_url"`
		BackdropSizes []string `json:"backdrop_sizes"`
		LogoSizes     []string `json:"logo_sizes"`
		PosterSizes   []string `json:"poster_sizes"`
		ProfileSizes  []string `json:"profile_sizes"`
		StillSizes    []string `json:"still_sizes"`
	} `json:"images"`
	ChangeKeys []string `json:"change_keys"`
}

type tmdbConfiguration struct {
	Config      tmdbConfigObj
	LastUpdated time.Time
	sync.RWMutex
}

const tmdbApihost = "api.themoviedb.org"
const dateFormat = "2006-01-02"

var idPattern = regexp.MustCompile("^(tv|movie):(.{1,})")

var tmdbConfig tmdbConfiguration

func fetchTmdbConfig() (tmdbConfigObj, error) {
	var lastUpdated time.Time
	tmdbConfig.RLock()
	lastUpdated = tmdbConfig.LastUpdated
	tmdbConfig.RUnlock()

	// Cache for two days, then update
	if time.Since(lastUpdated).Hours() > 48 {
		u := url.URL{}
		u.Scheme = "https"
		u.Host = tmdbApihost
		u.Path = "/3/configuration"
		q := u.Query()
		q.Set("api_key", settings.TmdbAPIKey)
		u.RawQuery = q.Encode()

		req, err := http.NewRequest("GET", u.String(), nil)
		if err != nil {
			return tmdbConfigObj{}, err
		}

		resp, err := Client.Do(req)
		if err != nil {
			return tmdbConfigObj{}, err
		}

		cfgObj := tmdbConfigObj{}
		err = json.NewDecoder(resp.Body).Decode(&cfgObj)
		if err != nil {
			return tmdbConfigObj{}, err
		}

		tmdbConfig.Lock()
		tmdbConfig.Config = cfgObj
		tmdbConfig.Unlock()
	}

	tmdbConfig.RLock()
	defer tmdbConfig.RUnlock()
	return tmdbConfig.Config, nil
}

func (t *TmdbAPI) SearchByTitle(title, contentType string, page int) (*SearchResults, error) {
	if title == "" {
		return &SearchResults{}, nil
	}

	config, err := fetchTmdbConfig()
	if err != nil {
		return nil, err
	}

	u := url.URL{}
	u.Scheme = "https"
	u.Host = tmdbApihost
	u.Path = "/3/search/"
	if contentType == "series" {
		u.Path += "tv"
	} else if contentType == "movie" {
		u.Path += "movie"
	} else {
		u.Path += "multi"
	}
	q := u.Query()
	q.Set("api_key", settings.TmdbAPIKey)
	q.Set("query", title)
	q.Set("include_adult", "true")
	q.Set("page", strconv.Itoa(page))
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := Client.Do(req)
	if err != nil {
		return nil, err
	}

	results := &TmdbTitleSearch{}
	err = json.NewDecoder(resp.Body).Decode(results)
	if err != nil {
		return nil, err
	}

	searchResults := SearchResults{}
	searchResults.TotalResults = results.TotalResults
	searchResults.PerPage = 20
	for _, result := range results.Results {
		sr := SearchResult{}
		sr.Poster = getPoster(config, result.PosterPath)
		sr.Agent = "tmdb"
		sr.Rating = float32(result.VoteAverage)
		sr.Plot = result.Overview
		if result.MediaType == "tv" {
			sr.ID = "tv:" + strconv.Itoa(result.ID)
			sr.Title = result.Name
			sr.ContentType = "series"
			sr.ReleasedAt, err = time.Parse(dateFormat, result.FirstAirDate)
			if err != nil {
				logger.LogWarn(err)
			}
		} else if result.MediaType == "movie" {
			sr.ID = "movie:" + strconv.Itoa(result.ID)
			sr.Title = result.Title
			sr.ContentType = "movie"
			sr.ReleasedAt, err = time.Parse(dateFormat, result.ReleaseDate)
			if err != nil {
				logger.LogWarn(err)
			}
		} else {
			continue
		}
		searchResults.Results = append(searchResults.Results, sr)
	}
	return &searchResults, nil
}

func (t *TmdbAPI) SearchByID(id string) (*IndividualResult, error) {
	matches := idPattern.FindAllStringSubmatch(id, -1)
	if len(matches) != 1 {
		return nil, fmt.Errorf("badly formatted id %s", id)
	}
	if len(matches[0]) != 3 {
		return nil, fmt.Errorf("regexp failed for id %s", id)
	}

	config, err := fetchTmdbConfig()
	if err != nil {
		return nil, err
	}

	mediaType := matches[0][1]
	idStr := matches[0][2]

	u := url.URL{}
	u.Scheme = "https"
	u.Host = tmdbApihost
	u.Path = "/3/" + mediaType + "/" + idStr
	q := u.Query()
	q.Set("append_to_response", "release_dates,external_ids")
	q.Set("api_key", settings.TmdbAPIKey)
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := Client.Do(req)
	if err != nil {
		return nil, err
	}

	result := &TmdbIndividualSearch{}
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return nil, err
	}

	ir := &IndividualResult{}
	ir.Agent = "tmdb"
	ir.Poster = getPoster(config, result.PosterPath)
	ir.Rating = float32(result.VoteAverage)
	ir.Plot = result.Overview

	for _, genre := range result.Genres {
		ir.Genres = append(ir.Genres, genre.Name)
	}

	if mediaType == "tv" {
		if len(result.EpisodeRunTime) < 1 {
			return nil, fmt.Errorf("could not determine runtime for item id %s", id)
		}
		ir.Runtime = result.EpisodeRunTime[0]
		ir.Title = result.Name
		ir.ContentType = "series"
		ir.ImdbID = result.ExternalIds.ImdbID
		ir.ReleasedAt, err = time.Parse(dateFormat, result.FirstAirDate)
		if err != nil {
			logger.LogWarn(err)
		}
		if result.LastAirDate != "" {
			ir.EndedAt, err = time.Parse(dateFormat, result.LastAirDate)
			if err != nil {
				return nil, fmt.Errorf("bad date %s for item id %s: %s", result.LastAirDate, id, err)
			}
		}
	} else if mediaType == "movie" {
		ir.Runtime = result.Runtime
		ir.Title = result.Title
		ir.ContentType = "movie"
		ir.ImdbID = result.ImdbID
		ir.ReleasedAt, err = time.Parse(dateFormat, result.ReleaseDate)
		if err != nil {
			logger.LogWarn(err)
		}
		// Get the soonest physical or dvd release date
		for _, res := range result.ReleaseDates.Results {
			// US only for now
			if res.Iso31661 == "US" {
				for _, releaseDate := range res.ReleaseDates {
					// 4 is digital 5 is physical
					if releaseDate.Type == 4 || releaseDate.Type == 5 {
						// If we haven't set a date yet or if the one we have now is sooner
						if ir.PhysicalDigitalRelease.IsZero() || releaseDate.ReleaseDate.Before(ir.PhysicalDigitalRelease) {
							ir.PhysicalDigitalRelease = releaseDate.ReleaseDate
						}
					}
				}
			}
		}
	}

	return ir, nil
}

func (t *TmdbAPI) SearchByImdbID(imdbID string) (*IndividualResult, error) {
	u := url.URL{}
	u.Scheme = "https"
	u.Host = tmdbApihost
	u.Path = "/3/find"
	q := u.Query()
	q.Set("append_to_response", "release_dates")
	q.Set("api_key", settings.TmdbAPIKey)
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := Client.Do(req)
	if err != nil {
		return nil, err
	}

	result := &TmdbImdbResult{}
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return nil, err
	}

	if len(result.TvResults) > 0 {
		return t.SearchByID("tv:" + strconv.Itoa(result.TvResults[0].ID))
	} else if len(result.MovieResults) > 0 {
		return t.SearchByID("movie:" + strconv.Itoa(result.MovieResults[0].ID))
	}

	return nil, fmt.Errorf("no results found for imdb id %s", imdbID)
}

func getPoster(config tmdbConfigObj, posterPath string) string {
	var bestSize string
	if len(config.Images.PosterSizes) > 0 {
		bestSize = config.Images.PosterSizes[0]
	}
	for _, size := range config.Images.PosterSizes {
		if strings.Contains(size, "original") {
			bestSize = size
		}
	}
	for _, size := range config.Images.PosterSizes {
		if strings.Contains(size, "780") {
			bestSize = size
		}
	}
	return config.Images.SecureBaseURL + path.Join(bestSize, posterPath)
}
