package schema

import (
	"time"
)

// User database and JSON schema
type User struct {
	tableName    struct{}  `pg:"public.\"Users\""`
	UserId       int64     `json:"user_id" pg:",pk"`
	FirstName    string    `json:"first_name" pg:"first_name"`
	LastName     string    `json:"last_name" pg:"last_name"`
	Email        string    `json:"email" pg:"email"`
	Birthday     time.Time `json:"birthday" pg:"birthday"`
	Password     string    `pg:"-" json:"password"`
	PasswordHash []byte    `json:"password_hash" pg:"password_hash"`
}

type Credentials struct {
	Email        string `json:"email" pg:"email"`
	Password     string `pg:"-" json:"password"`
	PasswordHash []byte `json:"password_hash" pg:"password_hash"`
}

type SystemConfiguration struct {
	FbLogsEnabled	bool `pg:"fb_logs_enabled" json:"fb_logs_enabled"`
	PlexScannerLogsEnabled	bool `pg:"plex_scanner_logs_enabled" json:"plex_scanner_logs_enabled"`
}

// FilebotConfiguration database + json model
type FilebotConfiguration struct {
	FbOutputDir          string   `pg:"fb_output_dir" json:"fb_output_dir"`
	FbSubtitlesLocale    string   `pg:"fb_subtitles_locale" json:"fb_subtitles_locale"`
	FbAction             string   `pg:"fb_action" json:"fb_action"`
	FbExcludeList        string   `pg:"fb_exclude_list" json:"fb_exclude_list"`
	FbAmcLog             string   `pg:"fb_amc_log" json:"fb_amc_log"`
	FbKodi               string   `pg:"fb_kodi" json:"fb_kodi"`
	FbPlex               string   `pg:"fb_plex" json:"fb_plex"`
	FbEmby               string   `pg:"fb_emby" json:"fb_emby"`
	FbPushover           string   `pg:"fb_pushover" json:"fb_pushover"`
	FbPushbullet         string   `pg:"fb_pushbullet" json:"fb_pushbullet"`
	FbDiscord            string   `pg:"fb_discord" json:"fb_discord"`
	FbGmail              string   `pg:"fb_gmail" json:"fb_gmail"`
	FbMail               string   `pg:"fb_mail" json:"fb_mail"`
	FbMailto             string   `pg:"fb_mailto" json:"fb_mailto"`
	FbExtractFolder      string   `pg:"fb_extract_folder" json:"fb_extract_folder"`
	FbExec               string   `pg:"fb_exec" json:"fb_exec"`
	FbIgnore             string   `pg:"fb_ignore" json:"fb_ignore"`
	FbHomeDir            string   `pg:"fb_home_dir" json:"fb_home_dir"`
	FbNamespace			 string   `pg:"fb_namespace" json:"fb_namespace"`
	FbDeploymentName	 string   `pg:"fb_deployment_name" json:"fb_deployment_name"`
	FbMinFileSize        int64    `pg:"fb_min_file_size" json:"fb_min_file_size"`
	FbMinLengthMs        int64	  `pg:"fb_min_length_ms" json:"fb_min_length_ms"`
	FbArtwork            bool     `pg:"fb_artwork" json:"fb_artwork"`
	FbSkipExtract        bool     `pg:"fb_skip_extract" json:"fb_skip_extract"`
	FbDeleteAfterExtract bool     `pg:"fb_delete_after_extract" json:"fb_delete_after_extract"`
	FbClean              bool     `pg:"fb_clean" json:"fb_clean"`
	FbUnsorted           bool     `pg:"fb_unsorted" json:"fb_unsorted"`
	FbReportError        bool     `pg:"fb_report_error" json:"fb_report_error"`
	FbStoreReport        bool     `pg:"fb_store_report" json:"fb_store_report"`
	FbExtras             bool	  `pg:"fb_extras" json:"fb_extras"`
	FilebotConfigurationID    bool   `pg:"filebot_configuration_id"`
}

type Configuration struct {
	tableName            struct{} `pg:"public.\"Configuration\""`
	SonarrApiKey         string   `pg:"sonarr_api_key"`
	RadarrApiKey         string   `pg:"radarr_api_key"`
	RadarrCategory       string   `pg:"radarr_category"`
	RadarrUrl            string   `pg:"radarr_url"`
	SonarrCategory       string   `pg:"sonarr_category"`
	SonarrUrl            string   `pg:"sonarr_url"`
	PlexNamespace		string		`pg:"plex_namespace"`
	PlexDeploymentName	string `pg:"plex_deployment_name"`
	PlexAuthToken		string `pg:"plex_auth_token"`
	PlexBaseUrl			string `pg:"plex_base_url"`
}

type Configurator struct {
	Label string `json:"label"`
	Placeholder string `json:"placeholder"`
	Description string `json:"description"`
	Property string `json:"property"`
	BoolValue bool `json:"bool_value"`
	IntValue int64 `json:"int_value"`
	StringValue string `json:"string_value"`
	Datatype string `json:"datatype"`
}
