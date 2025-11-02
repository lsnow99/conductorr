package api

import (
	"net/http"
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

func GetRadarrSystemStatus(w http.ResponseWriter, r *http.Request) {
	resp := RadarrSystemResource{
		AppName:                "Radarr",
		InstanceName:           "Conductorr",
		Version:                "5.28.0.10274",
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
