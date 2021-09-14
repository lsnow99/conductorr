package app

import (
	"github.com/lsnow99/conductorr/integration"
	"github.com/lsnow99/conductorr/scheduler"
)

var DM DownloaderManager
var IM IndexerManager
var MSM MediaServerManager
var LM LibraryManager
var Monitor SystemMonitor

func init() {
	LM = NewLibraryManager()
	Monitor.statusMap = make(map[string]SystemTypeStatus)
	DM.backupReleaseMap = make(map[int][]integration.Release)
	scheduler.RegisterTask(&LibraryManager{})
	scheduler.RegisterTask(&Monitor)
	scheduler.RegisterTask(&DM)
	scheduler.RegisterTask(&IM)
}
