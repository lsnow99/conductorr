package app

import (
	"errors"
	"sync"
	"time"

	"github.com/lsnow99/conductorr/dbstore"
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

func (*MediaServerManager) GetTaskName() string {
	return "Media Server Manager"
}

func (msm *MediaServerManager) GetFrequency() time.Duration {
	return time.Minute * 10
}

func (msm *MediaServerManager) DoTask() {
	msm.syncMediaPaths()
}

func (msm *MediaServerManager) syncMediaPaths() error {
	msm.RLock()
	mediaServers := msm.mediaServers
	msm.RUnlock()

	mediaPaths := make([]integration.MediaPath, 0)
	for _, mediaServer := range mediaServers {
		paths, err := mediaServer.GetMediaPaths()
		if err != nil {
			return err
		}
		mediaPaths = append(mediaPaths, paths...)
	}

	_, err := dbstore.GetAllMedia()
	if err != nil {
		return err
	}

	// for _, media := range medias {
	// 	sort.Slice(mediaPaths, func(i, j int) bool {
	// 		return mediaPaths[i] < mediaPaths[j]
	// 	})
	// }

	return nil
}

func (msm *MediaServerManager) RegisterMediaServer(id int, mediaServerType, name string, configuration map[string]interface{}) error {
	msm.Lock()
	defer msm.Unlock()

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

func (msm *MediaServerManager) getMediaServers() []ManagedMediaServer {
	msm.RLock()
	defer msm.RUnlock()
	return msm.mediaServers
}

func (msm *MediaServerManager) ImportMedia(path string) error {

	mediaServers := msm.getMediaServers()

	var hadError bool
	for _, mediaServer := range mediaServers {
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
