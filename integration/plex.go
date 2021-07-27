package integration

import (
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"path/filepath"

	"github.com/lsnow99/conductorr/logger"
)

type Plex struct {
	token   string
	baseUrl string
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
	fmt.Println("importing on plex")
	libraries, err := p.getLibraries()
	if err != nil {
		return err
	}
	fmt.Printf("got libraries %v\n", libraries)
	for _, library := range libraries {
		fmt.Printf("checking library %v\n", library)
		for _, location := range library.Locations {
			fmt.Printf("checking location %v with path %s\n", location, location.Path)
			pattern := location.Path + string(filepath.Separator) + "*"
			matched, err := filepath.Match(pattern, path)
			if err != nil {
				fmt.Println("got error 1 ", err)
				logger.LogDanger(err)
				continue
			}
			if matched {
				fmt.Printf("path %s is a match for %v", path, location.Path)
				u, err := url.Parse(p.baseUrl)
				if err != nil {
					fmt.Println("got error 2 ", err)
					return err
				}
				q := url.Values{}
				q.Add("X-Plex-Token", p.token)
				q.Add("path", path)
				u.Path = "/library/sections/" + library.Key + "/refresh"
				u.RawQuery = q.Encode()

				req, err := http.NewRequest("GET", u.String(), nil)
				if err != nil {
					fmt.Println("got error 3 ", err)
					return err
				}

				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					fmt.Println("got error 4 ", err)
					return err
				}

				if resp.StatusCode < 200 || resp.StatusCode >= 300 {
					fmt.Println("got error 5 ", resp.StatusCode)
					return fmt.Errorf("received non-200 level status code %d", resp.StatusCode)
				} else {
					fmt.Println("refresh success")
				}

				resp.Body.Close()
			} else {
				fmt.Printf("path %s is not a match for %v", path, location.Path)
			}
		}
	}
	return nil
}

func NewPlex(token, baseUrl string) (*Plex, error) {
	p := &Plex{
		token:   token,
		baseUrl: baseUrl,
	}

	return p, nil
}

func NewPlexFromConfig(configuration map[string]interface{}) (*Plex, error) {
	token, tOK := configuration["token"].(string)
	baseUrl, bOK := configuration["base_url"].(string)
	if !tOK || !bOK {
		return nil, errors.New("failed to parse plex configuration")
	}

	return NewPlex(token, baseUrl)
}

type Library struct {
	Text       string `xml:",chardata"`
	AllowSync  string `xml:"allowSync,attr"`
	Art        string `xml:"art,attr"`
	Filters    string `xml:"filters,attr"`
	Refreshing string `xml:"refreshing,attr"`
	Thumb      string `xml:"thumb,attr"`
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
	XMLName         xml.Name  `xml:"MediaContainer"`
	Text            string    `xml:",chardata"`
	Size            string    `xml:"size,attr"`
	AllowSync       string    `xml:"allowSync,attr"`
	Identifier      string    `xml:"identifier,attr"`
	MediaTagPrefix  string    `xml:"mediaTagPrefix,attr"`
	MediaTagVersion string    `xml:"mediaTagVersion,attr"`
	Title1          string    `xml:"title1,attr"`
	Libraries       []Library `xml:"Directory"`
}

func (p *Plex) getLibraries() ([]Library, error) {
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
	return librariesResponse.Libraries, nil
}
