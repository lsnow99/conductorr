package integration

import (
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"sync"
)

type Plex struct {
	token   string
	baseUrl string
}

type jobResultPtr *jobResult

type jobResult struct {
	mediaPaths []MediaPath
	err        error
}

func (p *Plex) TestConnection() error {
	u, err := url.Parse(p.baseUrl)
	if err != nil {
		return err
	}
	q := url.Values{}
	q.Add("X-Plex-Token", p.token)
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("received non-200 level status code %d", resp.StatusCode)
	}

	return nil
}

func (p *Plex) ImportMedia(path string) error {
	libraries, err := p.getLibraries()
	if err != nil {
		return err
	}
	for _, library := range libraries {
		u, err := url.Parse(p.baseUrl)
		if err != nil {
			return err
		}
		q := url.Values{}
		q.Add("X-Plex-Token", p.token)
		q.Add("path", path)
		u.Path = "/library/sections/" + library.Key + "/refresh"
		u.RawQuery = q.Encode()

		req, err := http.NewRequest("GET", u.String(), nil)
		if err != nil {
			return err
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}

		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			return fmt.Errorf("received non-200 level status code %d", resp.StatusCode)
		}

		resp.Body.Close()
	}
	return nil
}

func NewPlex(token, baseUrl string) *Plex {
	p := &Plex{
		token:   token,
		baseUrl: baseUrl,
	}

	return p
}

func NewPlexFromConfig(configuration map[string]interface{}) (*Plex, error) {
	token, tOK := configuration["token"].(string)
	baseUrl, bOK := configuration["base_url"].(string)
	if !tOK || !bOK {
		return nil, errors.New("failed to parse plex configuration")
	}

	return NewPlex(token, baseUrl), nil
}

type Directory struct {
	Text       string `xml:",chardata"`
	AllowSync  string `xml:"allowSync,attr"`
	Art        string `xml:"art,attr"`
	Filters    string `xml:"filters,attr"`
	Refreshing string `xml:"refreshing,attr"`
	Thumb      string `xml:"thumb,attr"`
	RatingKey  string `xml:"ratingKey,attr"`
	Key        string `xml:"key,attr"`
	Type       string `xml:"type,attr"`
	Title      string `xml:"title,attr"`
	Agent      string `xml:"agent,attr"`
	Scanner    string `xml:"scanner,attr"`
	Language   string `xml:"language,attr"`
	Uuid       string `xml:"uuid,attr"`
	UpdatedAt  string `xml:"updatedAt,attr"`
	CreatedAt  string `xml:"createdAt,attr"`
	Locations  []struct {
		Text string `xml:",chardata"`
		ID   string `xml:"id,attr"`
		Path string `xml:"path,attr"`
	} `xml:"Location"`
}

type MediaContainer struct {
	XMLName         xml.Name    `xml:"MediaContainer"`
	Text            string      `xml:",chardata"`
	Size            string      `xml:"size,attr"`
	AllowSync       string      `xml:"allowSync,attr"`
	Identifier      string      `xml:"identifier,attr"`
	MediaTagPrefix  string      `xml:"mediaTagPrefix,attr"`
	MediaTagVersion string      `xml:"mediaTagVersion,attr"`
	Title1          string      `xml:"title1,attr"`
	Directories     []Directory `xml:"Directory"`
	Videos          []Video     `xml:"Video"`
}

