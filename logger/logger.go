package logger

import (
	"log"
	"time"
)

const (
	Info = iota
	Warn
	Danger
)

type LogMessage struct {
	Level     uint8
	Message   string
	Timestamp time.Time
}

type logsContainer struct {
	sendChan    chan LogMessage
	recvChan    chan []LogMessage
	logMessages []LogMessage
}

var logs logsContainer

func init() {
	logs.sendChan = make(chan LogMessage)
	logs.recvChan = make(chan []LogMessage)
	go func ()  {
		for {
			select {
			case log := <-logs.sendChan:
				logs.logMessages = append(logs.logMessages, log)
			case <-logs.recvChan:
				logs.recvChan <- logs.logMessages
			}
		}
	}()
}

func GetLogs() []LogMessage {
	return <-logs.recvChan
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
	logs.sendChan <- LogMessage{
		Level:     level,
		Message:   err.Error(),
		Timestamp: time.Now(),
	}
}
