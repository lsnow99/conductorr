package app

import (
	"errors"
	"sync"

	"github.com/lsnow99/conductorr/integration"
	"github.com/lsnow99/conductorr/logger"
)

type ManagedMediaServer struct {
	ID              int
	Name            string
	MediaServerType string
	integration.MediaServer
}

type MediaServerManager struct {
	sync.RWMutex
	mediaServers []ManagedMediaServer
}

func (msm *MediaServerManager) RegisterMediaServer(id int, mediaServerType, name string, configuration map[string]interface{}) error {
	msvr, err := integration.NewMediaServerFromConfig(mediaServerType, configuration)
	if err != nil {
		return err
	}
	mmsvr := ManagedMediaServer{
		ID:              id,
		Name:            name,
		MediaServerType: mediaServerType,
		MediaServer:     msvr,
	}
	var added bool
	for i, mediaServer := range msm.mediaServers {
		if mediaServer.ID == id {
			msm.mediaServers[i] = mmsvr
			added = true
		}
	}
	if !added {
		msm.mediaServers = append(msm.mediaServers, mmsvr)
	}
	return nil
}

func (msm *MediaServerManager) ImportMedia(path string) error {
	var hadError bool
	for _, mediaServer := range msm.mediaServers {
		err := mediaServer.ImportMedia(path)
		if err != nil {
			hadError = true
			logger.LogDanger(err)
		}
	}

	if hadError {
		return errors.New("could not import media path in one or more media servers. See log for details")
	}

	return nil
}
