package backup

import (
	"sync"
	"time"
)

type bkup struct {
	filename  string
	timestamp time.Time
}

type BackupMapper struct {
	backups map[int]bkup
	sync.RWMutex
}

var bm BackupMapper

func init() {
	bm = BackupMapper{}
	bm.backups = make(map[int]bkup)
}

func GetBackupFile(id int) (string, time.Time, bool) {
	bm.RLock()
	defer bm.RUnlock()
	bkup, ok := bm.backups[id]
	return bkup.filename, bkup.timestamp, ok
}

func addBackup(filename string) int {
	bm.Lock()
	defer bm.Unlock()
	id := len(bm.backups) + 1
	bm.backups[id] = bkup{
		filename: filename,
		timestamp: time.Now(),
	}
	return id
}
