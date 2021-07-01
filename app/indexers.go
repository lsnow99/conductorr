package app

import (
	"time"

	"github.com/lsnow99/conductorr/integration"
)

type ManagedIndexer struct {
	Name      string
	ForMovies bool
	ForSeries bool
	integration.Indexer
}

type IndexerManager struct {
	indexers []ManagedIndexer
}

func (mi ManagedIndexer) GetName() string {
	return mi.Name
}

// TODO get rss feeds
func (im *IndexerManager) DoTask() {
}

func (im *IndexerManager) GetFrequency() time.Duration {
	return time.Minute * 60
}

func (im *IndexerManager) getIndexers() []ManagedIndexer {
	return im.indexers
}

func (im *IndexerManager) RegisterIndexer(downloadType string, userID int, name, apiKey, baseUrl string, forMovies, forSeries bool) {
	indexer := integration.NewXnab(userID, apiKey, baseUrl, name, downloadType)
	im.indexers = append(im.indexers, ManagedIndexer{
		Name:      name,
		ForMovies: forMovies,
		ForSeries: forSeries,
		Indexer:   indexer,
	})
}

func (im *IndexerManager) Search(media *integration.Media) ([]integration.Release, error) {
	results := make([]integration.Release, 0)
	for _, indexer := range im.indexers {
		indexer.TestConnection()
		indexerResults, err := indexer.Search(media)
		if err != nil {
			return nil, err
		}
		results = append(results, indexerResults...)
	}
	return results, nil
}
