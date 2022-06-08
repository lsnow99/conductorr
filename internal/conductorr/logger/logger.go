package logger

import (
	"os"
	"time"

	"github.com/lsnow99/conductorr/internal/conductorr/helpers"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
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

type LogHook struct{}

var logs logsContainer
const cacheSize = 500

func Init() {
	zerolog.TimestampFieldName = "timestamp"
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	log.Logger = zerolog.New(
		zerolog.MultiLevelWriter(zerolog.ConsoleWriter{Out: os.Stdout}, &lumberjack.Logger{
			Filename:   "./conductorr.log",
			MaxSize:    30, // megabytes
			MaxBackups: 3,
			MaxAge:     28,    //days
			Compress:   false,
		})).
		With().
		Caller().
		Logger().
		Hook(LogHook{})

	logs.sendChan = make(chan LogMessage, cacheSize * 2)
	logs.recvChan = make(chan getLogsReq)
	go func() {
		for {
			select {
			case log := <-logs.sendChan:
				// Append a log while ensuring there can never be more than cacheSize logs
				logs.logMessages = append(
					logs.logMessages[
						helpers.Min(0, helpers.Max(cacheSize-1-len(logs.logMessages), 1)):
						helpers.Min(len(logs.logMessages), cacheSize-1)],
				log)
			case req := <-logs.recvChan:
				req.logsChan <- logs.logMessages
			}
		}
	}()
}

func (h LogHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	timestamp := time.Now()
	e.Time("timestamp", timestamp)
	
	if level != zerolog.NoLevel && len(msg) > 0 {
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
