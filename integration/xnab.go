package integration

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/lsnow99/go-newznab/newznab"
)

type Xnab struct {
	userID       int
	apiKey       string
	baseUrl      string
	name         string
	caps         newznab.Capabilities
	client       newznab.Client
	downloadType string
}

func NewXnab(userID int, apiKey, baseUrl, name, downloadType string) *Xnab {
	x := &Xnab{
		userID:       userID,
		apiKey:       apiKey,
		baseUrl:      baseUrl,
		name:         name,
		client:       newznab.New(baseUrl, apiKey, userID, true),
		downloadType: downloadType,
	}
	return x
}

func (x *Xnab) TestConnection() error {
	caps, err := x.client.Capabilities()
	if err != nil {
		return err
	}
	x.caps = caps
	if x.caps.Searching.MovieSearch.Available == "yes" {
		// Check auth
		_, err = x.client.SearchWithQuery([]int{newznab.CategoryMovieAll}, "The", "movie")
	} else if x.caps.Searching.TvSearch.Available == "yes" {
		// Check auth
		_, err = x.client.SearchWithQuery([]int{newznab.CategoryTVAll}, "The", "series")
	} else {
		return errors.New("searching is not enabled")
	}
	return err
}

func (x *Xnab) prepareResponse(nzbs []newznab.NZB, contentTypeExtractor func(nzb newznab.NZB) string) []Release {
	releases := make([]Release, len(nzbs))
	for i, nzb := range nzbs {
		contentType := contentTypeExtractor(nzb)
		releases[i] = NewRelease(nzb.ID, nzb.Title, nzb.Description, nzb.DownloadURL, nzb.Category, nzb.Size, int64(nzb.Seeders), nzb.AirDate, nzb.PubDate, x, nzb.IMDBID, contentType)
	}
	return releases
}

func rssContentTypeExtractor(nzb newznab.NZB) string {
	for _, cat := range nzb.Category {
		if cat == strconv.Itoa(newznab.CategoryTVAll) {
			// TODO: figure out how to differentiate seasons
			return "episode"
		} else if cat == strconv.Itoa(newznab.CategoryMovieAll) {
			return "movie"
		}
	}
	return ""
}

func movieContentTypeExtractor(newznab.NZB) string   { return "movie" }
func seasonContentTypeExtractor(newznab.NZB) string  { return "season" }
func episodeContentTypeExtractor(newznab.NZB) string { return "episode" }

func (x *Xnab) SyncRSS(lastRSSID string) ([]Release, error) {
	results, err := x.client.LoadRSSFeedUntilNZBID([]int{newznab.CategoryMovieAll, newznab.CategoryTVAll}, 50, lastRSSID, 20)
	releases := x.prepareResponse(results, rssContentTypeExtractor)
	return releases, err
}

func (x *Xnab) SearchEpisode(seasonNum, episodeNum int, showTitle string, tvdbID *int, imdbID *string) ([]Release, error) {
	if x.caps.Searching.Search.Available != "yes" {
		return nil, fmt.Errorf("searching not enabled on indexer")
	}

	var allResults []newznab.NZB
	if tvdbID != nil && strings.Contains(x.caps.Searching.TvSearch.SupportedParams, "tvdbid") {
		nzbs, err := x.client.SearchWithParams([]int{newznab.CategoryTVAll}, "tvsearch", map[string][]string{
			"tvdbid":  {strconv.Itoa(*tvdbID)},
			"season":  {strconv.Itoa(seasonNum)},
			"episode": {strconv.Itoa(episodeNum)},
		})
		if err != nil {
			return nil, err
		}
		allResults = append(allResults, nzbs...)
	}
	if imdbID != nil && strings.Contains(x.caps.Searching.TvSearch.SupportedParams, "imdbid") {
		nzbs, err := x.client.SearchWithParams([]int{newznab.CategoryTVAll}, "tvsearch", map[string][]string{
			"imdbid":  {*imdbID},
			"season":  {strconv.Itoa(seasonNum)},
			"episode": {strconv.Itoa(episodeNum)},
		})
		if err != nil {
			return nil, err
		}
		allResults = append(allResults, nzbs...)
	}

	if len(allResults) == 0 {
		nzbs, err := x.client.SearchWithQuery([]int{newznab.CategoryTVAll}, fmt.Sprintf("%s S%02dE%02d", showTitle, seasonNum, episodeNum), "tvsearch")
		if err != nil {
			return nil, err
		}
		allResults = append(allResults, nzbs...)
	}

	return x.prepareResponse(allResults, episodeContentTypeExtractor), nil
}

func (x *Xnab) SearchSeason(seasonNum int, tvdbID *int, imdbID *string) ([]Release, error) {
	if x.caps.Searching.Search.Available != "yes" {
		return nil, fmt.Errorf("searching not enabled on indexer")
	}

	var allResults []newznab.NZB
	if tvdbID != nil && strings.Contains(x.caps.Searching.TvSearch.SupportedParams, "tvdbid") {
		nzbs, err := x.client.SearchWithParams([]int{newznab.CategoryTVAll}, "tvsearch", map[string][]string{
			"tvdbid": {strconv.Itoa(*tvdbID)},
			"season": {strconv.Itoa(seasonNum)},
		})
		if err != nil {
			return nil, err
		}
		allResults = append(allResults, nzbs...)
	}
	if imdbID != nil && strings.Contains(x.caps.Searching.TvSearch.SupportedParams, "imdbid") {
		nzbs, err := x.client.SearchWithParams([]int{newznab.CategoryTVAll}, "tvsearch", map[string][]string{
			"imdbid": {*imdbID},
			"season": {strconv.Itoa(seasonNum)},
		})
		if err != nil {
			return nil, err
		}
		allResults = append(allResults, nzbs...)
	}

	return x.prepareResponse(allResults, seasonContentTypeExtractor), nil
}

func (x *Xnab) SearchMovie(movieTitle string, year int, imdbID *string) ([]Release, error) {
	if x.caps.Searching.Search.Available != "yes" {
		return nil, fmt.Errorf("searching not enabled on indexer")
	}

	var allResults []newznab.NZB
	if imdbID != nil && strings.Contains(x.caps.Searching.TvSearch.SupportedParams, "imdbid") {
		nzbs, err := x.client.SearchWithIMDB([]int{newznab.CategoryMovieAll}, *imdbID)
		if err != nil {
			return nil, err
		}
		allResults = append(allResults, nzbs...)
	}

	if len(allResults) == 0 {
		nzbs, err := x.client.SearchWithQuery([]int{newznab.CategoryMovieAll}, fmt.Sprintf("%s %d", movieTitle, year), "moviesearch")
		if err != nil {
			return nil, err
		}
		allResults = append(allResults, nzbs...)
	}

	return x.prepareResponse(allResults, seasonContentTypeExtractor), nil
}
