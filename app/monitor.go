package app

import (
	"reflect"
	"sync"
	"time"

	"github.com/lsnow99/conductorr/dbstore"
	"github.com/lsnow99/conductorr/integration"
)

type SystemStatus struct {
	Msg        string `json:"msg,omitempty"`
	SystemName string `json:"system_name,omitempty"`
}

type SystemTypeStatus struct {
	Healthy  bool           `json:"healthy"`
	Msg      string         `json:"msg,omitempty"`
	Statuses []SystemStatus `json:"statuses,omitempty"`
}

type SystemMonitor struct {
	sync.RWMutex
	// statusMap maps system types to an array of statuses that describe the health of the system
	statusMap map[string]SystemTypeStatus
}

func (m *SystemMonitor) GetFrequency() time.Duration {
	return time.Minute * 20
}

type Service interface {
	TestConnection() error
	GetName() string
}

func (m *SystemMonitor) DoTask() {
	m.Lock()
	defer m.Unlock()
	m.statusMap["indexer"] = checkSystem("indexer", buildServices(IM.getIndexers()))
	m.statusMap["downloader"] = checkSystem("downloader", buildServices(DM.getDownloaders(false)))
	m.statusMap["path"] = getPathsStatus()
}

func (m *SystemMonitor) GetStatus() map[string]SystemTypeStatus {
	m.RLock()
	defer m.RUnlock()
	return m.statusMap
}

func buildServices(arr interface{}) (services []Service) {
	s := reflect.ValueOf(arr)
	for i := 0; i < s.Len(); i++ {
		elem := s.Index(i)
		if service, ok := elem.Interface().(Service); ok {
			services = append(services, service)
		}
	}
	return services
}

func checkSystem(systemType string, services []Service) (sts SystemTypeStatus) {
	for _, service := range services {
		if err := service.TestConnection(); err != nil {
			sts.Statuses = append(sts.Statuses, SystemStatus{
				Msg:        "Could not connect to " + systemType + " " + service.GetName() + ": " + err.Error(),
				SystemName: service.GetName(),
			})
		}
	}

	if len(services) == 0 {
		sts.Healthy = false
		sts.Msg = "No " + systemType + "s configured"
	} else if len(sts.Statuses) == len(services) {
		sts.Healthy = false
		sts.Msg = "All " + systemType + "s are failing"
	} else if len(sts.Statuses) > 0 {
		sts.Healthy = false
		sts.Msg = "Some " + systemType + "s are failing"
	} else {
		sts.Healthy = true
		sts.Msg = "All " + systemType + "s are connectable"
	}
	return sts
}

func getPathsStatus() (sts SystemTypeStatus) {
	paths, err := dbstore.GetPaths()
	if err != nil {
		sts.Healthy = false
		sts.Msg = "Could not retrieve paths from database"
		return
	}
	for _, path := range paths {
		if err := integration.CheckPath(path.Path); err != nil {
			sts.Statuses = append(sts.Statuses, SystemStatus{
				Msg: err.Error(),
				SystemName: path.Path,
			})
		}
	}

	if len(paths) == 0 {
		sts.Healthy = false
		sts.Msg = "No paths configured"
	} else if len(sts.Statuses) == len(paths) {
		sts.Healthy = false
		sts.Msg = "All paths are failing"
	} else if len(sts.Statuses) > 0 {
		sts.Healthy = false
		sts.Msg = "Some paths are failing"
	} else {
		sts.Healthy = true
		sts.Msg = "All paths are OK"
	}
	return
}