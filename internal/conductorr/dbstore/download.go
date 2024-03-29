package dbstore

import (
	"database/sql"

	"github.com/lsnow99/conductorr/pkg/constant"
)

func NewDownload(mediaID int, downloaderID int, identifier, status, friendlyName, releaseID string) (id int, err error) {
	row := db.QueryRow(`
		INSERT INTO download 
		(media_id, downloader_id, identifier, status, friendly_name, release_id) 
		VALUES (?, ?, ?, ?, ?, ?) 
		RETURNING id
	`, mediaID, downloaderID, identifier, status, friendlyName, releaseID)

	err = row.Scan(&id)
	return id, err
}

func UpdateDownloadStatusByIdentifier(identifier, status string) error {
	_, err := db.Exec(`
		UPDATE download
		SET status = ?
		WHERE identifier = ?
	`, status, identifier)
	return err
}

func GetFinishedDownloads() ([]Download, error) {
	rows, err := db.Query(`
		SELECT id, media_id, downloader_id, status, friendly_name, identifier, release_id
		FROM download
		WHERE status in ('done', 'cerror', 'error')
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	downloads := make([]Download, 0)
	for rows.Next() {
		download := Download{}
		if err := rows.Scan(&download.ID, &download.MediaID, &download.DownloaderID, &download.Status, &download.FriendlyName, &download.Identifier, &download.ReleaseID); err != nil {
			return nil, err
		}
		downloads = append(downloads, download)
	}
	return downloads, nil
}

func GetDownloads() ([]Download, error) {
	rows, err := db.Query(`
		SELECT id, media_id, downloader_id, status, friendly_name, identifier, release_id
		FROM download
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	downloads := make([]Download, 0)
	for rows.Next() {
		download := Download{}
		if err := rows.Scan(&download.ID, &download.MediaID, &download.DownloaderID, &download.Status, &download.FriendlyName, &download.Identifier, &download.ReleaseID); err != nil {
			return nil, err
		}
		downloads = append(downloads, download)
	}
	return downloads, nil
}

func GetActiveDownloads() ([]Download, error) {
	rows, err := db.Query(`
		SELECT id, media_id, downloader_id, status, friendly_name, identifier, release_id
		FROM download
		WHERE status IN (?, ?, ?, ?, ?)
	`, constant.StatusWaiting, constant.StatusPaused, constant.StatusDownloading, constant.StatusComplete, constant.StatusProcessing)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	downloads := make([]Download, 0)
	for rows.Next() {
		download := Download{}
		if err := rows.Scan(&download.ID, &download.MediaID, &download.DownloaderID, &download.Status, &download.FriendlyName, &download.Identifier, &download.ReleaseID); err != nil {
			return nil, err
		}
		downloads = append(downloads, download)
	}
	return downloads, nil
}

func GetDownloadByIdentifier(identifier string) (Download, error) {
	row := db.QueryRow(`SELECT id, media_id, downloader_id, status, friendly_name, identifier, release_id FROM download WHERE identifier = ?`, identifier)
	download := Download{}
	return download, row.Scan(&download.ID, &download.MediaID, &download.DownloaderID, &download.Status, &download.FriendlyName, &download.Identifier, &download.ReleaseID)
}

func DumpDownloads(tx *sql.Tx) ([]*Download, error) {
	rows, err := db.Query(`
		SELECT id, media_id, downloader_id, status, friendly_name, identifier, release_id
		FROM download
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	downloads := make([]*Download, 0)
	for rows.Next() {
		download := &Download{}
		if err := rows.Scan(&download.ID, &download.MediaID, &download.DownloaderID, &download.Status, &download.FriendlyName, &download.Identifier, &download.ReleaseID); err != nil {
			return nil, err
		}
		downloads = append(downloads, download)
	}
	return downloads, rows.Err()
}

func RestoreDownloads(tx *sql.Tx, downloads []*Download) error {
	for _, download := range downloads {
		_, err := tx.Exec(`
		INSERT INTO download 
		(id, media_id, downloader_id, identifier, status, friendly_name, release_id) 
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`, download.ID, download.MediaID, download.DownloaderID, download.Identifier, download.Status, download.FriendlyName, download.ReleaseID)
		if err != nil {
			return err
		}
	}
	return nil
}
