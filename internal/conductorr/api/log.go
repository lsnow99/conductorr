package api

import (
	"net/http"
	"time"

	"github.com/lsnow99/conductorr/internal/conductorr/logger"
)

type LogResponse struct {
	Timestamp time.Time `json:"timestamp,omitempty"`
	Variant   string    `json:"variant,omitempty"`
	Msg       string    `json:"msg,omitempty"`
}

func GetLogs(w http.ResponseWriter, r *http.Request) {
	logs := logger.GetLogs()

	logsResp := make([]LogResponse, 0)
	for _, log := range logs {
		var variant string
		switch log.Level {
		case logger.Danger:
			variant = "danger"
		case logger.Warn:
			variant = "warning"
		case logger.Info:
			variant = ""
		}
		logsResp = append(logsResp, LogResponse{
			Timestamp: log.Timestamp,
			Variant: variant,
			Msg: log.Message,
		})
	}

	Respond(w, r.Header.Get("hostname"), nil, logsResp, true)
}