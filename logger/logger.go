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

type getLogsReq struct {
	logsChan chan []LogMessage
}

type logsContainer struct {
	sendChan    chan LogMessage
	recvChan    chan getLogsReq
	logMessages []LogMessage
}

var logs logsContainer

func init() {
	logs.sendChan = make(chan LogMessage)
	logs.recvChan = make(chan getLogsReq)
	go func ()  {
		for {
			select {
			case log := <-logs.sendChan:
				logs.logMessages = append(logs.logMessages, log)
			case req := <-logs.recvChan:
				req.logsChan <- logs.logMessages
			}
		}
	}()
}

func GetLogs() []LogMessage {
	req := getLogsReq{
		logsChan: make(chan []LogMessage),
	}
	logs.recvChan <- req
	return <-req.logsChan
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
