package dbstore

import (
	"database/sql"
	"time"
)

type Profile struct {
	ID     int
	Name   sql.NullString
	Filter sql.NullString
	Sorter sql.NullString
}

type Media struct {
	ID            int
	Title         sql.NullString
	Description   sql.NullString
	ReleasedAt    sql.NullTime
	EndedAt       sql.NullTime
	ContentType   sql.NullString
	Poster        *[]byte
	ParentMediaID sql.NullInt32
	TmdbID        sql.NullInt32
	ImdbID        sql.NullString
	TvdbID		  sql.NullInt32
	TmdbRating    sql.NullInt32
	ImdbRating    sql.NullInt32
	Runtime       sql.NullInt32
	Status        string
	Path          sql.NullString
	Size          sql.NullInt64
	ProfileID     sql.NullInt32
	PathID        sql.NullInt32
	Number        sql.NullInt32
	Monitoring    bool
	Added		  time.Time
}

type Indexer struct {
	ID           int
	Name         string
	UserID       int
	BaseUrl      string
	ApiKey       string
	ForMovies    bool
	ForSeries    bool
	DownloadType string
	LastRSSID    string
}

type Downloader struct {
	ID             int
	Name           string
	DownloaderType string
	FileAction string
	Config         map[string]interface{}
}

type MediaServer struct {
	ID              int
	Name            string
	MediaServerType string
	Config          map[string]interface{}
}

type Download struct {
	ID           int
	MediaID      sql.NullInt32
	DownloaderID sql.NullInt32
	Status       string
	FriendlyName string
	Identifier   string
	ReleaseID    sql.NullString
}

type ReleaseHistory struct {
	ID string
	MediaID int
	LastAttempt time.Time
}

type Path struct {
	ID            int
	Path          string
	MoviesDefault bool
	SeriesDefault bool
}
