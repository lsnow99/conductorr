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
	"strings"
)

var gitPattern = regexp.MustCompile(`^(.*\..*?)\/(.*):([^@]+)@?(.*)?`)
var profilePattern = regexp.MustCompile(`^(?:profile):(filter|sorter):(.*)`)
var filePattern = regexp.MustCompile(`^(?:file):(.*)`)

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
	name string
	scriptType string
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
	return attemptToFetch(gs.GetURL(), allowInsecureRequests)
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

func (ps ProfileScript) Fetch(allowInsecureRequests bool) (string, error) {
	return "", fmt.Errorf("no profile fetcher function provided")
}

func (ps ProfileScript) CanonicalizedImportPath() string {
	return fmt.Sprintf("profile:%s:%s", ps.scriptType, ps.name)
}

func (ps ProfileScript) GetName() string {
	return ps.name
}

func (ps ProfileScript) GetScriptType() string {
	return ps.scriptType
}

func (ws WebScript) Fetch(allowInsecureRequests bool) (string, error) {
	return attemptToFetch(ws.u, allowInsecureRequests)
}

func (ws WebScript) CanonicalizedImportPath() string {
	return ws.u.String()
}

func (ws WebScript) GetURL() (url.URL) {
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
		if len(matches[0]) != 3 {
			return nil, fmt.Errorf("regular expression returned unexpected amount of submatches")
		}
		ps := ProfileScript{
			name: matches[0][2],
			scriptType: matches[0][1],
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
	// return nil, fmt.Errorf("no matching import scheme")
}

func (cslpm *CSLPackageManager) Resolve(importPath string) (string, error) {
	is, err := cslpm.parseImport(importPath)
	if err != nil {
		return "", err
	}
	return cslpm.scriptFetcher(is, importPath, cslpm.allowInsecureRequests)
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
