package dbstore

import "github.com/lsnow99/conductorr/constant"

func NewDownload(mediaID int, downloaderID int, identifier string, status string, friendlyName string) (id int, err error) {
	row := db.QueryRow(`
		INSERT INTO download 
		(media_id, downloader_id, identifier, status, friendly_name) 
		VALUES (?, ?, ?, ?, ?) 
		RETURNING id
	`, mediaID, downloaderID, identifier, status, friendlyName)

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

func GetDoneDownloads() ([]Download, error) {
	rows, err := db.Query(`
		SELECT id, media_id, downloader_id, status, friendly_name, identifier
		FROM download
		WHERE status = ?
	`, constant.StatusDone)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	downloads := make([]Download, 0)
	for rows.Next() {
		download := Download{}
		if err := rows.Scan(&download.ID, &download.MediaID, &download.DownloaderID, &download.Status, &download.FriendlyName, &download.Identifier); err != nil {
			return nil, err
		}
		downloads = append(downloads, download)
	}
	return downloads, nil
}

func GetDownloads() ([]Download, error) {
	rows, err := db.Query(`
		SELECT id, media_id, downloader_id, status, friendly_name, identifier
		FROM download
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	downloads := make([]Download, 0)
	for rows.Next() {
		download := Download{}
		if err := rows.Scan(&download.ID, &download.MediaID, &download.DownloaderID, &download.Status, &download.FriendlyName, &download.Identifier); err != nil {
			return nil, err
		}
		downloads = append(downloads, download)
	}
	return downloads, nil
}

func GetActiveDownloads() ([]Download, error) {
	rows, err := db.Query(`
		SELECT id, media_id, downloader_id, status, friendly_name, identifier
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
		if err := rows.Scan(&download.ID, &download.MediaID, &download.DownloaderID, &download.Status, &download.FriendlyName, &download.Identifier); err != nil {
			return nil, err
		}
		downloads = append(downloads, download)
	}
	return downloads, nil
}

func GetDownloadByIdentifier(identifier string) (Download, error) {
	row := db.QueryRow(`SELECT id, media_id, downloader_id, status, friendly_name, identifier FROM download WHERE identifier = ?`, identifier)
	download := Download{}
	return download, row.Scan(&download.ID, &download.MediaID, &download.DownloaderID, &download.Status, &download.FriendlyName, &download.Identifier)
}