package dbstore

import (
	"database/sql"
	"encoding/json"
)

func NewDownloader(downloaderType, name string, config map[string]interface{}) (id int, err error) {
	configStr, err := json.Marshal(config)
	if err != nil {
		return 0, err
	}
	row := db.QueryRow(`
	INSERT INTO downloader (downloader_type, name, config) 
	VALUES (?, ?, ?) 
	RETURNING id`, downloaderType, name, configStr)

	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func GetDownloaders() (downloaders []Downloader, err error) {
	rows, err := db.Query(`SELECT id, downloader_type, name, config FROM downloader WHERE true`)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	} else if err == sql.ErrNoRows {
		return downloaders, nil
	}
	defer rows.Close()

	for rows.Next() {
		downloader := Downloader{}
		var configStr string
		err := rows.Scan(&downloader.ID, &downloader.DownloaderType, &downloader.Name, &configStr)
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal([]byte(configStr), &downloader.Config); err != nil {
			return nil, err
		}
		downloaders = append(downloaders, downloader)
	}

	return downloaders, nil
}

func UpdateDownloader(id int, name string, config map[string]interface{}) error {
	configStr, err := json.Marshal(config)
	if err != nil {
		return err
	}
	_, err = db.Exec(`UPDATE downloader SET config = ?, name = ? WHERE id = ?`, configStr, name, id)
	return err
}

func DeleteDownloader(id int) error {
	_, err := db.Exec(`DELETE FROM downloader WHERE id = ?`, id)
	return err
}

func GetDownloader(id int) (downloader Downloader, err error) {
	row := db.QueryRow(`SELECT id, downloader_type, name, config FROM downloader WHERE id = ?`, id)
	var configStr string
	if err = row.Scan(&downloader.ID, &downloader.DownloaderType, &downloader.Name, &configStr); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(configStr), &downloader.Config); err != nil {
		return
	}
	return
}