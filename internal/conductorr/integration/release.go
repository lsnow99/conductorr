package integration

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/lsnow99/conductorr/pkg/constant"
	"github.com/lsnow99/conductorr/pkg/csllib"
	"github.com/lsnow99/conductorr/internal/csl"
	"github.com/lsnow99/conductorr/internal/conductorr/settings"
)

type Release struct {
	ID           string    `json:"id,omitempty"`
	Title        string    `json:"title,omitempty"`
	Description  string    `json:"description,omitempty"`
	DownloadURL  string    `json:"downloadUrl,omitempty"`
	Categories   []string  `json:"categories,omitempty"`
	Size         int64     `json:"size,omitempty"`
	Seeders      int64     `json:"seeders,omitempty"`
	AirDate      time.Time `json:"airDate,omitempty"`
	PubDate      time.Time `json:"pubDate,omitempty"`
	Age          int       `json:"age,omitempty"`
	RipType      string    `json:"ripType,omitempty"`
	Resolution   string    `json:"resolution,omitempty"`
	Encoding     string    `json:"encoding,omitempty"`
	DownloadType string    `json:"downloadType,omitempty"`
	Indexer      string    `json:"indexer,omitempty"`
	Warnings     []string  `json:"warnings,omitempty"`
	IndexerID    int       `json:"indexerId,omitempty"`
	ImdbID       string    `json:"imdbId,omitempty"`
	ContentType  string    `json:"contentType,omitempty"`

	// HighPriority whether to download with high priority
	HighPriority bool `json:"highPriority,omitempty"`
}

func NewRelease(id, title, description, downloadURL string, categories []string, size, seeders int64, airDate, pubDate time.Time, indexer *Xnab, imdbID, contentType string) Release {
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
		DownloadType: indexer.downloadType,
		Indexer:      indexer.name,
		ImdbID:       imdbID,
		ContentType:  contentType,
	}

	// Calculate age in days
	release.Age = int(time.Since(pubDate).Hours() / 24)

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
of the given release. The CSL script should either true or false, OR it can return
a string/list of strings that will be treated the same as a return false. The string(s)
will be displayed as errors to the user for why the release failed the filter.
True means to include the release, false means to filter it out. Returns a slice
of all included and excluded releases, as well as a non-nil error if the CSL script
fails to execute properly.

If this function ends due to an error, no guarantees are made for the returned release
slices.
*/
func FilterReleases(releases []Release, filter string, runtime *int64) ([]Release, []Release, error) {
	csl := csl.NewCSL()
	cslpm := csllib.NewCSLPackageManager(csllib.DefaultFetcher, settings.DebugMode)
	if err := csl.PreprocessScript(filter, "", cslpm); err != nil {
		return nil, nil, err
	}
	sexprs, err := csl.Parse(filter)
	if err != nil {
		return nil, nil, err
	}

	included := make([]Release, 0, len(releases))
	excluded := make([]Release, 0, len(releases))
	for _, release := range releases {
		val, trace := csl.Invoke(sexprs, makeCSLRelease(release, runtime))
		if trace.Err != nil {
			return nil, nil, trace.Err
		}
		include, ok := val.(bool)
		if !ok {
			str, ok := val.(string)
			if ok {
				release.Warnings = append(release.Warnings, str)
			} else if l, ok := val.(csllib.List); ok {
				warnings := make([]string, 0, len(l))
				for _, elem := range l {
					if s, ok := elem.(string); ok {
						warnings = append(warnings, s)
					} else {
						return nil, nil, fmt.Errorf("csl script returned non boolean value %v", val)
					}
				}
				release.Warnings = append(release.Warnings, warnings...)
			} else {
				return nil, nil, fmt.Errorf("csl script returned non boolean value %v", val)
			}
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
func SortReleases(releases *[]Release, sorter string, runtime *int64) error {
	if releases == nil {
		return nil
	}

	csl := csl.NewCSL()
	cslpm := csllib.NewCSLPackageManager(csllib.DefaultFetcher, settings.DebugMode)
	if err := csl.PreprocessScript(sorter, "", cslpm); err != nil {
		return err
	}
	sexprs, err := csl.Parse(sorter)
	if err != nil {
		return err
	}

	sort.SliceStable(*releases, func(i, j int) bool {
		val, trace := csl.Invoke(sexprs, makeCSLRelease((*releases)[i], runtime), makeCSLRelease((*releases)[j], runtime))
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

func makeCSLRelease(release Release, runtime *int64) csllib.List {
	var rt interface{}
	if runtime == nil {
		rt = nil
	} else {
		rt = *runtime
	}
	return csllib.List{
		release.Title,
		release.Indexer,
		release.DownloadType,
		release.ContentType,
		release.RipType,
		release.Resolution,
		release.Encoding,
		release.Seeders,
		release.Age,
		release.Size,
		rt,
	}
}
