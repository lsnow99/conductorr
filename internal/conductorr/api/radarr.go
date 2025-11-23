package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/lsnow99/conductorr/internal/conductorr/dbstore"
	"github.com/lsnow99/conductorr/internal/conductorr/services/search"
	"github.com/rs/zerolog/log"
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
	Id               int    `json:"id,omitempty"`
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

type RadarrCommand struct {
	Name        string `json:"name"`
	CommandName string `json:"commandName"`
	MovieIds    []int  `json:"movieIds"`
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

func GetRadarrMovies(w http.ResponseWriter, r *http.Request) {
	medias, err := dbstore.GetRecentlyAddedMedia(100)
	if err != nil {
		RespondRaw(w, r, err, nil)
		return
	}

	resp := make([]RadarrMovie, 0, len(medias))
	for _, result := range medias {
		if result.ContentType.Valid && result.ContentType.String == "movie" {
			movie := RadarrMovie{
				Id: result.ID,
			}
			if result.Title.Valid {
				movie.Title = result.Title.String
			}
			resp = append(resp, movie)
		}
	}
	RespondRaw(w, r, nil, resp)
}

func UpsertRadarrMovie(w http.ResponseWriter, r *http.Request) {
	movie := RadarrMovie{}
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		RespondRaw(w, r, err, nil)
		return
	}
	RespondRaw(w, r, fmt.Errorf("not implemented"), nil)
}

func GetRadarrMovieLookup(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	term := query.Get("term")

	results, err := search.SearchFuzzy(term, "movie", 1)
	if err != nil {
		RespondRaw(w, r, err, nil)
		return
	}

	// Get media that already exists
	tmdbIDs := make([]int, 0, len(results.Results))
	for _, result := range results.Results {
		split := strings.Split(result.ID, ":")
		if len(split) != 2 {
			continue
		}
		id, err := strconv.Atoi(split[1])
		if err != nil {
			continue
		}
		tmdbIDs = append(tmdbIDs, id)
	}
	existingMedias, err := dbstore.GetMediaByTmdbIDs(tmdbIDs)
	if err != nil {
		RespondRaw(w, r, err, nil)
		return
	}
	tmdbIDToMedia := make(map[int]*dbstore.Media)
	for _, media := range existingMedias {
		if media.TmdbID.Valid {
			tmdbIDToMedia[int(media.TmdbID.Int32)] = media
		}
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

		movie := RadarrMovie{
			Title:       result.Title,
			IsAvailable: true,
			TmdbId:      id,
			ImdbId:      "na",
			Tags:        []int{},
			TitleSlug:   fmt.Sprintf("%d", id),
		}
		if media, ok := tmdbIDToMedia[id]; ok {
			movie.Id = media.ID
			movie.Monitored = true
		}

		resp = append(
			resp,
			movie,
		)
	}
	RespondRaw(w, r, nil, resp)
}

func RunRadarrCommand(w http.ResponseWriter, r *http.Request) {
	command := RadarrCommand{}
	err := json.NewDecoder(r.Body).Decode(&command)
	if err != nil {
		RespondRaw(w, r, err, nil)
		return
	}
	switch command.CommandName {
	case "MoviesSearch":
		RespondRaw(w, r, fmt.Errorf("not implemented"), nil)
		return
	default:
		log.Warn().Msg(fmt.Sprintf("unknown radarr command %s", command.CommandName))
	}
	RespondRaw(w, r, nil, command)
}

func RedirectRadarrMovie(w http.ResponseWriter, r *http.Request) {
	RespondRaw(w, r, fmt.Errorf("Not implemented"), nil)
}
