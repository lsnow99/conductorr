package scheduler

import (
	"context"
	"time"
)

type Task interface {
	GetFrequency() time.Duration
	GetTaskName() string
	DoTask()
}

type ManagedTask struct {
	Task
	NextRunTime time.Time
	ID          int
	startChan   chan struct{}
}

type statusReq struct {
	respCh chan []*ManagedTask
}

type updateReq struct {
	nextRunTime time.Time
	taskID      int
}

var startCh = make(chan struct{})
var registerCh = make(chan Task)
var statusCh = make(chan statusReq)
var updateCh = make(chan updateReq)
var ctx context.Context = context.Background()

func init() {
	var tasks []*ManagedTask
	var counter = 1
	go func() {
		for {
			select {
			case <-startCh:
				for _, task := range tasks {
					task.startChan <- struct{}{}
				}
			case t := <-registerCh:
				mt := &ManagedTask{
					Task:      t,
					startChan: make(chan struct{}),
					ID:        counter,
				}
				counter++
				tasks = append(tasks, mt)
				go loopTask(mt)
			case sr := <-statusCh:
				sr.respCh <- tasks
			case ur := <-updateCh:
				for _, task := range tasks {
					if task.ID == ur.taskID {
						task.NextRunTime = ur.nextRunTime
						break
					}
				}
			}
		}
	}()
}

func loopTask(mt *ManagedTask) {
	<-mt.startChan
	mt.DoTask()

	mt.NextRunTime = time.Now().Add(mt.GetFrequency())

	timer := time.NewTimer(mt.GetFrequency())

	for {
		select {
		case <-timer.C:
			go mt.DoTask()
			ur := updateReq{
				nextRunTime: mt.NextRunTime.Add(mt.GetFrequency()),
				taskID:      mt.ID,
			}
			updateCh <- ur
			timer.Reset(mt.GetFrequency())
		case <-ctx.Done():
			return
		}
	}
}

func RegisterTask(t Task) {
	registerCh <- t
}

func StartTasks() {
	startCh <- struct{}{}
}

func GetTaskStatuses() []*ManagedTask {
	sr := statusReq{
		respCh: make(chan []*ManagedTask),
	}
	statusCh <- sr
	return <-sr.respCh
}
