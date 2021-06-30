package integration

import "sync"

// syncQueue synchronized queue
type syncQueue struct {
	sync.Mutex
	Releases []*Release
}

var queue syncQueue

func addToQueue(release *Release) {
	queue.Lock()
	defer queue.Unlock()
	queue.Releases = append(queue.Releases, release)
}

func DownloadMediaRelease(release *Release) {

}