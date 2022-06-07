package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	Info = iota
	Warn
	Danger
)

type LogMessage struct {
	Level     zerolog.Level
	Message   string
	Timestamp time.Time
}

type getLogsReq struct {
	logsChan chan []LogMessage
}

type logsContainer struct {
	sendChan    chan LogMessage
	recvChan    chan getLogsReq
	logMessages []LogMessage
}

var logs logsContainer

type LogHook struct{}

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Caller().Logger()

	logs.sendChan = make(chan LogMessage)
	logs.recvChan = make(chan getLogsReq)
	go func() {
		for {
			select {
			case log := <-logs.sendChan:
				// Append a log while chopping off one at the beginning if there are at least 500 logs
				logs.logMessages = append([]LogMessage{log}, logs.logMessages[min(len(logs.logMessages), 1):min(len(logs.logMessages), 500)]...)
			case req := <-logs.recvChan:
				req.logsChan <- logs.logMessages
			}
		}
	}()
}

func (h LogHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	if level != zerolog.NoLevel {
		timestamp := time.Now()
		e.Time("timestamp", timestamp)
		logs.sendChan <- LogMessage{
			Level:     level,
			Message:   msg,
			Timestamp: timestamp,
		}
	}
}

func GetLogs() []LogMessage {
	req := getLogsReq{
		logsChan: make(chan []LogMessage),
	}
	logs.recvChan <- req
	return <-req.logsChan
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
