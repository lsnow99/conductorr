package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"time"
)

type Jellyfin struct {
	apiKey  string
	baseUrl string
}

type JellyfinUsers []struct {
	Name                      string    `json:"Name"`
	ServerID                  string    `json:"ServerId"`
	ID                        string    `json:"Id"`
	HasPassword               bool      `json:"HasPassword"`
	HasConfiguredPassword     bool      `json:"HasConfiguredPassword"`
	HasConfiguredEasyPassword bool      `json:"HasConfiguredEasyPassword"`
	EnableAutoLogin           bool      `json:"EnableAutoLogin"`
	LastLoginDate             time.Time `json:"LastLoginDate"`
	LastActivityDate          time.Time `json:"LastActivityDate"`
	Configuration             struct {
		PlayDefaultAudioTrack      bool          `json:"PlayDefaultAudioTrack"`
		SubtitleLanguagePreference string        `json:"SubtitleLanguagePreference"`
		DisplayMissingEpisodes     bool          `json:"DisplayMissingEpisodes"`
		GroupedFolders             []interface{} `json:"GroupedFolders"`
		SubtitleMode               string        `json:"SubtitleMode"`
		DisplayCollectionsView     bool          `json:"DisplayCollectionsView"`
		EnableLocalPassword        bool          `json:"EnableLocalPassword"`
		OrderedViews               []interface{} `json:"OrderedViews"`
		LatestItemsExcludes        []interface{} `json:"LatestItemsExcludes"`
		MyMediaExcludes            []interface{} `json:"MyMediaExcludes"`
		HidePlayedInLatest         bool          `json:"HidePlayedInLatest"`
		RememberAudioSelections    bool          `json:"RememberAudioSelections"`
		RememberSubtitleSelections bool          `json:"RememberSubtitleSelections"`
		EnableNextEpisodeAutoPlay  bool          `json:"EnableNextEpisodeAutoPlay"`
	} `json:"Configuration"`
	Policy struct {
		IsAdministrator                  bool          `json:"IsAdministrator"`
		IsHidden                         bool          `json:"IsHidden"`
		IsDisabled                       bool          `json:"IsDisabled"`
		BlockedTags                      []interface{} `json:"BlockedTags"`
		EnableUserPreferenceAccess       bool          `json:"EnableUserPreferenceAccess"`
		AccessSchedules                  []interface{} `json:"AccessSchedules"`
		BlockUnratedItems                []interface{} `json:"BlockUnratedItems"`
		EnableRemoteControlOfOtherUsers  bool          `json:"EnableRemoteControlOfOtherUsers"`
		EnableSharedDeviceControl        bool          `json:"EnableSharedDeviceControl"`
		EnableRemoteAccess               bool          `json:"EnableRemoteAccess"`
		EnableLiveTvManagement           bool          `json:"EnableLiveTvManagement"`
		EnableLiveTvAccess               bool          `json:"EnableLiveTvAccess"`
		EnableMediaPlayback              bool          `json:"EnableMediaPlayback"`
		EnableAudioPlaybackTranscoding   bool          `json:"EnableAudioPlaybackTranscoding"`
		EnableVideoPlaybackTranscoding   bool          `json:"EnableVideoPlaybackTranscoding"`
		EnablePlaybackRemuxing           bool          `json:"EnablePlaybackRemuxing"`
		ForceRemoteSourceTranscoding     bool          `json:"ForceRemoteSourceTranscoding"`
		EnableContentDeletion            bool          `json:"EnableContentDeletion"`
		EnableContentDeletionFromFolders []interface{} `json:"EnableContentDeletionFromFolders"`
		EnableContentDownloading         bool          `json:"EnableContentDownloading"`
		EnableSyncTranscoding            bool          `json:"EnableSyncTranscoding"`
		EnableMediaConversion            bool          `json:"EnableMediaConversion"`
		EnabledDevices                   []interface{} `json:"EnabledDevices"`
		EnableAllDevices                 bool          `json:"EnableAllDevices"`
		EnabledChannels                  []interface{} `json:"EnabledChannels"`
		EnableAllChannels                bool          `json:"EnableAllChannels"`
		EnabledFolders                   []interface{} `json:"EnabledFolders"`
		EnableAllFolders                 bool          `json:"EnableAllFolders"`
		InvalidLoginAttemptCount         int           `json:"InvalidLoginAttemptCount"`
		LoginAttemptsBeforeLockout       int           `json:"LoginAttemptsBeforeLockout"`
		MaxActiveSessions                int           `json:"MaxActiveSessions"`
		EnablePublicSharing              bool          `json:"EnablePublicSharing"`
		BlockedMediaFolders              []interface{} `json:"BlockedMediaFolders"`
		BlockedChannels                  []interface{} `json:"BlockedChannels"`
		RemoteClientBitrateLimit         int           `json:"RemoteClientBitrateLimit"`
		AuthenticationProviderID         string        `json:"AuthenticationProviderId"`
		PasswordResetProviderID          string        `json:"PasswordResetProviderId"`
		SyncPlayAccess                   string        `json:"SyncPlayAccess"`
	} `json:"Policy"`
}

