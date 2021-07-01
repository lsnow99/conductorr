package app

import (
	"github.com/lsnow99/conductorr/scheduler"
)

var DM DownloaderManager
var IM IndexerManager
var Monitor SystemMonitor

func init() {
	scheduler.RegisterTask(Monitor)
	scheduler.RegisterTask(&DM)
	scheduler.RegisterTask(&IM)
}
