package dbstore

import (
	"database/sql"
	"encoding/json"
)

func NewMediaServer(mediaServerType, name string, config map[string]interface{}) (id int, err error) {
	configStr, err := json.Marshal(config)
	if err != nil {
		return 0, err
	}
	row := db.QueryRow(`
	INSERT INTO media_server (media_server_type, name, config) 
	VALUES (?, ?, ?) 
	RETURNING id`, mediaServerType, name, configStr)

	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func GetMediaServers() (mediaServers []MediaServer, err error) {
	rows, err := db.Query(`SELECT id, media_server_type, name, config FROM media_server WHERE true`)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	} else if err == sql.ErrNoRows {
		return mediaServers, nil
	}
	defer rows.Close()

	for rows.Next() {
		mediaServer := MediaServer{}
		var configStr string
		err := rows.Scan(&mediaServer.ID, &mediaServer.MediaServerType, &mediaServer.Name, &configStr)
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal([]byte(configStr), &mediaServer.Config); err != nil {
			return nil, err
		}
		mediaServers = append(mediaServers, mediaServer)
	}

	return mediaServers, nil
}

func UpdateMediaServer(id int, name string, config map[string]interface{}) error {
	configStr, err := json.Marshal(config)
	if err != nil {
		return err
	}
	_, err = db.Exec(`UPDATE media_server SET config = ?, name = ? WHERE id = ?`, configStr, name, id)
	return err
}

func DeleteMediaServer(id int) error {
	_, err := db.Exec(`DELETE FROM media_server WHERE id = ?`, id)
	return err
}

func GetMediaServer(id int) (mediaServer MediaServer, err error) {
	row := db.QueryRow(`SELECT id, media_server_type, name, config FROM media_server WHERE id = ?`, id)
	var configStr string
	if err = row.Scan(&mediaServer.ID, &mediaServer.MediaServerType, &mediaServer.Name, &configStr); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(configStr), &mediaServer.Config); err != nil {
		return
	}
	return
}