type JellyfinItemsResponse struct {
	Items []struct {
		Name              string      `json:"Name"`
		ServerID          string      `json:"ServerId"`
		ID                string      `json:"Id"`
		HasSubtitles      bool        `json:"HasSubtitles,omitempty"`
		Container         string      `json:"Container"`
		PremiereDate      time.Time   `json:"PremiereDate"`
		Path              string      `json:"Path"`
		OfficialRating    string      `json:"OfficialRating,omitempty"`
		ChannelID         interface{} `json:"ChannelId"`
		CommunityRating   float64     `json:"CommunityRating,omitempty"`
		RunTimeTicks      int64       `json:"RunTimeTicks"`
		ProductionYear    int         `json:"ProductionYear"`
		IndexNumber       int         `json:"IndexNumber"`
		ParentIndexNumber int         `json:"ParentIndexNumber"`
		IsFolder          bool        `json:"IsFolder"`
		Type              string      `json:"Type"`
		UserData          struct {
			PlaybackPositionTicks int    `json:"PlaybackPositionTicks"`
			PlayCount             int    `json:"PlayCount"`
			IsFavorite            bool   `json:"IsFavorite"`
			Played                bool   `json:"Played"`
			Key                   string `json:"Key"`
		} `json:"UserData"`
		SeriesName            string `json:"SeriesName"`
		SeriesID              string `json:"SeriesId"`
		SeasonID              string `json:"SeasonId"`
		SeriesPrimaryImageTag string `json:"SeriesPrimaryImageTag"`
		SeasonName            string `json:"SeasonName"`
		VideoType             string `json:"VideoType"`
		ImageTags             struct {
			Primary string `json:"Primary"`
		} `json:"ImageTags"`
		BackdropImageTags       []interface{} `json:"BackdropImageTags"`
		LocationType            string        `json:"LocationType"`
		MediaType               string        `json:"MediaType"`
		ParentBackdropItemID    string        `json:"ParentBackdropItemId,omitempty"`
		ParentBackdropImageTags []string      `json:"ParentBackdropImageTags,omitempty"`
	} `json:"Items"`
}

func NewJellyfin(apiKey, baseUrl string) *Jellyfin {
	p := &Jellyfin{
		apiKey:  apiKey,
		baseUrl: baseUrl,
	}

	return p
}

func NewJellyfinFromConfig(configuration map[string]interface{}) (*Jellyfin, error) {
	apiKey, aOK := configuration["api_key"].(string)
	baseUrl, bOK := configuration["base_url"].(string)
	if !aOK || !bOK {
		return nil, fmt.Errorf("failed to parse jellyfin configuration")
	}

	return NewJellyfin(apiKey, baseUrl), nil
}

func (j *Jellyfin) TestConnection() error {
	u, err := url.Parse(j.baseUrl)
	if err != nil {
		return err
	}
	u.Path = path.Join(u.Path, "/System/Configuration")

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return err
	}
	req.Header.Add("X-Emby-Token", j.apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if len(data) == 0 {
		return fmt.Errorf("could not retrieve jellyfin configuration")
	}

	return nil
}

func (j *Jellyfin) ImportMedia(mediaPath string) error {
	u, err := url.Parse(j.baseUrl)
	if err != nil {
		return err
	}
	u.Path = path.Join(u.Path, "/Library/Media/Updated")

	payloadData := map[string]string{
		"Path": mediaPath,
	}

	payload, err := json.Marshal(payloadData)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", u.String(), bytes.NewReader(payload))
	if err != nil {
		return err
	}
	req.Header.Add("X-Emby-Token", j.apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func (j *Jellyfin) GetMediaPaths() ([]MediaPath, error) {
	users, err := j.getUsers()
	if err != nil {
		return nil, err
	}

	chosenUserID := ""
	for _, user := range users {
		if user.Policy.IsAdministrator {
			chosenUserID = user.ID
		}
	}

	u, err := url.Parse(j.baseUrl)
	if err != nil {
		return nil, err
	}
	q := url.Values{}
	q.Set("includeItemTypes", "Movie,Episode")
	q.Set("mediaType", "Video")
	q.Set("recursive", "true")
	q.Set("userId", chosenUserID)
	q.Set("fields", "Path")
	u.RawQuery = q.Encode()
	u.Path = path.Join(u.Path, "/Items")

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-Emby-Token", j.apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	itemsResp := JellyfinItemsResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&itemsResp); err != nil {
		return nil, err
	}

	mPaths := make([]MediaPath, 0, len(itemsResp.Items))
	for _, item := range itemsResp.Items {
		mPaths = append(mPaths, MediaPath{
			Title:            item.Name,
			ParentTitle:      item.SeasonName,
			GrandparentTitle: item.SeriesName,
			Path:             item.Path,
		})
	}

	return mPaths, nil
}

func (j *Jellyfin) getUsers() (JellyfinUsers, error) {
	u, err := url.Parse(j.baseUrl)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, "/Users")

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-Emby-Token", j.apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	users := JellyfinUsers{}
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		return nil, err
	}

	return users, nil
}
