package integration

type Indexer interface {
	Search(media *Media) ([]Release, error)
	TestConnection() error
}
