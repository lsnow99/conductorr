package integration

import "fmt"

const (
	Movie = iota
	TVSeries
	TVSeason
	TVEpisode
)

type Media struct {
	ID int
	PathID int
	// Title of the content
	Title string
	// Year of the content
	Year int
	// Description of the content
	Description string
	// Category the category to be passed to the downloader
	Category string
	// HighPriority whether to assign high priority to the media
	HighPriority bool
	// URL fetched by the indexer
	URL string
	// ContentType either Movie or TVShow
	ContentType int
	// ParentMedia for tv shows, can either be a season media type or series
	ParentMedia *Media
	// ImdbID assigned by Imdb
	ImdbID string
	// TvdbID assigned by Tvdb
	TvdbID int
	// Episode (tv episodes and series only)
	Number int
	// Runtime minutes duration of media
	Runtime int64
}

func (m Media) QueryString() string {

	if m.ContentType == TVEpisode {
		if m.ParentMedia == nil || m.ParentMedia.ParentMedia == nil {
			return ""
		}
		return fmt.Sprintf("%s S%02dE%02d", m.ParentMedia.ParentMedia.Title, m.ParentMedia.Number, m.Number)
	} else if m.ContentType == Movie {
		return fmt.Sprintf("%s %d", m.Title, m.Year)
	} else if m.ContentType == TVSeason {
		if m.ParentMedia == nil {
			return ""
		}
		return fmt.Sprintf("%s Season %d", m.ParentMedia.Title, m.Number)
	} else if m.ContentType == TVSeries {
		return fmt.Sprintf("%s %d", m.Title, m.Year)
	}
	return ""
}