package api

import (
	"net/http"

	"github.com/lsnow99/conductorr/internal/conductorr/app"
	"github.com/lsnow99/conductorr/internal/conductorr/scheduler"
)

func GetStatus(w http.ResponseWriter, r *http.Request) {
	Respond(w, r.Header.Get("hostname"), nil, app.Monitor.GetStatus(), true)
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks := make([]map[string]interface{}, 0)
	taskStatuses := scheduler.GetTaskStatuses()
	for _, task := range taskStatuses {
		tasks = append(tasks, map[string]interface{}{
			"next_run_time": task.NextRunTime,
			"id": task.ID,
			"name": task.GetTaskName(),
		})
	}
	Respond(w, r.Header.Get("hostname"), nil, tasks, true)
}
