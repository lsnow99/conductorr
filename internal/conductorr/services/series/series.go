package series

import (
	"fmt"
	"time"
)

type Episode struct {
	ImdbID      string
	Title       string
	Description string
	Runtime     int
	Season      int
	Episode     int
	Aired       time.Time
}

type SeriesAPI interface {
	GetEpisodes(imdbID string) ([]Episode, error)
}

type ErrImdbIDNotFound struct {
	ImdbID string
	Agent  string
}

var seriesAPI SeriesAPI

func init() {
	seriesAPI = &TVMaze{}
}

func GetEpisodes(imdbID string) ([]Episode, error) {
	return seriesAPI.GetEpisodes(imdbID)
}

func (e ErrImdbIDNotFound) Error() string {
	return fmt.Sprintf("no results found for IMDB ID %s on %s", e.ImdbID, e.Agent)
}
