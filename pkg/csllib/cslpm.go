package csllib

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

var gitPattern = regexp.MustCompile(`^(.*\..*?)\/(.*):([^@]+)@?(.*)?`)
var profilePattern = regexp.MustCompile(`^(?:profile):(filter|sorter):(.*?)(?:\?|$)(.*)`)
var filePattern = regexp.MustCompile(`^(?:file):(.*)`)

type basicAuthDetails struct {
	username string
	password string
}

type ScriptFetcher func(is ImportableScript, importPath string, allowInsecureRequests bool) (string, error)

type ImportableScript interface {
	Fetch(allowInsecureRequests bool) (script string, err error)
	CanonicalizedImportPath() string
}

type GitScript struct {
	host     string
	repo     string
	filePath string
	version  string
}

type ProfileScript struct {
	name       string
	scriptType string
	host       string
	authToken  string
	username   string
	password   string
	port       int
}

type WebScript struct {
	u url.URL
}

type FileScript struct {
	filePath string
}

func NewCSLPackageManager(scriptFetcher ScriptFetcher, allowInsecureRequests bool) *CSLPackageManager {
	cslpm := new(CSLPackageManager)
	if scriptFetcher == nil {
		cslpm.scriptFetcher = DefaultFetcher
	} else {
		cslpm.scriptFetcher = scriptFetcher
	}
	cslpm.allowInsecureRequests = allowInsecureRequests
	return cslpm
}

func (gs GitScript) Fetch(allowInsecureRequests bool) (string, error) {
	return attemptToFetch(gs.GetURL(), nil, nil, allowInsecureRequests)
}

func (gs GitScript) CanonicalizedImportPath() string {
	return fmt.Sprintf("%s/%s:%s@%s", gs.host, gs.repo, gs.filePath, gs.version)
}

func (gs GitScript) GetURL() (u url.URL) {
	u.Scheme = "https"
	u.Host = gs.host
	u.Path = path.Join(gs.repo, "raw", gs.version, gs.filePath)
	return
}

// Fetch for profile scripts assumes that the host and port have been specified, and
// that at least either the auth token or username and password are set.
// A request will first be made over https to the specified Conductorr instance, and
// upon failure, if allowInsecureRequests is set, http will be attempted as a fallback.
func (ps ProfileScript) Fetch(allowInsecureRequests bool) (string, error) {
	if ps.host == "" {
		return "", fmt.Errorf("no host provided for profile %s, cannot resolve", ps.name)
	}

	var creds *basicAuthDetails
	u := url.URL{}
	u.Scheme = "https"
	u.Host = fmt.Sprintf("%s:%d", ps.host, ps.port)
	header := http.Header{}
	u.Path = fmt.Sprintf("/api/v1/profile/%s/%s/raw", ps.name, ps.scriptType)

	if ps.authToken != "" {
		header.Set("X-Auth-Token", ps.authToken)
	} else if ps.username != "" && ps.password != "" {
		creds = &basicAuthDetails{
			username: ps.username,
			password: ps.password,
		}
	} else {
		return "", fmt.Errorf("no credentials provided")
	}
	script, err := attemptToFetch(u, header, creds, allowInsecureRequests)
	return script, err
}

func (ps ProfileScript) CanonicalizedImportPath() string {
	var queryStr string
	if ps.host != "" {
		vals := url.Values{}
		// Only set the values if either the auth token is set or both the password and username
		var valid bool
		if ps.authToken != "" {
			vals.Set("auth_token", ps.authToken)
			valid = true
		} else if ps.password != "" && ps.username != "" {
			vals.Set("password", ps.password)
			vals.Set("username", ps.username)
			valid = false
		}

		if valid {
			vals.Set("host", ps.host)
			vals.Set("port", strconv.Itoa(ps.port))
			queryStr = "?" + vals.Encode()
		}
	}
	return fmt.Sprintf("profile:%s:%s%s", ps.scriptType, ps.name, queryStr)
}

func (ps ProfileScript) GetName() string {
	return ps.name
}

func (ps ProfileScript) GetScriptType() string {
	return ps.scriptType
}

func (ws WebScript) Fetch(allowInsecureRequests bool) (string, error) {
	return attemptToFetch(ws.u, nil, nil, allowInsecureRequests)
}

