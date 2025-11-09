package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/lsnow99/conductorr/internal/conductorr/dbstore"
	"github.com/lsnow99/conductorr/internal/conductorr/services/search"
)

type RadarrSystemResource struct {
	AppName                       string `json:"appName"`
	InstanceName                  string `json:"instanceName"`
	Version                       string `json:"version"`
	BuildTime                     string `json:"buildTime"`
	IsDebug                       bool   `json:"isDebug"`
	IsProduction                  bool   `json:"isProduction"`
	IsAdmin                       bool   `json:"isAdmin"`
	IsUserInteractive             bool   `json:"isUserInteractive"`
	StartupPath                   string `json:"startupPath"`
	AppData                       string `json:"appData"`
	OsName                        string `json:"osName"`
	OsVersion                     string `json:"osVersion"`
	IsNetCore                     bool   `json:"isNetCore"`
	IsLinux                       bool   `json:"isLinux"`
	IsOsx                         bool   `json:"isOsx"`
	IsWindows                     bool   `json:"isWindows"`
	IsDocker                      bool   `json:"isDocker"`
	Mode                          string `json:"mode"`
	Branch                        string `json:"branch"`
	DatabaseType                  string `json:"databaseType"`
	DatabaseVersion               string `json:"databaseVersion"`
	Authentication                string `json:"authentication"`
	MigrationVersion              int    `json:"migrationVersion"`
	UrlBase                       string `json:"urlBase"`
	RuntimeVersion                string `json:"runtimeVersion"`
	RuntimeName                   string `json:"runtimeName"`
	StartTime                     string `json:"startTime"`
	PackageVersion                string `json:"packageVersion"`
	PackageAuthor                 string `json:"packageAuthor"`
	PackageUpdateMechanism        string `json:"packageUpdateMechanism"`
	PackageUpdateMechanismMessage string `json:"packageUpdateMechanismMessage"`
}

// For JellySeerr compatibility, this is all we need
// https://github.com/seerr-team/seerr/blob/develop/server/api/servarr/base.ts#L46
// TODO: Add full compatibility
type RadarrQualityProfile struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type RadarrRootFolder struct {
	Id         int    `json:"id"`
	Path       string `json:"path"`
	FreeSpace  int    `json:"freeSpace"`
	TotalSpace int    `json:"totalSpace"`
	Accessible bool   `json:"accessible"`
}

type RadarrTag struct {
	Id    int    `json:"id"`
	Label string `json:"label"`
}

type RadarrMovie struct {
	Id               int    `json:"id"`
	Title            string `json:"title"`
	IsAvailable      bool   `json:"isAvailable"`
	Monitored        bool   `json:"monitored"`
	TmdbId           int    `json:"tmdbId"`
	ImdbId           string `json:"imdbId"`
	TitleSlug        string `json:"titleSlug"`
	FolderName       string `json:"folderName"`
	Path             string `json:"path"`
	ProfileId        int    `json:"profileId"`
	QualityProfileId int    `json:"qualityProfileId"`
	Added            string `json:"added"`
	HasFile          bool   `json:"hasFile"`
	Tags             []int  `json:"tags"`
	MovieFile        struct {
		Id           int    `json:"id"`
		MovieId      int    `json:"movieId"`
		RelativePath string `json:"relativePath,omitempty"`
		Path         string `json:"path,omitempty"`
		Size         int    `json:"size"`
		DateAdded    string `json:"dateAdded"`
		SceneName    string `json:"sceneName,omitempty"`
		ReleaseGroup string `json:"releaseGroup,omitempty"`
		Edition      string `json:"edition,omitempty"`
		IndexerFlags int    `json:"indexerFlags,omitempty"`
		MediaInfo    struct {
			Id                    int    `json:"id"`
			AudioBitrate          int    `json:"audioBitrate"`
			AudioChannels         int    `json:"audioChannels"`
			AudioCodec            string `json:"audioCodec,omitempty"`
			AudioLanguages        string `json:"audioLanguages,omitempty"`
			AudioStreamCount      int    `json:"audioStreamCount"`
			VideoBitDepth         int    `json:"videoBitDepth"`
			VideoBitrate          int    `json:"videoBitrate"`
			VideoCodec            string `json:"videoCodec,omitempty"`
			VideoFps              int    `json:"videoFps"`
			VideoDynamicRange     string `json:"videoDynamicRange,omitempty"`
			VideoDynamicRangeType string `json:"videoDynamicRangeType,omitempty"`
			Resolution            string `json:"resolution,omitempty"`
			RunTime               string `json:"runTime,omitempty"`
			ScanType              string `json:"scanType,omitempty"`
			Subtitles             string `json:"subtitles,omitempty"`
		} `json:"mediaInfo,omitempty"`
		OriginalFilePath    string `json:"originalFilePath,omitempty"`
		QualityCutoffNotMet bool   `json:"qualityCutoffNotMet"`
	} `json:"movieFile,omitempty"`
}

