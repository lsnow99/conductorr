package logger

import (
	"log"
	"sync"
	"time"
)

const (
	Info = iota
	Warn
	Danger
)

type LogMessage struct {
	Level uint8
	Message string
	Timestamp time.Time
}

type logsContainer struct {
	sync.Mutex
	logMessages []LogMessage
}

var logs logsContainer

func GetLogs() []LogMessage {
	logs.Lock()
	defer logs.Unlock()
	return logs.logMessages
}

func LogToStdout(err error) {
	log.Println(err.Error())
}

func LogInfo(err error) {
	addLog(err, Info)
}

func LogWarn(err error) {
	addLog(err, Warn)
}

func LogDanger(err error) {
	addLog(err, Danger)
}

func addLog(err error, level uint8) {
	logs.Lock()
	defer logs.Unlock()
	logs.logMessages = append(logs.logMessages, LogMessage{
		Level: level,
		Message: err.Error(),
		Timestamp: time.Now(),
	})
}