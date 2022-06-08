package api

import (
	"net/http"
	"time"

	"github.com/lsnow99/conductorr/internal/conductorr/logger"
	"github.com/rs/zerolog"
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
		case zerolog.ErrorLevel:
			variant = "danger"
		case zerolog.WarnLevel:
			variant = "warning"
		case zerolog.InfoLevel:
			variant = ""
		}
		logsResp = append(logsResp, LogResponse{
			Timestamp: log.Timestamp,
			Variant:   variant,
			Msg:       log.Message,
		})
	}

	Respond(w, r, nil, logsResp, true)
}

func GetLogFile(w http.ResponseWriter, r *http.Request) {
	filename := "conductorr.log"
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	http.ServeFile(w, r, "conductorr.log")
}
