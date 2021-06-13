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

func (x *Xnab) Search(media Media) ([]newznab.NZB, error) {
	if x.caps.Searching.Search.Available != "yes" {
		return nil, fmt.Errorf("searching not enabled on indexer")
	}

	if media.ContentType == Movie {
		if x.caps.Searching.MovieSearch.Available == "yes" {
			if strings.Contains(x.caps.Searching.MovieSearch.SupportedParams, "imdbid") {
				return x.client.SearchWithIMDB([]int{newznab.CategoryMovieAll}, media.ImdbID)
			}
			return x.client.SearchWithQuery([]int{newznab.CategoryMovieAll}, media.QueryString(), "movie")
		}
		return x.client.SearchWithQuery([]int{newznab.CategoryMovieAll}, media.QueryString(), "search")
	} else if media.ContentType == TVShow {
		if x.caps.Searching.TvSearch.Available == "yes" {
			if strings.Contains(x.caps.Searching.TvSearch.SupportedParams, "tvdbid") {
				return x.client.SearchWithTVDB([]int{newznab.CategoryTVAll}, media.ParentMedia.TvdbID, media.Season, media.Episode)
			}
			return x.client.SearchWithQuery([]int{newznab.CategoryTVAll}, media.QueryString(), "tvsearch")
		}
		return x.client.SearchWithQuery([]int{newznab.CategoryTVAll}, media.QueryString(), "search")
	}
	return nil, fmt.Errorf("unrecognized media content type: %d", media.ContentType)
}
