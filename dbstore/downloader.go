package dbstore

import "encoding/json"

func NewDownloader(downloaderType, name string, config map[string]interface{}) error {
	configStr, err := json.Marshal(config)
	if err != nil {
		return err
	}
	_, err = db.Exec(`INSERT INTO downloader (downloader_type, name, config) VALUES (?, ?, ?)`, downloaderType, name, configStr)
	return err
}

func GetDownloaders() (downloaders []Downloader, err error) {
	rows, err := db.Query(`SELECT id, downloader_type, name, config FROM downloader WHERE true`)
	if err != nil {
		return nil, err
	}

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