package schema

import (
	"encoding/xml"
	"time"
)

// Credentials receive credentials from login page
type Credentials struct {
	Username     string `json:"username" pg:"username"`
	Password     string `pg:"-" json:"password"`
	PasswordHash []byte `json:"password_hash" pg:"password_hash"`
}

// SystemConfiguration database + json model
type SystemConfiguration struct {
	tableName              struct{} `pg:"system_configuration"`
	FbLogsEnabled          bool     `pg:"filebot_logs_enabled" json:"filebot_logs_enabled"`
	PlexScannerLogsEnabled bool     `pg:"plex_scanner_logs_enabled" json:"plex_scanner_logs_enabled"`
	Username               string   `pg:"username" json:"username"`
	Password               []byte   `pg:"password"`
	SystemConfigurationID  bool     `pg:"system_configuration_id, pk" json:"system_configuration_id"`
}

// Jobs database + json model
type Jobs struct {
	tableName          struct{}  `pg:"jobs"`
	GrabberInternalID  int64     `pg:"grabber_internal_id"`
	JobID              int64     `pg:"job_id, pk"`
	TorrentLinkerID    string    `pg:"torrent_linker_id"`
	NZBLinkerID        string    `pg:"nzb_linker_id"`
	TimeGrabbed        time.Time `pg:"time_grabbed"`
	Title              string    `pg:"title"`
	ImdbID             string    `pg:"imdb_id"`
	ReleaseTitle       string    `pg:"release_title"`
	ContentType        string    `pg:"content_type"`
	DownloadClient     string    `pg:"download_client"`
	DownloadDirectory  string    `pg:"download_directory"`
	Status             string    `pg:"status"`
	FilebotLogs        string    `pg:"filebot_logs"`
	ScannerLogs        string    `pg:"scanner_logs"`
	GrabbedQuality     string    `pg:"grabbed_quality"`
	GrabbedSize        string    `pg:"grabbed_size"`
	TimeFilebotStarted time.Time `pg:"time_filebot_started"`
	TimeFilebotDone    time.Time `pg:"time_filebot_done"`
	TimeScanStarted    time.Time `pg:"time_scan_started"`
	TimeScanDone       time.Time `pg:"time_scan_done"`
}

// FilebotConfiguration database + json model
type FilebotConfiguration struct {
	tableName              struct{} `pg:"filebot_configuration"`
	FbOutputDir            string   `pg:"fb_output_dir" json:"fb_output_dir"`
	FbSubtitlesLocale      string   `pg:"fb_subtitles_locale" json:"fb_subtitles_locale"`
	FbAction               string   `pg:"fb_action" json:"fb_action"`
	FbExcludeList          string   `pg:"fb_exclude_list" json:"fb_exclude_list"`
	FbAmcLog               string   `pg:"fb_amc_log" json:"fb_amc_log"`
	FbKodi                 string   `pg:"fb_kodi" json:"fb_kodi"`
	FbPlex                 string   `pg:"fb_plex" json:"fb_plex"`
	FbEmby                 string   `pg:"fb_emby" json:"fb_emby"`
	FbPushover             string   `pg:"fb_pushover" json:"fb_pushover"`
	FbPushbullet           string   `pg:"fb_pushbullet" json:"fb_pushbullet"`
	FbDiscord              string   `pg:"fb_discord" json:"fb_discord"`
	FbGmail                string   `pg:"fb_gmail" json:"fb_gmail"`
	FbMail                 string   `pg:"fb_mail" json:"fb_mail"`
	FbMailto               string   `pg:"fb_mailto" json:"fb_mailto"`
	FbExtractFolder        string   `pg:"fb_extract_folder" json:"fb_extract_folder"`
	FbExec                 string   `pg:"fb_exec" json:"fb_exec"`
	FbIgnore               string   `pg:"fb_ignore" json:"fb_ignore"`
	FbHomeDir              string   `pg:"fb_home_dir" json:"fb_home_dir"`
	FbNamespace            string   `pg:"fb_namespace" json:"fb_namespace"`
	FbDeploymentName       string   `pg:"fb_deployment_name" json:"fb_deployment_name"`
	FbMinFileSize          int64    `pg:"fb_min_file_size" json:"fb_min_file_size"`
	FbMinLengthMs          int64    `pg:"fb_min_length_ms" json:"fb_min_length_ms"`
	FbArtwork              bool     `pg:"fb_artwork" json:"fb_artwork"`
	FbSkipExtract          bool     `pg:"fb_skip_extract" json:"fb_skip_extract"`
	FbDeleteAfterExtract   bool     `pg:"fb_delete_after_extract" json:"fb_delete_after_extract"`
	FbClean                bool     `pg:"fb_clean" json:"fb_clean"`
	FbUnsorted             bool     `pg:"fb_unsorted" json:"fb_unsorted"`
	FbReportError          bool     `pg:"fb_report_error" json:"fb_report_error"`
	FbStoreReport          bool     `pg:"fb_store_report" json:"fb_store_report"`
	FbExtras               bool     `pg:"fb_extras" json:"fb_extras"`
	FilebotConfigurationID bool     `pg:"filebot_configuration_id, pk"`
}