type Video struct {
	Text                  string `xml:",chardata"`
	RatingKey             string `xml:"ratingKey,attr"`
	Key                   string `xml:"key,attr"`
	Guid                  string `xml:"guid,attr"`
	Studio                string `xml:"studio,attr"`
	Type                  string `xml:"type,attr"`
	Title                 string `xml:"title,attr"`
	OriginalTitle         string `xml:"originalTitle,attr"`
	GrandparentTitle      string `xml:"grandparentTitle,attr"`
	ParentTitle           string `xml:"parentTitle,attr"`
	ContentRating         string `xml:"contentRating,attr"`
	Summary               string `xml:"summary,attr"`
	Rating                string `xml:"rating,attr"`
	AudienceRating        string `xml:"audienceRating,attr"`
	Year                  string `xml:"year,attr"`
	Tagline               string `xml:"tagline,attr"`
	Thumb                 string `xml:"thumb,attr"`
	Art                   string `xml:"art,attr"`
	Duration              string `xml:"duration,attr"`
	OriginallyAvailableAt string `xml:"originallyAvailableAt,attr"`
	AddedAt               string `xml:"addedAt,attr"`
	UpdatedAt             string `xml:"updatedAt,attr"`
	AudienceRatingImage   string `xml:"audienceRatingImage,attr"`
	ChapterSource         string `xml:"chapterSource,attr"`
	PrimaryExtraKey       string `xml:"primaryExtraKey,attr"`
	RatingImage           string `xml:"ratingImage,attr"`
	ViewCount             string `xml:"viewCount,attr"`
	LastViewedAt          string `xml:"lastViewedAt,attr"`
	Media                 struct {
		Text                  string `xml:",chardata"`
		ID                    string `xml:"id,attr"`
		Duration              string `xml:"duration,attr"`
		Bitrate               string `xml:"bitrate,attr"`
		Width                 string `xml:"width,attr"`
		Height                string `xml:"height,attr"`
		AspectRatio           string `xml:"aspectRatio,attr"`
		AudioChannels         string `xml:"audioChannels,attr"`
		AudioCodec            string `xml:"audioCodec,attr"`
		VideoCodec            string `xml:"videoCodec,attr"`
		VideoResolution       string `xml:"videoResolution,attr"`
		Container             string `xml:"container,attr"`
		VideoFrameRate        string `xml:"videoFrameRate,attr"`
		AudioProfile          string `xml:"audioProfile,attr"`
		VideoProfile          string `xml:"videoProfile,attr"`
		OptimizedForStreaming string `xml:"optimizedForStreaming,attr"`
		Has64bitOffsets       string `xml:"has64bitOffsets,attr"`
		Part                  struct {
			Text                  string `xml:",chardata"`
			ID                    string `xml:"id,attr"`
			Key                   string `xml:"key,attr"`
			Duration              string `xml:"duration,attr"`
			File                  string `xml:"file,attr"`
			Size                  string `xml:"size,attr"`
			AudioProfile          string `xml:"audioProfile,attr"`
			Container             string `xml:"container,attr"`
			VideoProfile          string `xml:"videoProfile,attr"`
			Has64bitOffsets       string `xml:"has64bitOffsets,attr"`
			HasThumbnail          string `xml:"hasThumbnail,attr"`
			OptimizedForStreaming string `xml:"optimizedForStreaming,attr"`
		} `xml:"Part"`
	} `xml:"Media"`
	Genre []struct {
		Text string `xml:",chardata"`
		Tag  string `xml:"tag,attr"`
	} `xml:"Genre"`
	Director struct {
		Text string `xml:",chardata"`
		Tag  string `xml:"tag,attr"`
	} `xml:"Director"`
	Writer []struct {
		Text string `xml:",chardata"`
		Tag  string `xml:"tag,attr"`
	} `xml:"Writer"`
	Country struct {
		Text string `xml:",chardata"`
		Tag  string `xml:"tag,attr"`
	} `xml:"Country"`
	Role []struct {
		Text string `xml:",chardata"`
		Tag  string `xml:"tag,attr"`
	} `xml:"Role"`
}

func (p *Plex) getLibraries() ([]Directory, error) {
	u, err := url.Parse(p.baseUrl)
	if err != nil {
		return nil, err
	}
	q := url.Values{}
	q.Add("X-Plex-Token", p.token)
	u.Path = "/library/sections"
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("received non-200 level status code %d", resp.StatusCode)
	}

	librariesResponse := MediaContainer{}
	if err := xml.NewDecoder(resp.Body).Decode(&librariesResponse); err != nil {
		return nil, err
	}
	return librariesResponse.Directories, nil
}

func (p *Plex) getDirectoryListing(key string) (mediaContainer MediaContainer, err error) {
	var u *url.URL
	var req *http.Request
	var resp *http.Response

	u, err = url.Parse(p.baseUrl)
	if err != nil {
		return
	}
	q := url.Values{}
	q.Add("X-Plex-Token", p.token)
	u.Path = key
	u.RawQuery = q.Encode()

	req, err = http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return
	}

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		err = fmt.Errorf("received non-200 level status code %d", resp.StatusCode)
		return
	}

	err = xml.NewDecoder(resp.Body).Decode(&mediaContainer)
	return
}

func (p *Plex) getMediaInLibrary(wg *sync.WaitGroup, library Directory, result jobResultPtr) {
	defer wg.Done()

	listing, err := p.getDirectoryListing("/library/sections/" + library.Key + "/allLeaves")
	if err != nil {
		result.err = err
		return
	}
	result.mediaPaths = make([]MediaPath, len(listing.Videos))
	for i, video := range listing.Videos {
		result.mediaPaths[i] = MediaPath{
			Title: video.Title,
			Path:  video.Media.Part.File,
		}
	}
}

func (p *Plex) GetMediaPaths() ([]MediaPath, error) {
	libraries, err := p.getLibraries()
	if err != nil {
		return nil, err
	}

	results := make([]jobResultPtr, len(libraries))
	wg := sync.WaitGroup{}
	for i, library := range libraries {
		var result jobResultPtr = &jobResult{}
		results[i] = result
		wg.Add(1)
		go p.getMediaInLibrary(&wg, library, result)
	}
	wg.Wait()

	mediaPaths := make([]MediaPath, 0)
	for _, result := range results {
		if result == nil || result.mediaPaths == nil {
			continue
		}
		if result.err != nil {
			return nil, err
		}
		mediaPaths = append(mediaPaths, result.mediaPaths...)
	}
	return mediaPaths, nil
}
