package app

import (
	"github.com/lsnow99/conductorr/integration"
	"github.com/lsnow99/conductorr/scheduler"
)

var DM DownloaderManager
var IM IndexerManager
var MSM MediaServerManager
var Monitor SystemMonitor

func init() {
	Monitor.statusMap = make(map[string]SystemTypeStatus)
	DM.backupReleaseMap = make(map[int][]integration.Release)
	scheduler.RegisterTask(&Monitor)
	scheduler.RegisterTask(&DM)
	scheduler.RegisterTask(&IM)
}
