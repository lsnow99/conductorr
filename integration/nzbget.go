package integration

import (
	"errors"
	"net/url"
	"strconv"

	"github.com/ethereum/go-ethereum/rpc"
)

type NZBGet struct {
	username  string
	password  string
	baseUrl   string
	rpcClient *rpc.Client
	downloads []Download
}

type DownloadEntry struct {
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
	ParStatus    string
	UnpackStatus string
	MoveStatus   string
	ScriptStatus []struct {
		Name   string
		Status string
	}
	DeleteStatus     string
	MarkStatus       string
	ScriptStatuses   string
	PostTotalTimeSec int
	ParTimeSec       int
	RepairTimeSec    int
	UnpackTimeSec    int
	PostInfoText     string
}

func NewNZBGet(username, password, baseUrl string) *NZBGet {
	n := &NZBGet{
		username: username,
		password: password,
		baseUrl:  baseUrl,
	}

	return n
}

func NewNZBGetFromConfig(configuration map[string]interface{}) (*NZBGet, error) {
	username, uOK := configuration["username"].(string)
	password, pOK := configuration["password"].(string)
	baseUrl, bOK := configuration["base_url"].(string)
	if !uOK || !pOK || !bOK {
		return nil, errors.New("failed to parse nzbget configuration")
	}

	return NewNZBGet(username, password, baseUrl), nil
}

func (n *NZBGet) Init() error {
	endpoint, _ := url.Parse(n.baseUrl)
	endpoint.Path = "jsonrpc"
	endpoint.User = url.UserPassword(n.username, n.password)

	rpcClient, err := rpc.Dial(endpoint.String())
	if err != nil {
		return err
	}
	n.rpcClient = rpcClient

	return nil
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

func (n *NZBGet) AddMedia(media *Media) error {
	var code, priority int

	if media.HighPriority {
		priority = 100
	}
	if err := n.rpcClient.Call(&code,
		"append",
		"",
		media.URL,
		media.Category,
		priority,
		media.HighPriority,
		false,
		"",
		0,
		"all",
		[]string{"conductorr", "true"}); err != nil {
		return err
	}

	if code <= 0 {
		return errors.New("error adding media to nzbget")
	}

	media.Identifier = strconv.Itoa(code)

	return nil
}

func (n *NZBGet) PollDownloads(names []string) error {
	entries := make([]DownloadEntry, 0)

	if err := n.rpcClient.Call(&entries, "listgroups"); err != nil {
		return err
	}

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
		switch entry.Status {
		case "PAUSED":
			d.Status = Paused
			break
		case "DOWNLOADING":
			d.Status = Downloading
			break
		case "FETCHING":
		case "QUEUED":
		case "PP_QUEUED":
			d.Status = Waiting
			break
		case "PP_FINISHED":
			d.Status = Done
			break
		default:
			d.Status = Processing
			break
		}
		d.FriendlyName = entry.NZBName
		d.BytesLeft = convertLoHi(entry.RemainingSizeLo, entry.RemainingSizeHi)
		d.FullSize = convertLoHi(entry.FileSizeLo, entry.FileSizeHi)
		downloads = append(downloads, d)
	}
	n.downloads = downloads
	return nil
}

func (n *NZBGet) GetDownloads() []Download {
	return n.downloads
}

func convertLoHi(lo, hi uint) uint64 {
	low := uint64(lo)
	high := uint64(hi)
	return (high << 32) | low
}
