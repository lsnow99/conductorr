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
}

func NewMediaServerFromConfig(mediaServerType string, config map[string]interface{}) (MediaServer, error) {
	switch mediaServerType {
	case "plex":
		return NewPlexFromConfig(config)
	}
	return nil, fmt.Errorf("unrecognized media server type: %s", mediaServerType)
}