func (ws WebScript) CanonicalizedImportPath() string {
	return ws.u.String()
}

func (ws WebScript) GetURL() url.URL {
	return ws.u
}

func (fs FileScript) Fetch(allowInsecureRequests bool) (string, error) {
	data, err := ioutil.ReadFile(fs.filePath)
	return string(data), err
}

func (fs FileScript) CanonicalizedImportPath() string {
	return fs.filePath
}

func (cslpm *CSLPackageManager) parseImport(importStmt string) (ImportableScript, error) {
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
			// Default to HEAD if no version has been specified
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
		if len(matches[0]) != 4 {
			return nil, fmt.Errorf("regular expression returned unexpected amount of submatches")
		}

		vals, err := url.ParseQuery(matches[0][3])
		if err != nil {
			return nil, err
		}

		var port = 8282
		portStr := vals.Get("port")
		if portStr != "" {
			port, err = strconv.Atoi(portStr)
			if err != nil {
				return nil, err
			}
		}

		ps := ProfileScript{
			name:       matches[0][2],
			scriptType: matches[0][1],
			host:       vals.Get("host"),
			port:       port,
			username:   vals.Get("username"),
			password:   vals.Get("password"),
			authToken:  vals.Get("auth_token"),
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
	} else {
		// Assume it is a file
		fs := FileScript{
			filePath: importStmt,
		}
		return fs, nil
	}
}

func (cslpm *CSLPackageManager) Resolve(importPath string) (string, error) {
	is, err := cslpm.parseImport(importPath)
	if err != nil {
		return "", err
	}
	return cslpm.scriptFetcher(is, importPath, cslpm.allowInsecureRequests)
}

func attemptToFetch(u url.URL, header http.Header, creds *basicAuthDetails, allowInsecureRequests bool) (string, error) {
	req, err := http.NewRequest("GET", u.String(), nil)
	req.Header = header
	if creds != nil {
		req.SetBasicAuth(creds.username, creds.password)
	}

	if err != nil {
		return "", err
	}
	resp, err := http.DefaultClient.Do(req)
	// Need to check the error before using resp
	if err != nil {
		if u.Scheme == "https" && allowInsecureRequests {
			// Fallback to http if allowed
			u.Scheme = "http"
			return attemptToFetch(u, header, creds, allowInsecureRequests)
		} else {
			// Otherwise return the error
			return "", err
		}
	}
	defer resp.Body.Close()

	if err != nil || (resp.StatusCode < 200 || resp.StatusCode > 299) {
		if u.Scheme == "https" && allowInsecureRequests {
			// Fallback to http if allowed
			u.Scheme = "http"
			return attemptToFetch(u, header, creds, allowInsecureRequests)
		} else {
			// Otherwise return the error
			return "", fmt.Errorf("received non 2xx response code %d", resp.StatusCode)
		}
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// DefaultFetcher implements the Fetcher interface and uses a simple filesystem cache to resolve scripts
func DefaultFetcher(is ImportableScript, importPath string, allowInsecureRequests bool) (string, error) {
	if _, ok := is.(FileScript); ok {
		return is.Fetch(allowInsecureRequests)
	}

	cacheDir, err := GetDefaultCacheDir()
	if err != nil {
		return "", err
	}

	filename := GetDefaultCacheName(importPath)
	filename = filepath.Join(cacheDir, filename)
	data, err := os.ReadFile(filename)
	if err == nil {
		return string(data), nil
	}

	script, err := is.Fetch(allowInsecureRequests)
	if err != nil {
		return "", err
	}

	return script, os.WriteFile(filename, []byte(script), os.ModePerm)
}

func GetDefaultCacheName(importPath string) string {
	hasher := sha256.New()
	hasher.Write([]byte(importPath))
	return hex.EncodeToString(hasher.Sum(nil)) + ".csl"
}

func GetDefaultCacheDir() (string, error) {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		return "", err
	}
	cacheDir = filepath.Join(cacheDir, "csl")
	if err := os.MkdirAll(cacheDir, os.ModePerm); err != nil {
		return "", err
	}
	return cacheDir, nil
}
