package dbstore

import (
	"database/sql"
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
	TmdbRating    sql.NullInt32
	ImdbRating    sql.NullInt32
	Runtime       sql.NullInt32
	Status        string
	Path          sql.NullString
	Size          sql.NullInt64
	ProfileID	  sql.NullInt32
}

type Indexer struct {
	ID      int
	Name    string
	UserID sql.NullInt32
	BaseUrl string
	ApiKey  string
	ForMovies bool
	ForSeries bool
	DownloadType string
}

type Downloader struct {
	ID int
	Name string
	DownloaderType string
	Config map[string]interface{}
}