package scheduler

import (
	"context"
	"time"
)

type Task interface {
	GetFrequency() time.Duration
	DoTask()
}

var Tasks []*Task
var ctx context.Context = context.Background()

func RegisterTask(t Task) {
	Tasks = append(Tasks, &t)
	go func() {
		ticker := time.NewTicker(t.GetFrequency())
		for {
			select {
			case <- ticker.C:
				t.DoTask()
			case <- ctx.Done():
				return
			}
		}
	}()
}