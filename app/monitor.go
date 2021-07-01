package app

import (
	"reflect"
	"time"
)

type SystemStatus struct {
	msg        string
	systemName string
}

type SystemTypeStatus struct {
	healthy  bool
	msg      string
	statuses []SystemStatus
}

type SystemMonitor struct {
	// statusMap maps system types to an array of statuses that describe the health of the system
	statusMap map[string]SystemTypeStatus
}

func (m SystemMonitor) GetFrequency() time.Duration {
	return time.Minute * 20
}

type Service interface {
	TestConnection() error
	GetName() string
}

func (m SystemMonitor) DoTask() {
	m.statusMap["indexer"] = checkSystem("indexer", buildServices(IM.getIndexers()))
	m.statusMap["downloader"] = checkSystem("downloader", buildServices(DM.getDownloaders()))
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
			sts.statuses = append(sts.statuses, SystemStatus{
				msg:        "Could not connect to downloader " + service.GetName() + ": " + err.Error(),
				systemName: service.GetName(),
			})
		}
	}

	if len(services) == 0 {
		sts.healthy = false
		sts.msg = "No " + systemType + "s configured"
	} else if len(sts.statuses) == len(services) {
		sts.healthy = false
		sts.msg = "All " + systemType + "s are failing"
	} else if len(sts.statuses) > 0 {
		sts.healthy = false
		sts.msg = "Some downloaders are failing"
	} else {
		sts.healthy = true
		sts.msg = "All downloaders are connectable"
	}
	return sts
}
