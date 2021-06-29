package integration

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/lsnow99/conductorr/constant"
	"github.com/lsnow99/conductorr/csl"
)

type Release struct {
	ID           string    `json:"id,omitempty"`
	Title        string    `json:"title,omitempty"`
	Description  string    `json:"description,omitempty"`
	DownloadURL  string    `json:"download_url,omitempty"`
	Categories   []string  `json:"categories,omitempty"`
	Size         int64     `json:"size,omitempty"`
	Seeders      int64     `json:"seeders,omitempty"`
	AirDate      time.Time `json:"air_date,omitempty"`
	PubDate      time.Time `json:"pub_date,omitempty"`
	Age          int       `json:"age,omitempty"`
	RipType      string    `json:"rip_type,omitempty"`
	Resolution   string    `json:"resolution,omitempty"`
	Encoding     string    `json:"encoding,omitempty"`
	DownloadType string    `json:"download_type,omitempty"`
	Indexer      string    `json:"indexer,omitempty"`
	Warnings     []string  `json:"warnings,omitempty"`
	Media        *Media    `json:"media,omitempty"`

	// Identifier for the media, could be the hash or the nzb name
	Identifier string
	// HighPriority whether to download with high priority
	HighPriority bool
	// FriendlyName friendly name from indexer
	FriendlyName string
}

func NewRelease(id, title, description, downloadURL string, categories []string, size, seeders int64, airDate, pubDate time.Time, media *Media, indexer *Xnab) Release {
	release := Release{
		ID:           id,
		Title:        title,
		Description:  description,
		DownloadURL:  downloadURL,
		Categories:   categories,
		Size:         size,
		Seeders:      seeders,
		AirDate:      airDate,
		PubDate:      pubDate,
		Media:        media,
		DownloadType: indexer.downloadType,
		Indexer:      indexer.name,
	}

	// Calculate age in days
	release.Age = int(time.Now().Sub(pubDate).Hours() / 24)

	// Parse the quality, resolution, encoding, and rip type information
	release.RipType = searchKeywords(title, constant.RipTypes, false)
	release.Resolution = searchKeywords(title, constant.ResolutionTypes, false)
	release.Encoding = searchKeywords(title, constant.EncodingTypes, false)

	if release.Resolution == "" && release.Encoding == "Xvid" {
		release.Resolution = "352p"
	}
	if release.Resolution == "" && release.RipType == "DVDRip" {
		release.Resolution = "480i"
	}

	if release.RipType == "" && release.Encoding == "Xvid" {
		release.RipType = "DVDRip"
	}

	return release
}

func searchKeywords(title string, keywordMap map[string]constant.ParsedOption, secondTry bool) string {
	for k, v := range keywordMap {
		for _, keyword := range v.ParseStrings {
			if secondTry {
				title = strings.ToUpper(title)
				keyword = strings.ToUpper(keyword)
			}
			if strings.Contains(title, keyword) {
				return k
			}
		}
	}
	if !secondTry {
		return searchKeywords(title, keywordMap, true)
	}
	return ""
}

/*
FilterReleases takes a pointer to a slice of Releases, as well as a CSL script
that is expected to always return a boolean. The CSL script will be run once
for each release, and will have the `a` variable set to the CSL list representation
of the given release. The CSL script should always return either true or false.
True means to include the release, false means to filter it out. Returns a slice
of all excluded releases.

If this function ends due to an error, correctness is not guaranteed, and
the releases slice may be corrupt. If data should not be modified in the event
of an error, it is recommended to make a copy of your releases slice before
calling FilterReleases
*/
func FilterReleases(releases []Release, filter string) ([]Release, []Release, error) {
	sexprs, err := csl.Parse(filter)
	if err != nil {
		return nil, nil, err
	}
	included := make([]Release, 0, len(releases))
	excluded := make([]Release, 0, len(releases))
	env := make(map[string]interface{})
	for _, release := range releases {
		env["a"] = makeCSLRelease(release)
		val, trace := csl.Eval(sexprs, env)
		if trace.Err != nil {
			return nil, nil, trace.Err
		}
		include, ok := val.(bool)
		if !ok {
			return nil, nil, fmt.Errorf("csl script returned non boolean value %v", val)
		}
		if include {
			included = append(included, release)
		} else {
			excluded = append(excluded, release)
		}
	}
	return included, excluded, nil
}

/*
SortReleases takes a pointer to a slice of Releases, as well as a CSL script
whose job is to return true if release A should be preferred over release B,
and false otherwise. The CSL script will be run with the `a` and `b` variables
set to the CSL list representation of the pair of releases. The CSL script
should always return either true or false. True means release A > release B.

If this function ends due to an error, correctness is not guaranteed, and
the releases slice may be corrupt. If data should not be modified in the event
of an error, it is recommended to make a copy of your releases slice before
calling SortReleases
*/
func SortReleases(releases *[]Release, sorter string) error {
	if releases == nil {
		return nil
	}
	sexprs, err := csl.Parse(sorter)
	if err != nil {
		return err
	}
	env := make(map[string]interface{})
	sort.Slice(*releases, func(i, j int) bool {
		env["a"] = makeCSLRelease((*releases)[i])
		env["b"] = makeCSLRelease((*releases)[j])
		val, trace := csl.Eval(sexprs, env)
		if trace.Err != nil {
			err = trace.Err
		}
		lessThan, ok := val.(bool)
		if !ok {
			err = fmt.Errorf("csl script returned non boolean value %v", val)
			return false
		}
		return lessThan
	})
	if err != nil {
		return err
	}
	return nil
}

func makeCSLRelease(release Release) csl.List {
	return csl.List{
		Elems: []interface{}{
			release.Title,
			release.Indexer,
			release.DownloadType,
			release.Media.ContentType,
			release.RipType,
			release.Resolution,
			release.Encoding,
			release.Seeders,
			release.Age,
			release.Size,
			release.Media.Runtime,
		},
	}
}
