package scheduler

import (
	"context"
	"time"
)

type Task interface {
	GetFrequency() time.Duration
	DoTask()
}

type managedTask struct {
	Task
	startChan chan struct{}
}

var tasks []*managedTask
var ctx context.Context = context.Background()

func RegisterTask(t Task) {
	mt := &managedTask{
		Task: t,
		startChan: make(chan struct{}),
	}
	tasks = append(tasks, mt)
	go func() {
		<-mt.startChan
		t.DoTask()
		ticker := time.NewTicker(t.GetFrequency())
		for {
			select {
			case <- ticker.C:
				go func ()  {
					t.DoTask()
				}()
			case <- ctx.Done():
				return
			}
		}
	}()
}

func StartTasks() {
	for _, task := range tasks {
		task.startChan <- struct{}{}
	}
}