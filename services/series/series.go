package series

import "time"

type Episode struct {
	ImdbID string
	Title string
	Description string
	Runtime int
	Season int
	Episode int
	Aired time.Time
}

type SeriesAPI interface {
	GetEpisodes(imdbID string) ([]Episode, error)
}

var seriesAPI SeriesAPI

func init() {
	seriesAPI = &TVMaze{}
}

func GetEpisodes(imdbID string) ([]Episode, error) {
	return seriesAPI.GetEpisodes(imdbID)
}