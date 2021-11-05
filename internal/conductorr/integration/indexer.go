package integration

type Indexer interface {
	SearchEpisode(seasonNum, episodeNum int, showTitle string, tvdbID *int, imdbID *string) ([]Release, error)
	SearchSeason(seasonNum int, tvdbID *int, imdbID *string) ([]Release, error)
	SearchMovie(movieTitle string, year int, imdbID *string) ([]Release, error)
	SyncRSS(lastRSSID string) ([]Release, error)
	TestConnection() error
}
