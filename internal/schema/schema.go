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

type APIResponse struct {
	Status   int64  `json:"status"`
	Action   string `json:"action"`
	Response string `json:"response"`
}

type Configuration struct {
	tableName            struct{} `pg:"public.\"Configuration\""`
	SonarrApiKey         string   `pg:"sonarr_api_key"`
	RadarrApiKey         string   `pg:"radarr_api_key"`
	RadarrCategory       string   `pg:"radarr_category"`
	RadarrUrl            string   `pg:"radarr_url"`
	SonarrCategory       string   `pg:"sonarr_category"`
	SonarrUrl            string   `pg:"sonarr_url"`
	FbOutputDir          string   `pg:"fb_output_dir"`
	FbSubtitlesLocale    string   `pg:"fb_subtitles_locale"`
	FbAction             string   `pg:"fb_action"`
	FbExcludeList        string   `pg:"fb_exclude_list"`
	FbAmcLog             string   `pg:"fb_amc_log"`
	FbKodi               string   `pg:"fb_kodi"`
	FbPlex               string   `pg:"fb_plex"`
	FbEmby               string   `pg:"fb_emby"`
	FbPushover           string   `pg:"fb_pushover"`
	FbPushbullet         string   `pg:"fb_pushbullet"`
	FbDiscord            string   `pg:"fb_discord"`
	FbGmail              string   `pg:"fb_gmail"`
	FbMail               string   `pg:"fb_mail"`
	FbMailto             string   `pg:"fb_mailto"`
	FbReportError        bool   `pg:"fb_report_error"`
	FbStoreReport        bool   `pg:"fb_store_report"`
	FbExtractFolder      string   `pg:"fb_extract_folder"`
	FbExec               string   `pg:"fb_exec"`
	FbIgnore             string   `pg:"fb_ignore"`
	FbMinFileSize        int64    `pg:"fb_min_file_size"`
	FbMinLengthMs        int64    `pg:"fb_min_length_ms"`
	ConfigurationId      int64    `pg:"configuration_id"`
	FbArtwork            bool     `pg:"fb_artwork"`
	FbSkipExtract        bool     `pg:"fb_skip_extract"`
	FbDeleteAfterExtract bool     `pg:"fb_delete_after_extract"`
	FbClean              bool     `pg:"fb_clean"`
	FbUnsorted           bool     `pg:"fb_unsorted"`
	FbHomeDir            string     `pg:"fb_home_dir"`
	FbExtras             bool     `pg:"fb_extras"`
	PlexNamespace		string		`pg:"plex_namespace"`
	PlexDeploymentName	string `pg:"plex_deployment_name"`
	PlexAuthToken		string `pg:"plex_auth_token"`
	PlexBaseUrl			string `pg:"plex_base_url"`
	FbNamespace			string `pg:"fb_namespace"`
	FbDeploymentName	string `pg:"fb_deployment_name"`
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
