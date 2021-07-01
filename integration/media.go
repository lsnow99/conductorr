package integration

import "fmt"

const (
	Paused = iota
	Waiting
	Downloading
	Processing
	Error
	Done
)

const (
	Movie = iota
	TVShow
)

type Media struct {
	ID int
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
	// ParentMedia for tv shows, this is usually the media object that encompasses all episodes
	ParentMedia *Media
	// ImdbID assigned by Imdb
	ImdbID string
	// TvdbID assigned by Tvdb
	TvdbID int
	// Season (tv shows only)
	Season int
	// Episode (tv shows only)
	Episode int
	// Runtime minutes duration of media
	Runtime int
}

func (m Media) QueryString() string {
	var query string
	if m.ParentMedia != nil {
		query += m.ParentMedia.Title + " "
	}
	query += m.Title + " "

	if m.ContentType == TVShow {
		query += fmt.Sprintf("S%02dE%02d", m.Season, m.Episode)
	} else if m.ContentType == Movie {
		query += fmt.Sprintf("(%d)", m.Year)
	}
	return query
}