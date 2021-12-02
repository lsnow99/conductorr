package series

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type TVMaze struct {}

type EpisodeResults []EpisodeResult

type EpisodeResult struct {
	ID       int       `json:"id"`
	URL      string    `json:"url"`
	Name     string    `json:"name"`
	Season   int       `json:"season"`
	Number   int       `json:"number"`
	Type     string    `json:"type"`
	Airdate  string    `json:"airdate"`
	Airtime  string    `json:"airtime"`
	Airstamp time.Time `json:"airstamp"`
	Runtime  int       `json:"runtime"`
	Image    struct {
		Medium   string `json:"medium"`
		Original string `json:"original"`
	} `json:"image"`
	Summary string `json:"summary"`
	Links   struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
}

type TVMazeShowResult struct {
	ID             int      `json:"id"`
	URL            string   `json:"url"`
	Name           string   `json:"name"`
	Type           string   `json:"type"`
	Language       string   `json:"language"`
	Genres         []string `json:"genres"`
	Status         string   `json:"status"`
	Runtime        int      `json:"runtime"`
	AverageRuntime int      `json:"averageRuntime"`
	Premiered      string   `json:"premiered"`
	OfficialSite   string   `json:"officialSite"`
	Schedule       struct {
		Time string   `json:"time"`
		Days []string `json:"days"`
	} `json:"schedule"`
	Rating struct {
		Average float64 `json:"average"`
	} `json:"rating"`
	Weight  int `json:"weight"`
	Network struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Country struct {
			Name     string `json:"name"`
			Code     string `json:"code"`
			Timezone string `json:"timezone"`
		} `json:"country"`
	} `json:"network"`
	WebChannel interface{} `json:"webChannel"`
	DvdCountry interface{} `json:"dvdCountry"`
	Externals  struct {
		Tvrage  int    `json:"tvrage"`
		Thetvdb int    `json:"thetvdb"`
		Imdb    string `json:"imdb"`
	} `json:"externals"`
	Image struct {
		Medium   string `json:"medium"`
		Original string `json:"original"`
	} `json:"image"`
	Summary string `json:"summary"`
	Updated int    `json:"updated"`
	Links   struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Previousepisode struct {
			Href string `json:"href"`
		} `json:"previousepisode"`
	} `json:"_links"`
}

func (t *TVMaze) GetEpisodes(imdbID string) ([]Episode, error) {
	id, err := t.getShowID(imdbID)
	if err != nil {
		return nil, err
	}

	u := url.URL{}
	u.Scheme = "https"
	u.Host = "api.tvmaze.com"
	u.Path = "shows/" + strconv.Itoa(id) + "/episodes"

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	results := EpisodeResults{}
	err = json.NewDecoder(resp.Body).Decode(&results)
	if err != nil {
		return nil, err
	}

	episodes := make([]Episode, len(results))

	for i, result := range results {
		episodes[i] = Episode{
			Title:       result.Name,
			Season:      result.Season,
			Episode:     result.Number,
			Aired:       result.Airstamp,
			Description: result.Summary,
		}
	}

	return episodes, nil
}

func (t *TVMaze) getShowID(imdbID string) (int, error) {
	u := url.URL{}
	u.Scheme = "https"
	u.Host = "api.tvmaze.com"
	u.Path = "lookup/shows"
	q := u.Query()
	q.Set("imdb", imdbID)
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return 0, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return 0, &ErrImdbIDNotFound{
			ImdbID: imdbID,
			Agent: "tvmaze",
		}
	}

	sr := TVMazeShowResult{}
	if err := json.NewDecoder(resp.Body).Decode(&sr); err != nil {
		return 0, err
	}
	return sr.ID, nil
}
