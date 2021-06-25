package integration

import (
	"strings"
	"time"

	"github.com/lsnow99/conductorr/constant"
)

type Release struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DownloadURL string    `json:"download_url"`
	Categories  []string  `json:"categories"`
	Size        int64     `json:"size"`
	AirDate     time.Time `json:"air_date"`
	PubDate     time.Time `json:"pub_date"`
	Media       *Media    `json:"-"`
	Age         int       `json:"age"`
	RipType     string    `json:"rip_type"`
	Resolution  string    `json:"resolution"`
	Encoding    string    `json:"encoding"`
	DownloadType string `json:"download_type"`
	Indexer string `json:"indexer"`
}

func NewRelease(id, title, description, downloadURL string, categories []string, size int64, airDate, pubDate time.Time, media *Media, indexer *Xnab) Release {
	release := Release{
		ID: id,
		Title: title,
		Description: description,
		DownloadURL: downloadURL,
		Categories: categories,
		Size: size,
		AirDate: airDate,
		PubDate: pubDate,
		Media: media,
		DownloadType: indexer.downloadType,
		Indexer: indexer.name,
	}

	// Calculate age in days
	release.Age = int(time.Now().Sub(pubDate).Hours() / 24)

	// Parse the quality, resolution, encoding, and rip type information
	release.RipType = searchKeywords(title, constant.RipTypes, false)
	release.Resolution = searchKeywords(title, constant.ResolutionTypes, false)
	release.Encoding = searchKeywords(title, constant.EncodingTypes, false)

	return release
}

func searchKeywords(title string, keywordMap map[string][]string, secondTry bool) string {
	for k, v := range keywordMap {
		for _, keyword := range v {
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