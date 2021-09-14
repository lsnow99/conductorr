package dbstore

import (
	"database/sql"
	"encoding/json"
)

func NewDownloader(downloaderType, name, fileAction string, config map[string]interface{}) (id int, err error) {
	configStr, err := json.Marshal(config)
	if err != nil {
		return 0, err
	}
	row := db.QueryRow(`
	INSERT INTO downloader (downloader_type, name, file_action config) 
	VALUES (?, ?, ?, ?) 
	RETURNING id`, downloaderType, name, fileAction, configStr)

	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func GetDownloaders() (downloaders []Downloader, err error) {
	rows, err := db.Query(`SELECT id, downloader_type, name, file_action, config FROM downloader WHERE true`)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	} else if err == sql.ErrNoRows {
		return downloaders, nil
	}
	defer rows.Close()

	for rows.Next() {
		downloader := Downloader{}
		var configStr string
		err := rows.Scan(&downloader.ID, &downloader.DownloaderType, &downloader.Name, &downloader.FileAction, &configStr)
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

func UpdateDownloader(id int, name, fileAction string, config map[string]interface{}) error {
	configStr, err := json.Marshal(config)
	if err != nil {
		return err
	}
	_, err = db.Exec(`UPDATE downloader SET config = ?, name = ?, file_action = ? WHERE id = ?`, configStr, name, fileAction, id)
	return err
}

func DeleteDownloader(id int) error {
	_, err := db.Exec(`DELETE FROM downloader WHERE id = ?`, id)
	return err
}

func GetDownloader(id int) (downloader Downloader, err error) {
	row := db.QueryRow(`SELECT id, downloader_type, name, file_action, config FROM downloader WHERE id = ?`, id)
	var configStr string
	if err = row.Scan(&downloader.ID, &downloader.DownloaderType, &downloader.Name, &downloader.FileAction, &configStr); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(configStr), &downloader.Config); err != nil {
		return
	}
	return
}

func DumpDownloaders(tx *sql.Tx) ([]*Downloader, error) {
	downloaders := make([]*Downloader, 0)

	rows, err := db.Query(`SELECT id, downloader_type, name, file_action, config FROM downloader WHERE true`)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	} else if err == sql.ErrNoRows {
		return downloaders, nil
	}
	defer rows.Close()

	for rows.Next() {
		downloader := &Downloader{}
		var configStr string
		err := rows.Scan(&downloader.ID, &downloader.DownloaderType, &downloader.Name, &downloader.FileAction, &configStr)
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

func RestoreDownloaders(tx *sql.Tx, downloaders []*Downloader) error {
	for _, downloader := range downloaders {
		configStr, err := json.Marshal(downloader.Config)
		if err != nil {
			return err
		}
		_, err = tx.Exec(`
		INSERT INTO downloader (id, downloader_type, name, file_action, config) 
		VALUES (?, ?, ?, ?, ?)`, downloader.ID, downloader.DownloaderType, downloader.Name, downloader.FileAction, configStr)

		if err != nil {
			return err
		}
	}
	return nil
}
