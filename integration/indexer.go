package integration

type Indexer interface {
	Search(media *Media) ([]Release, error)
	SyncRSS(lastRSSID string) ([]Release, error)
	TestConnection() error
}
