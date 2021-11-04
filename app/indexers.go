package app

import (
	"errors"
	"sync"
	"time"

	"github.com/lsnow99/conductorr/dbstore"
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

// sync rss feeds
func (im *IndexerManager) DoTask() {
	indexers := im.getIndexers()
	for _, indexer := range indexers {
		results, err := indexer.SyncRSS(indexer.LastRSSID)
		if err != nil {
			logger.LogWarn(err)
			continue
		}
		if len(results) > 0 {
			im.updateIndexerLastRSSID(indexer.ID, results[len(results)-1].ID)
			monitoringMedia, err := dbstore.GetMonitoringMedia()
			if err != nil || monitoringMedia == nil {
				return
			}
			for _, result := range results {
				for _, media := range *monitoringMedia {
					if result.ImdbID == media.ImdbID.String && result.ImdbID != "" {
						// Found matching release
					}
				}
			}
		}
	}
}

func (im *IndexerManager) updateIndexerLastRSSID(indexerID int, rssID string) {
	im.Lock()
	defer im.Unlock()
	for i, indexer := range im.indexers {
		if indexer.ID == indexerID {
			indexer.LastRSSID = rssID
			im.indexers[i] = indexer
		}
	}
}

func (*IndexerManager) GetTaskName() string {
	return "Indexer Manager"
}

func (im *IndexerManager) GetFrequency() time.Duration {
	return time.Minute * 60
}

func (im *IndexerManager) getIndexers() []ManagedIndexer {
	im.RLock()
	defer im.RUnlock()
	return im.indexers
}

func (im *IndexerManager) RegisterIndexer(id int, downloadType string, userID int, name, apiKey, baseUrl string, forMovies, forSeries bool, lastRSSID string) {
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

	var updated bool
	for i, indexer := range im.indexers {
		if indexer.ID == id {
			im.indexers[i] = mi
			updated = true
		}
	}
	if !updated {
		// Only use the provided last rss id if this is the first time we are registering the indexer
		mi.LastRSSID = lastRSSID
		im.indexers = append(im.indexers, mi)
	}
}

func (im *IndexerManager) doSearch(searchFunc func(indexer integration.Indexer)([]integration.Release, error)) ([]integration.Release, error) {
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
		go func(wg *sync.WaitGroup, indexer integration.Indexer, results *[]integration.Release) {
			indexerResults, err := searchFunc(indexer)
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

func (im *IndexerManager) SearchEpisode(seasonNum, episodeNum int, showTitle string, tvdbID *int, imdbID *string) ([]integration.Release, error) {
	return im.doSearch(func(indexer integration.Indexer) ([]integration.Release, error) {
		return indexer.SearchEpisode(seasonNum, episodeNum, showTitle, tvdbID, imdbID)
	})
}

func (im *IndexerManager) SearchSeason(seasonNum int, showTitle string, tvdbID *int, imdbID *string) ([]integration.Release, error) {
	return im.doSearch(func(indexer integration.Indexer) ([]integration.Release, error) {
		return indexer.SearchSeason(seasonNum, tvdbID, imdbID)
	})
}

func (im *IndexerManager) SearchMovie(movieTitle string, year int, imdbID *string) ([]integration.Release, error) {
	return im.doSearch(func(indexer integration.Indexer) ([]integration.Release, error) {
		return indexer.SearchMovie(movieTitle, year, imdbID)
	})
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
