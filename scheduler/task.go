package scheduler

import (
	"context"
	"time"
)

type Task interface {
	GetFrequency() time.Duration
	DoTask() error
}

var Tasks []*Task
var ctx context.Context

func RegisterTask(t Task) {
	Tasks = append(Tasks, &t)
	go func() {
		ticker := time.NewTicker(t.GetFrequency())
		for {
			select {
			case <- ticker.C:
				err := t.DoTask()
				if err != nil {

				}
			case <- ctx.Done():
				return
			}
		}
	}()
}