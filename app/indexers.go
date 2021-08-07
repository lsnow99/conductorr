package app

import (
	"errors"
	"sync"
	"time"

	"github.com/lsnow99/conductorr/integration"
	"github.com/lsnow99/conductorr/logger"
)

type ManagedIndexer struct {
	ID        int
	Name      string
	ForMovies bool
	ForSeries bool
	LastRSSID string
	integration.Indexer
}

type IndexerManager struct {
	indexers []ManagedIndexer
	sync.RWMutex
}

func (mi ManagedIndexer) GetName() string {
	return mi.Name
}

// TODO get rss feeds
func (im *IndexerManager) DoTask() {
	indexers := im.getIndexers()
	for _, indexer := range indexers {
		indexer.SyncRSS()
	}
}

func (im *IndexerManager) GetFrequency() time.Duration {
	return time.Minute * 60
}

func (im *IndexerManager) getIndexers() []ManagedIndexer {
	im.RLock()
	defer im.RUnlock()
	return im.indexers
}

func (im *IndexerManager) RegisterIndexer(id int, downloadType string, userID int, name, apiKey, baseUrl string, forMovies, forSeries bool) {
	indexer := integration.NewXnab(userID, apiKey, baseUrl, name, downloadType)
	mi := ManagedIndexer{
		Name:      name,
		ForMovies: forMovies,
		ForSeries: forSeries,
		Indexer:   indexer,
	}
	err := mi.TestConnection()
	if err != nil {
		logger.LogWarn(err)
	}

	im.Lock()
	defer im.Unlock()
	
	var added bool
	for i, indexer := range im.indexers {
		if indexer.ID == id {
			im.indexers[i] = mi
			added = true
		}
	}
	if !added {
		im.indexers = append(im.indexers, mi)
	}
}

func (im *IndexerManager) Search(media *integration.Media) ([]integration.Release, error) {
	results := make([]integration.Release, 0)
	indexers := im.getIndexers()

	if len(indexers) == 0 {
		return results, errors.New("no indexers registered")
	}
	var wg sync.WaitGroup
	resultsArrs := make([]*[]integration.Release, len(indexers))
	for i, indexer := range indexers {
		resultsArrs[i] = &[]integration.Release{}
		wg.Add(1)
		go func (wg *sync.WaitGroup, indexer integration.Indexer, results *[]integration.Release)  {
			indexerResults, err := indexer.Search(media)
			if err != nil {
				logger.LogWarn(err)
			}
			*results = indexerResults
			defer wg.Done()
		}(&wg, indexer, resultsArrs[i])
	}
	wg.Wait()
	for _, resultsArr := range resultsArrs {
		if resultsArr != nil {
			results = append(results, *resultsArr...)
		}
	}
	return results, nil
}

func (im *IndexerManager) DeleteIndexer(id int) {
	im.Lock()
	defer im.Unlock()
	for i, indexer := range im.indexers {
		if indexer.ID == id {
			im.indexers = append(im.indexers[:i], im.indexers[i+1:]...)
			return
		}
	}
}
