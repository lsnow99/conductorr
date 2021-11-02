package integration

import (
	"errors"
	"fmt"
	"strings"

	"github.com/mrobinsn/go-newznab/newznab"
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
		_, err = x.client.SearchWithQuery([]int{newznab.CategoryTVAll}, "The", "movie")
	} else {
		return errors.New("searching is not enabled")
	}
	return err
}

func (x *Xnab) prepareResponse(nzbs []newznab.NZB, media *Media) []Release {
	releases := make([]Release, len(nzbs))
	for i, nzb := range nzbs {
		releases[i] = NewRelease(nzb.ID, nzb.Title, nzb.Description, nzb.DownloadURL, nzb.Category, nzb.Size, int64(nzb.Seeders), nzb.AirDate, nzb.PubDate, media, x, nzb.IMDBID)
	}
	return releases
}

func (x *Xnab) SyncRSS(lastRSSID string) ([]Release, error) {
	results, err := x.client.LoadRSSFeedUntilNZBID([]int{newznab.CategoryMovieAll, newznab.CategoryTVAll}, 50, lastRSSID, 20)
	releases := x.prepareResponse(results, nil)
	return releases, err
}

func (x *Xnab) Search(media *Media) ([]Release, error) {
	if x.caps.Searching.Search.Available != "yes" {
		return nil, fmt.Errorf("searching not enabled on indexer")
	}

	if media.ContentType == Movie {
		if x.caps.Searching.MovieSearch.Available == "yes" {
			if strings.Contains(x.caps.Searching.MovieSearch.SupportedParams, "imdbid") {
				nzbs, err := x.client.SearchWithIMDB([]int{newznab.CategoryMovieAll}, strings.ReplaceAll(media.ImdbID, "t", ""))
				return x.prepareResponse(nzbs, media), err
			}
			nzbs, err := x.client.SearchWithQuery([]int{newznab.CategoryMovieAll}, media.QueryString(), "movie")
			return x.prepareResponse(nzbs, media), err
		}
		nzbs, err := x.client.SearchWithQuery([]int{newznab.CategoryMovieAll}, media.QueryString(), "search")
		return x.prepareResponse(nzbs, media), err
	} else if media.ContentType == TVSeries {
		if x.caps.Searching.TvSearch.Available == "yes" {
			if strings.Contains(x.caps.Searching.TvSearch.SupportedParams, "tvdbid") && media.TvdbID != 0 {
				nzbs, err := x.client.SearchWithTVDB([]int{newznab.CategoryTVAll}, media.TvdbID, media.Season, media.Episode)
				return x.prepareResponse(nzbs, media), err
			}
			q := media.QueryString()
			nzbs, err := x.client.SearchWithQuery([]int{newznab.CategoryTVAll}, q, "tvsearch")
			return x.prepareResponse(nzbs, media), err
		}
		nzbs, err := x.client.SearchWithQuery([]int{newznab.CategoryTVAll}, media.QueryString(), "search")
		return x.prepareResponse(nzbs, media), err
	} else if media.ContentType == TVSeason {
		if x.caps.Searching.TvSearch.Available == "yes" {
			// if strings.Contains
			if strings.Contains(x.caps.Searching.TvSearch.SupportedParams, "tvdbid") && media.ParentMedia.TvdbID != 0 {
				nzbs, err := x.client.SearchWithTVDB()
			}
		}
	} else if media.ContentType == TVEpisode {

	}
	return nil, fmt.Errorf("unrecognized media content type: %d", media.ContentType)
}
