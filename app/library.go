package app

import (
	"time"

	"github.com/lsnow99/conductorr/dbstore"
	"github.com/lsnow99/conductorr/integration"
	"github.com/lsnow99/conductorr/logger"
)

type updatePathRequest struct {
	mediaID int
	pathOK  bool
}

type getPathRequest struct {
	mediaID  int
	respChan chan bool
}

type getAllPathsRequest struct {
	respChan chan map[int]bool
}

type LibraryManager struct {
	// map media ids to a boolean of whether the media is downloaded or not
	pathStatuses map[int]bool
	// channels to handle access to pathStatuses
	updateChan chan updatePathRequest
	getChan    chan getPathRequest
	getAllChan chan getAllPathsRequest
}

func NewLibraryManager() LibraryManager {
	lm := LibraryManager{}
	lm.pathStatuses = make(map[int]bool)
	lm.initChannels()
	return lm
}

func (lm *LibraryManager) initChannels() {
	lm.updateChan = make(chan updatePathRequest)
	lm.getChan = make(chan getPathRequest)
	lm.getAllChan = make(chan getAllPathsRequest)
	go func() {
		for {
			select {
			case req := <-lm.updateChan:
				lm.pathStatuses[req.mediaID] = req.pathOK
			case req := <-lm.getChan:
				id := req.mediaID
				pathGood, inMap := lm.pathStatuses[id]
				req.respChan <- pathGood && inMap
			case req := <-lm.getAllChan:
				newMap := make(map[int]bool)
				for k, v := range lm.pathStatuses {
					newMap[k] = v
				}
				req.respChan <- newMap
			}
		}
	}()
}

func (lm *LibraryManager) GetFrequency() time.Duration {
	return time.Minute * 30
}

func (lm *LibraryManager) DoTask() {
	lm.checkPaths()
}

func (lm *LibraryManager) checkPaths() {
	medias, err := dbstore.GetAllMedia()
	if err != nil {
		logger.LogDanger(err)
		return
	}
	for _, media := range medias {
		if media.Path.Valid {
			lm.CheckMediaPath(media.ID, media.Path.String)
		}
	}
}

func (lm *LibraryManager) CheckMediaPath(mediaID int, path string) bool {
	status := integration.CheckPath(path) == nil
	lm.SetMediaPathStatus(mediaID, status)
	return status
}

func (lm *LibraryManager) SetMediaPathStatus(id int, status bool) {
	req := updatePathRequest{
		mediaID: id,
		pathOK:  status,
	}
	lm.updateChan <- req
}

func (lm *LibraryManager) GetMediaPathStatus(id int) bool {
	req := getPathRequest{
		mediaID:  id,
		respChan: make(chan bool),
	}
	lm.getChan <- req
	return <-req.respChan
}

func (lm *LibraryManager) GetAllMediaPathStatuses() map[int]bool {
	req := getAllPathsRequest{
		respChan: make(chan map[int]bool),
	}
	lm.getAllChan <- req
	return <-req.respChan
}
