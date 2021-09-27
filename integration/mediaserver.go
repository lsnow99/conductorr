package integration

import "fmt"

type MediaServer interface {
	/*
		ImportMedia imports media from a filepath into the media server
	*/
	ImportMedia(string) error
	/*
		TestConnection returns a non-nil error if the media server fails to connect
	*/
	TestConnection() error
	/*
		GetMediaPaths returns a slice of all media paths available on the server
	*/
	GetMediaPaths() ([]MediaPath, error)
}

type MediaPath struct {
	Title string
	ParentTitle string
	GrandparentTitle string
	Path string
}

func NewMediaServerFromConfig(mediaServerType string, config map[string]interface{}) (MediaServer, error) {
	switch mediaServerType {
	case "plex":
		return NewPlexFromConfig(config)
	case "jellyfin":
		return NewJellyfinFromConfig(config)
	}
	return nil, fmt.Errorf("unrecognized media server type: %s", mediaServerType)
}