// PlexConfiguration database + json model
type PlexConfiguration struct {
	tableName           struct{} `pg:"plex_configuration"`
	PlexNamespace       string   `pg:"plex_namespace" json:"plex_namespace"`
	PlexDeploymentName  string   `pg:"plex_deployment_name" json:"plex_deployment_name"`
	PlexAuthToken       string   `pg:"plex_auth_token" json:"plex_auth_token"`
	PlexBaseURL         string   `pg:"plex_base_url" json:"plex_base_url"`
	PlexConfigurationID bool     `pg:"plex_configuration_id, pk"`
}

// SonarrConfiguration database + json model
type SonarrConfiguration struct {
	tableName             struct{} `pg:"sonarr_configuration"`
	SonarrAPIKey          string   `pg:"sonarr_api_key" json:"sonarr_api_key"`
	SonarrCategory        string   `pg:"sonarr_category" json:"sonarr_category"`
	SonarrURL             string   `pg:"sonarr_url" json:"sonarr_url"`
	SonarrConfigurationID bool     `pg:"sonarr_configuration_id, pk"`
}

// RadarrConfiguration database + json model
type RadarrConfiguration struct {
	tableName             struct{} `pg:"radarr_configuration"`
	RadarrAPIKey          string   `pg:"radarr_api_key" json:"radarr_api_key"`
	RadarrCategory        string   `pg:"radarr_category" json:"radarr_category"`
	RadarrURL             string   `pg:"radarr_url" json:"radarr_url"`
	RadarrConfigurationID bool     `pg:"radarr_configuration_id, pk"`
}

// Configurator json model to transfer configuration fields to ui
type Configurator struct {
	Label       string `json:"label"`
	Placeholder string `json:"placeholder"`
	Description string `json:"description"`
	Property    string `json:"property"`
	BoolValue   bool   `json:"bool_value"`
	IntValue    int64  `json:"int_value"`
	StringValue string `json:"string_value"`
	Datatype    string `json:"datatype"`
}

// LinkRequest json model for receiving a request to link a newly grabbed item
type LinkRequest struct {
	ContentIdentifier        int64  `json:"content_id"`
	TorrentLinkerID          string `json:"torrent_linker_id"`
	NZBLinkerID              string `json:"nzb_linker_id"`
	ContentCategory          string `json:"content_category"`
	ContentName              string `json:"content_name"`
	ReleaseTitle             string `json:"release_title"`
	DownloadClientIdentifier string `json:"dl_client_id"`
}

// ImportRequest json model for receiving a request to run filebot on a download
type ImportRequest struct {
	DownloadDirectory        string `json:"dl_dir"`
	DownloadContentID        string `json:"dl_content_id"`
	ContentName              string `json:"content_name"`
	DownloadClientIdentifier string `json:"dl_client_id"`
}

/*
MediaContainer matches the response format from a Plex /library/sections API
request (generated at: https://www.onlinetool.io/xmltogo/)
*/
type MediaContainer struct {
	XMLName         xml.Name `xml:"MediaContainer"`
	Text            string   `xml:",chardata"`
	Size            string   `xml:"size,attr"`
	AllowSync       string   `xml:"allowSync,attr"`
	Identifier      string   `xml:"identifier,attr"`
	MediaTagPrefix  string   `xml:"mediaTagPrefix,attr"`
	MediaTagVersion string   `xml:"mediaTagVersion,attr"`
	Title1          string   `xml:"title1,attr"`
	Directory       []struct {
		Text                string `xml:",chardata"`
		AllowSync           string `xml:"allowSync,attr"`
		Art                 string `xml:"art,attr"`
		Composite           string `xml:"composite,attr"`
		Filters             string `xml:"filters,attr"`
		Refreshing          string `xml:"refreshing,attr"`
		Thumb               string `xml:"thumb,attr"`
		Key                 string `xml:"key,attr"`
		Type                string `xml:"type,attr"`
		Title               string `xml:"title,attr"`
		Agent               string `xml:"agent,attr"`
		Scanner             string `xml:"scanner,attr"`
		Language            string `xml:"language,attr"`
		UUID                string `xml:"uuid,attr"`
		UpdatedAt           string `xml:"updatedAt,attr"`
		CreatedAt           string `xml:"createdAt,attr"`
		ScannedAt           string `xml:"scannedAt,attr"`
		EnableAutoPhotoTags string `xml:"enableAutoPhotoTags,attr"`
		Location            struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Path string `xml:"path,attr"`
		} `xml:"Location"`
	} `xml:"Directory"`
}

// History xml model matching filebot history.xml schema
type History struct {
	XMLName   xml.Name   `xml:"history"`
	Sequences []Sequence `xml:"sequence"`
}

// Sequence xml model matching filebot history.xml schema
type Sequence struct {
	XMLName xml.Name `xml:"sequence"`
	Date    string   `xml:"date,attr"`
	Renames []Rename `xml:"rename"`
}

// Rename xml model matching filebot history.xml schema
type Rename struct {
	XMLName xml.Name `xml:"rename"`
	Dir     string   `xml:"dir,attr"`
	From    string   `xml:"from,attr"`
	To      string   `xml:"to,attr"`
}
