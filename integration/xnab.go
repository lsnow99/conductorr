package integration

import (
	"fmt"
	"strings"

	"github.com/mrobinsn/go-newznab/newznab"
)

type Xnab struct {
	userID  int
	apiKey  string
	baseUrl string
	name    string
	caps    newznab.Capabilities
	client  newznab.Client
}

func NewXnab(userID int, apiKey, baseUrl, name string) *Xnab {
	x := &Xnab{
		userID:  userID,
		apiKey:  apiKey,
		baseUrl: baseUrl,
		name:    name,
		client:  newznab.New(baseUrl, apiKey, userID, true),
	}
	return x
}

func (x *Xnab) TestConnection() error {
	caps, err := x.client.Capabilities()
	x.caps = caps
	return err
}

func prepareResponse(nzbs []newznab.NZB, media *Media) ([]Release) {
	releases := make([]Release, len(nzbs))
	for i, nzb := range nzbs {
		releases[i] = NewRelease(nzb.ID, nzb.Title, nzb.Description, nzb.DownloadURL, nzb.Category, nzb.Size, nzb.AirDate, nzb.PubDate, media)
	}
}

func (x *Xnab) Search(media *Media) ([]Release, error) {
	if x.caps.Searching.Search.Available != "yes" {
		return nil, fmt.Errorf("searching not enabled on indexer")
	}

	if media.ContentType == Movie {
		if x.caps.Searching.MovieSearch.Available == "yes" {
			if strings.Contains(x.caps.Searching.MovieSearch.SupportedParams, "imdbid") {
				nzbs, err := x.client.SearchWithIMDB([]int{newznab.CategoryMovieAll}, strings.ReplaceAll(media.ImdbID, "t", ""))
				return prepareResponse(nzbs, media), err
			}
			nzbs, err := x.client.SearchWithQuery([]int{newznab.CategoryMovieAll}, media.QueryString(), "movie")
			return prepareResponse(nzbs, media), err
		}
		nzbs, err := x.client.SearchWithQuery([]int{newznab.CategoryMovieAll}, media.QueryString(), "search")
		return prepareResponse(nzbs, media), err
	} else if media.ContentType == TVShow {
		if x.caps.Searching.TvSearch.Available == "yes" {
			if strings.Contains(x.caps.Searching.TvSearch.SupportedParams, "tvdbid") {
				nzbs, err := x.client.SearchWithTVDB([]int{newznab.CategoryTVAll}, media.ParentMedia.TvdbID, media.Season, media.Episode)
				return prepareResponse(nzbs, media), err
			}
			nzbs, err := x.client.SearchWithQuery([]int{newznab.CategoryTVAll}, media.QueryString(), "tvsearch")
			return prepareResponse(nzbs, media), err
		}
		nzbs, err := x.client.SearchWithQuery([]int{newznab.CategoryTVAll}, media.QueryString(), "search")
		return prepareResponse(nzbs, media), err
	}
	return nil, fmt.Errorf("unrecognized media content type: %d", media.ContentType)
}
