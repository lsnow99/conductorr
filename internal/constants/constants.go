package constants

import "github.com/lsnow2017/conductorr/internal/schema"

var FilebotConfig = []schema.Configurator {
	{Label:"Output Directory", Placeholder:"/home/plex/Library", Description:"The root folder for your libraries. Filebot will create a Movies and TV Shows directory within it", Property:"fb_output_dir", StringValue:"", Datatype:"string"},
	{Label:"Subtitles Locale", Placeholder:"en", Description:"Subtitle locale(s) (ex: en,fr,de)", Property:"fb_subtitles_locale", StringValue:"en", Datatype:"string"},
	{Label:"Action", Placeholder:"copy", Description:"Filebot action to take (move, copy, duplicate)", Property:"fb_action", StringValue:"copy", Datatype:"string"},
	{Label:"Excludes File", Placeholder:".excludes", Description:"Excludes file name", Property:"fb_exclude_list", StringValue:".excludes", Datatype:"string"},
	{Label:"AMC Log", Placeholder:"amc.log", Description:"AMC script log file", Property:"fb_amc_log", StringValue:"amc.log", Datatype:"string"},
	{Label:"Kodi", Placeholder:"host[:port]", Description:"Kodi connection to trigger rescan", Property:"fb_kodi", StringValue:"", Datatype:"string"},
	{Label:"Plex", Placeholder:"host:token", Description:"Plex connection to trigger rescan. It is recommended to do this via Conductorr, but Filebot has this option as well", Property:"fb_plex", StringValue:"", Datatype:"string"},
	{Label:"Emby", Placeholder:"host:apikey", Description:"Emby connection to trigger rescan", Property:"fb_emby", StringValue:"", Datatype:"string"},
	{Label:"Pushover", Placeholder:"user[:apikey]", Description:"Send update notifications to your devices via Pushover. It is recommended that you create your own API key.", Property:"fb_pushover", StringValue:"", Datatype:"string"},
	{Label:"Pushbullet", Placeholder:"apikey", Description:"Send full reports to all your PushBullet devices", Property:"fb_pushbullet", StringValue:"", Datatype:"string"},
	{Label:"Discord", Placeholder:"webhook", Description:"Send update notifications to your Discord channel", Property:"fb_discord", StringValue:"", Datatype:"string"},
	{Label:"Gmail", Placeholder:"username:password", Description:"Use the following gmail account to send and receive full reports. You must use an App Password for security reasons.", Property:"fb_gmail", StringValue:"", Datatype:"string"},
	{Label:"Mail", Placeholder:"host:port:from[:username:password]", Description:"Send email via private mail server", Property:"fb_mail", StringValue:"", Datatype:"string"},
	{Label:"Mailto", Placeholder:"email", Description:"Send report to the given email address", Property:"fb_mailto", StringValue:"", Datatype:"string"},
	{Label:"Report Error", Placeholder:"", Description:"Report errors via email", Property:"fb_report_error", StringValue:"", BoolValue:true, Datatype:"bool"},
	{Label:"Store Report", Placeholder:"", Description:"Save reports to local filesystem", Property:"fb_store_report", StringValue:"", BoolValue:true, Datatype:"bool"},
	{Label:"Extract Folder", Placeholder:"/path/to/folder", Description:"Extract archives to this folder", Property:"fb_extract_folder", StringValue:"", Datatype:"string"},
	{Label:"Skip Extract", Placeholder:"", Description:"Do not extract archives", Property:"fb_skip_extract", StringValue:"", BoolValue:true, Datatype:"bool"},
	{Label:"Delete After Extract", Placeholder:"", Description:"Delete archives after extraction", Property:"fb_delete_after_extract", StringValue:"", BoolValue:true, Datatype:"bool"},
	{Label:"Clean", Placeholder:"", Description:"Automatically remove empty folders and clutter files that may be left behind after moving the video files, or temporary extracted files after copying", Property:"fb_clean", StringValue:"", BoolValue:true, Datatype:"bool"},
	{Label:"Unsorted", Placeholder:"", Description:"Move media files that cannot be identified to a separate folder", Property:"fb_unsorted", StringValue:"", BoolValue:true, Datatype:"bool"},
	{Label:"Ignore", Placeholder:"regex", Description:"Ignore filepaths that match the given regular expression", Property:"fb_ignore", StringValue:"", Datatype:"string"},
	{Label:"Min File Size", Placeholder:"0", Description:"Only process video files larger than the given number (in bytes). Defaults to 50 MB (use value -1 for default)", Property:"fb_min_file_size", StringValue:"", IntValue:-1, Datatype:"int"},
	{Label:"Min File Length (ms)", Placeholder:"0", Description:"Only process videos longer than the given number (in milliseconds). Defaults to 10 minutes (use value -1 for default)", Property:"fb_min_length_ms", StringValue:"", IntValue:-1, Datatype:"int"},
}

var SonarrConfig = []schema.Configurator {
	{Label:"Sonarr URL", Placeholder:"http://sonarr.cluster-ip.local:8989", Description:"Full URL to communicate with Sonarr", Property:"sonarr_url", StringValue:"", Datatype:"string"},
	{Label:"Sonarr API Key", Placeholder:"api key", Description:"Sonarr API Key, found in the Sonarr General Settings", Property:"sonarr_api_key", StringValue:"", Datatype:"string"},
	{Label:"Sonarr Category", Placeholder:"Series", Description:"Category in Sonarr settings that is sent to download clients (Must be set!)", Property:"sonarr_category", Datatype:"string"},
}

var RadarrConfig = []schema.Configurator {
	{Label:"Radarr URL", Placeholder:"http://radarr.cluster-ip.local:8989", Description:"Full URL to communicate with Radarr", Property:"radarr_url", StringValue:"", Datatype:"string"},
	{Label:"Radarr API Key", Placeholder:"api key", Description:"Radarr API Key, found in the Radarr General Settings", Property:"radarr_api_key", StringValue:"", Datatype:"string"},
	{Label:"Radarr Category", Placeholder:"Movies", Description:"Category in Radarr settings that is sent to download clients (Must be set!)", Property:"radarr_category", Datatype:"string"},
}

var PlexConfig = []schema.Configurator{
	{Label: "Plex Base URL", Placeholder: "https://plex.yourdomain.com:443", Description: "Base URL for connecting to Plex", Property: "plex_base_url", StringValue: "", Datatype: "string"},
	{Label: "Plex Auth Token", Placeholder: "auth token", Description:"Auth token for authenticating with Plex. See the Plex documentation for help retrieving it.", Property: "plex_auth_token", StringValue:"", Datatype: "string"},
	{Label: "Plex Namespace", Placeholder: "plex", Description:"Kubernetes Namespace that Plex runs in", Property:"plex_namespace", StringValue:"", Datatype:"string"},
	{Label: "Plex Deployment Name", Placeholder: "mediaserver", Description:"Name for your Plex deployment", Property:"plex_deployment_name", StringValue:"", Datatype:"string"},
}