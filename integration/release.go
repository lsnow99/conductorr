package integration

import "time"

type Release struct {
	ID          string    `json:"id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	DownloadURL string    `json:"download_url,omitempty"`
	Categories  []string  `json:"categories,omitempty"`
	Size        int64     `json:"size,omitempty"`
	AirDate     time.Time `json:"air_date,omitempty"`
	PubDate     time.Time `json:"pub_date,omitempty"`
	Media       *Media    `json:"-"`
}

func NewRelease(id, title, description, downloadURL string, categories []string, size int64, airDate, pubDate time.Time, media *Media) Release {
	return Release{
		ID: id,
		Title: title,
		Description: description,
		DownloadURL: downloadURL,
		Categories: categories,
		Size: size,
		AirDate: airDate,
		PubDate: pubDate,
		Media: media,
	}
}