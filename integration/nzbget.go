package integration

import (
	"encoding/base64"
	"errors"
	"io/ioutil"
	"mime"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/rpc"
	"github.com/lsnow99/conductorr/constant"
)

type NZBGet struct {
	username  string
	password  string
	baseUrl   string
	rpcClient *rpc.Client
	downloads []Download
}

type ItemEntry struct {
	NZBID              int
	NZBFilename        string
	NZBName            string
	Kind               string
	URL                string
	DestDir            string
	FinalDir           string
	Category           string
	FileSizeLo         uint
	FileSizeHi         uint
	FileSizeMB         int
	RemainingSizeLo    uint
	RemainingSizeHi    uint
	RemainingSizeMB    int
	PausedSizeLo       uint
	PausedSizeHi       uint
	PausedSizeMB       int
	FileCount          int
	RemainingFileCount int
	RemainingParCount  int
	MinPostTime        int
	MaxPostTime        int
	MaxPriority        int
	ActiveDownloads    int
	Status             string
	TotalArticles      int
	SuccessArticles    int
	FailedArticles     int
	Health             int
	CriticalHealth     int
	DownloadedSizeLo   uint
	DownlaodedSizeHi   uint
	DownloadedSizeMB   int
	DownloadTimeSec    int
	MessageCount       int
	DupeKey            string
	DupeScore          int
	DupeMode           string
	Parameters         []struct {
		Name  string
		Value string
	}
	ParStatus        string
	UnpackStatus     string
	MoveStatus       string
	DeleteStatus     string
	MarkStatus       string
	PostTotalTimeSec int
	ParTimeSec       int
	RepairTimeSec    int
	UnpackTimeSec    int
	PostInfoText     string
}

func NewNZBGet(username, password, baseUrl string) (*NZBGet, error) {
	n := &NZBGet{
		username: username,
		password: password,
		baseUrl:  baseUrl,
	}

	endpoint, err := url.Parse(baseUrl)
	if err != nil {
		return nil, err
	}
	endpoint.Path = "jsonrpc"
	endpoint.User = url.UserPassword(username, password)

	rpcClient, err := rpc.Dial(endpoint.String())
	if err != nil {
		return nil, err
	}
	n.rpcClient = rpcClient

	return n, nil
}

func NewNZBGetFromConfig(configuration map[string]interface{}) (*NZBGet, error) {
	username, uOK := configuration["username"].(string)
	password, pOK := configuration["password"].(string)
	baseUrl, bOK := configuration["base_url"].(string)
	if !uOK || !pOK || !bOK {
		return nil, errors.New("failed to parse nzbget configuration")
	}

	return NewNZBGet(username, password, baseUrl)
}

// func (n *NZBGet) GetConfig() string {
// 	cfg := NZBGetConfig{
// 		Username: n.username,
// 		Password: n.password,
// 		URL:      n.baseUrl,
// 	}
// 	data, _ := json.Marshal(cfg)

// 	return string(data)
// }

func (n *NZBGet) TestConnection() error {
	var ver string
	return n.rpcClient.Call(&ver, "version")
}

func (n *NZBGet) AddRelease(release Release) (string, error) {
	var code, priority int

	if release.HighPriority {
		priority = 100
	}

	resp, err := http.Get(release.DownloadURL)
	if err != nil {
		return "", err
	}

	_, params, err := mime.ParseMediaType(resp.Header.Get("Content-Disposition"))
	if err != nil {
		return "", err
	}

	filename := params["filename"]

	bod, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	nzbContent := base64.StdEncoding.EncodeToString(bod)

	if err := n.rpcClient.Call(&code,
		"append",
		filename,
		nzbContent,
		"",
		priority,
		release.HighPriority,
		false,
		"",
		0,
		"all",
		[]string{"conductorr", "true"}); err != nil {
		return "", err
	}

	if code <= 0 {
		return "", errors.New("error adding media to nzbget")
	}

	return strconv.Itoa(code), nil
}

func (n *NZBGet) PollDownloads(names []string) ([]Download, error) {
	entries := make([]ItemEntry, 0)

	if err := n.rpcClient.Call(&entries, "listgroups"); err != nil {
		return nil, err
	}

	history := make([]ItemEntry, 0)
	if err := n.rpcClient.Call(&history, "history"); err != nil {
		return nil, err
	}

	entries = append(entries, history...)

	downloads := make([]Download, 0)
	for _, entry := range entries {
		found := false
		for _, p := range entry.Parameters {
			if p.Name == "conductorr" {
				found = true
			}
		}
		if !found {
			continue
		}

		d := Download{}
		if contains([]string{"PAUSED"}, entry.Status) {
			d.Status = constant.StatusPaused
		} else if contains([]string{"DOWNLOADING"}, entry.Status) {
			d.Status = constant.StatusDownloading
		} else if contains([]string{"FETCHING", "QUEUED", "PP_QUEUED"}, entry.Status) {
			d.Status = constant.StatusWaiting
		} else if strings.Contains(entry.Status, "SUCCESS") {
			d.Status = constant.StatusComplete
		} else if strings.Contains(entry.Status, "FAILURE") || strings.Contains(entry.Status, "WARNING") || strings.Contains(entry.Status, "DELETED") && !strings.Contains(entry.Status, "DELETED/COPY") {
			d.Status = constant.StatusError
		} else {
			d.Status = constant.StatusProcessing
		}

		if entry.FinalDir != "" {
			d.FinalDir = entry.FinalDir
		} else {
			d.FinalDir = entry.DestDir
		}
		d.FriendlyName = entry.NZBName
		d.BytesLeft = convertLoHi(entry.RemainingSizeLo, entry.RemainingSizeHi)
		d.FullSize = convertLoHi(entry.FileSizeLo, entry.FileSizeHi)
		d.Identifier = entry.NZBFilename

		downloads = append(downloads, d)
	}
	
	return downloads, nil
}

func contains(arr []string, s string) bool {
	for _, str := range arr {
		if s == str {
			return true
		}
	}
	return false
}

func convertLoHi(lo, hi uint) uint64 {
	low := uint64(lo)
	high := uint64(hi)
	return (high << 32) | low
}
