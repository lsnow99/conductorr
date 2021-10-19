package csl

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"regexp"
	"strings"
)

const ExampleGit = "github.com/lsnow99/mycslscripts:main.csl"

var gitPattern = regexp.MustCompile(`^(.*\..*?)\/(.*):([^@]+)@?(.*)?`)
var profilePattern = regexp.MustCompile(`^(?:profile):(.*)`)
var filePattern = regexp.MustCompile(`^(?:file):(.*)`)

type ScriptFetcher func(is ImportableScript) (string, error)

type ImportableScript interface {
	Fetch(bool) (string, error)
}

type GitScript struct {
	host     string
	repo     string
	filePath string
	version  string
}

type ProfileScript struct {
	name string
}

type WebScript struct {
	u url.URL
}

type FileScript struct {
	filePath string
}

func (gs GitScript) Fetch(allowInsecureRequests bool) (string, error) {
	u := url.URL{}
	u.Scheme = "https"
	u.Host = gs.host
	u.Path = path.Join(gs.repo, "raw", gs.version, gs.filePath)
	return attemptToFetch(u, allowInsecureRequests)
}

func (ps ProfileScript) Fetch(allowInsecureRequests bool) (string, error) {
	return "", fmt.Errorf("no profile fetcher function provided")
}

func (ws WebScript) Fetch(allowInsecureRequests bool) (string, error) {
	return attemptToFetch(ws.u, allowInsecureRequests)
}

func (fs FileScript) Fetch(allowInsecureRequests bool) (string, error) {
	data, err := ioutil.ReadFile(fs.filePath)
	return string(data), err
}

func ParseImport(importStmt string) (ImportableScript, error) {
	if strings.HasPrefix(importStmt, "http") {
		u, err := url.Parse(importStmt)
		if err != nil {
			return nil, err
		}
		if u == nil {
			return nil, fmt.Errorf("url.Parse returned nil url")
		}
		ws := WebScript{
			u: *u,
		}
		return ws, nil
	} else if matches := gitPattern.FindAllStringSubmatch(importStmt, -1); len(matches) > 0 {
		if len(matches[0]) != 5 {
			return nil, fmt.Errorf("regular expression returned unexpected amount of submatches")
		}
		var version = matches[0][4]
		if version == "" {
			version = "HEAD"
		}
		gs := GitScript{
			host:     matches[0][1],
			repo:     matches[0][2],
			filePath: matches[0][3],
			version:  version,
		}
		return gs, nil
	} else if matches := profilePattern.FindAllStringSubmatch(importStmt, -1); len(matches) > 0 {
		if len(matches[0]) != 2 {
			return nil, fmt.Errorf("regular expression returned unexpected amount of submatches")
		}
		ps := ProfileScript{
			name: matches[0][1],
		}
		return ps, nil
	} else if matches := filePattern.FindAllStringSubmatch(importStmt, -1); len(matches) > 0 {
		if len(matches[0]) != 2 {
			return nil, fmt.Errorf("regular expression returned unexpected amount of submatches")
		}
		fs := FileScript{
			filePath: matches[0][1],
		}
		return fs, nil
	}
	return nil, fmt.Errorf("no matching import scheme")
}

func attemptToFetch(u url.URL, allowInsecureRequests bool) (string, error) {
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return "", err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		if u.Scheme == "https" && allowInsecureRequests {
			u.Scheme = "http"
			return attemptToFetch(u, allowInsecureRequests)
		} else {
			return "", fmt.Errorf("received non 2xx response code %d", resp.StatusCode)
		}
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