func GetRadarrSystemStatus(w http.ResponseWriter, r *http.Request) {
	resp := RadarrSystemResource{
		AppName:                "Conductorr",
		InstanceName:           "Conductorr",
		Version:                "1.0.0",
		BuildTime:              "2025-10-06T20:48:04Z",
		IsDebug:                false,
		IsProduction:           true,
		IsAdmin:                false,
		IsUserInteractive:      true,
		StartupPath:            "/app/radarr/bin",
		AppData:                "/config",
		OsName:                 "alpine",
		OsVersion:              "3.22.2",
		IsNetCore:              true,
		IsLinux:                true,
		IsOsx:                  false,
		IsWindows:              false,
		IsDocker:               true,
		Mode:                   "console",
		Branch:                 "master",
		DatabaseType:           "sqLite",
		DatabaseVersion:        "3.49.2",
		Authentication:         "basic",
		MigrationVersion:       242,
		UrlBase:                "",
		RuntimeVersion:         "6.0.35",
		RuntimeName:            "netcore",
		StartTime:              "2025-11-02T00:07:19Z",
		PackageVersion:         "5.28.0.10274-ls286",
		PackageAuthor:          "[linuxserver.io](https://linuxserver.io)",
		PackageUpdateMechanism: "docker",
	}
	RespondRaw(w, r, nil, resp)
}

func GetRadarrQualityProfiles(w http.ResponseWriter, r *http.Request) {
	dbProfiles, err := dbstore.GetProfiles()
	if err != nil {
		RespondRaw(w, r, err, nil)
		return
	}

	resp := make([]RadarrQualityProfile, 0, len(dbProfiles))
	for _, profile := range dbProfiles {
		resp = append(
			resp,
			RadarrQualityProfile{
				Id:   profile.ID,
				Name: profile.Name.String,
			},
		)
	}
	RespondRaw(w, r, nil, resp)
}

func GetRadarrRootFolder(w http.ResponseWriter, r *http.Request) {
	dbPaths, err := dbstore.GetPaths()
	if err != nil && err != sql.ErrNoRows {
		RespondRaw(w, r, err, nil)
		return
	}

	resp := make([]RadarrRootFolder, len(dbPaths))
	for i, path := range dbPaths {
		resp[i] = RadarrRootFolder{
			Id:   path.ID,
			Path: path.Path,
			// TODO: just get this
			FreeSpace:  1_000_000,
			TotalSpace: 1_000_000,
			Accessible: true,
		}
	}
	RespondRaw(w, r, nil, resp)
}

func GetRadarrTags(w http.ResponseWriter, r *http.Request) {
	resp := [...]RadarrTag{}
	RespondRaw(w, r, nil, resp)
}

func GetRadarrMovieLookup(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	term := query.Get("term")

	results, err := search.SearchFuzzy(term, "movie", 1)
	if err != nil {
		RespondRaw(w, r, err, nil)
		return
	}
	resp := make([]RadarrMovie, 0, len(results.Results))
	for _, result := range results.Results {
		// ugh, this is jank.
		split := strings.Split(result.ID, ":")
		if len(split) != 2 {
			continue
		}
		id, err := strconv.Atoi(split[1])
		if err != nil {
			continue
		}

		resp = append(
			resp,
			RadarrMovie{
				Id:    id,
				Title: result.Title,
				IsAvailable: true,
				// Monitored:   false,
				TmdbId:    id,
				ImdbId:    "na",
				Tags:      []int{},
				TitleSlug: fmt.Sprintf("%d", id),
			},
		)
	}
	RespondRaw(w, r, nil, resp)
}
