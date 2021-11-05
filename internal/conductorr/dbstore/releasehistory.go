package dbstore

import "time"

func GetReleaseHistoryForMedia(mediaID int) ([]ReleaseHistory, error) {
	history := make([]ReleaseHistory, 0)

	rows, err := db.Query(`SELECT id, media_id, last_attempt FROM release_history WHERE media_id = ?`, mediaID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		releaseHistory := ReleaseHistory{}
		err := rows.Scan(&releaseHistory.ID, &releaseHistory.MediaID, &releaseHistory.LastAttempt)
		if err != nil {
			return nil, err
		}
		history = append(history, releaseHistory)
	}

	return history, rows.Err()
}

func UpdateReleaseHistory(mediaID int, releaseID string) error {
	_, err := db.Exec(`
	INSERT INTO release_history (id, media_id)
	VALUES (?, ?)
	ON CONFLICT DO UPDATE
		SET media_id=EXCLUDED.media_id, last_attempt = ?
	`, releaseID, mediaID, time.Now())
	return err
}

func DeleteReleaseHistory(releaseID string) error {
	_, err := db.Exec(`
	DELETE FROM release_history WHERE id = ?
	`, releaseID)
	return err
}
