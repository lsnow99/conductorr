package dbstore

import (
	"context"
	"database/sql"
	"strings"
	"time"
)

func SearchMedia(title string, contentType string, page int) ([]*Media, int, error) {
	title = strings.ToUpper(title)
	row := db.QueryRow(`
		SELECT COUNT(id)
		FROM media
		WHERE UPPER(title) LIKE '%' || ? || '%'
		AND content_type LIKE '%' || ? || '%'
		AND content_type NOT IN ('episode', 'season')
		`, title, contentType)

	var count int

	if err := row.Scan(&count); err != nil {
		return nil, 0, err
	}

	rows, err := db.Query(`
		SELECT id, title, description, released_at, ended_at, content_type,
			parent_media_id, tmdb_id, imdb_id, tmdb_rating, imdb_rating,
			runtime, status, profile_id, path_id, size, item_number, monitoring, path
		FROM media
		WHERE UPPER(title) LIKE '%' || ? || '%' 
		AND content_type LIKE '%' || ? || '%'
		AND content_type NOT IN ('episode', 'season')
		LIMIT 10 
		OFFSET ?
		`, title, contentType, (page-1)*10)

	if err != nil && err != sql.ErrNoRows {
		return nil, 0, err
	}
	defer rows.Close()

	medias := make([]*Media, 0, 10)
	for rows.Next() {
		media := Media{}
		if err := rows.Scan(&media.ID, &media.Title, &media.Description, &media.ReleasedAt,
			&media.EndedAt, &media.ContentType, &media.ParentMediaID,
			&media.TmdbID, &media.ImdbID, &media.TmdbRating, &media.ImdbRating,
			&media.Runtime, &media.Status, &media.ProfileID, &media.PathID, &media.Size, 
			&media.Number, &media.Monitoring, &media.Path); 
			err != nil {
			return nil, 0, err
		}
		medias = append(medias, &media)
	}

	return medias, count, nil
}

func GetAllMedia() ([]*Media, error) {
	rows, err := db.Query(`
		SELECT id, title, description, released_at, ended_at, content_type,
			parent_media_id, tmdb_id, imdb_id, tmdb_rating, imdb_rating,
			runtime, status, profile_id, path_id, size, item_number, monitoring, path
		FROM media
		WHERE content_type in ('episode', 'movie')
		`)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	defer rows.Close()

	medias := make([]*Media, 0, 10)
	for rows.Next() {
		media := Media{}
		if err := rows.Scan(&media.ID, &media.Title, &media.Description, &media.ReleasedAt,
			&media.EndedAt, &media.ContentType, &media.ParentMediaID,
			&media.TmdbID, &media.ImdbID, &media.TmdbRating, &media.ImdbRating,
			&media.Runtime, &media.Status, &media.ProfileID, &media.PathID, &media.Size, 
			&media.Number, &media.Monitoring, &media.Path); 
			err != nil {
			return nil, err
		}
		medias = append(medias, &media)
	}
	return medias, nil
}

func GetAllMediaMap() (map[int]*Media, error) {
	rows, err := db.Query(`
		SELECT id, title, description, released_at, ended_at, content_type,
			parent_media_id, tmdb_id, imdb_id, tmdb_rating, imdb_rating,
			runtime, status, profile_id, path_id, size, item_number, monitoring, path
		FROM media
		`)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	defer rows.Close()

	medias := make(map[int]*Media)
	for rows.Next() {
		media := Media{}
		if err := rows.Scan(&media.ID, &media.Title, &media.Description, &media.ReleasedAt,
			&media.EndedAt, &media.ContentType, &media.ParentMediaID,
			&media.TmdbID, &media.ImdbID, &media.TmdbRating, &media.ImdbRating,
			&media.Runtime, &media.Status, &media.ProfileID, &media.PathID, &media.Size, 
			&media.Number, &media.Monitoring, &media.Path); 
			err != nil {
			return nil, err
		}
		medias[media.ID] = &media
	}
	return medias, nil
}

/* UpsertMedia inserts the new media described by the parameters, and on a conflict of the item_number
and parent_media_id, it will only update the metadata
*/
func UpsertMedia(title *string, description *string, releasedAt *time.Time, endedAt *time.Time,
	contentType *string, parentMediaID *int, tmdbID *int, imdbID *string,
	tmdbRating *int, imdbRating *int, runtime *int, poster *[]byte, genres []string, profileID,
	pathID *int, number *int, monitoring bool) (id int, err error) {

	tx, err := db.BeginTx(context.TODO(), nil)
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	if err != nil {
		return 0, err
	}

	released := ptrToNullTime(releasedAt)
	ended := ptrToNullTime(endedAt)

	row := tx.QueryRow(`
		INSERT INTO media (title, description, released_at, ended_at, content_type,
			parent_media_id, tmdb_id, imdb_id, tmdb_rating, imdb_rating, runtime, 
			poster, profile_id, path_id, item_number, monitoring)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		ON CONFLICT DO UPDATE
			SET title=EXCLUDED.title, description=EXCLUDED.description, 
				released_at=EXCLUDED.released_at, ended_at=EXCLUDED.ended_at,
				content_type=EXCLUDED.content_type,
				tmdb_id=EXCLUDED.tmdb_id, imdb_id=EXCLUDED.imdb_id, 
				tmdb_rating=EXCLUDED.tmdb_rating, imdb_rating=EXCLUDED.imdb_rating, 
				runtime=EXCLUDED.runtime, poster=EXCLUDED.poster
		RETURNING id
		`, ptrToNullString(title), ptrToNullString(description), released,
		ended, ptrToNullString(contentType), ptrToNullInt32(parentMediaID),
		ptrToNullInt32(tmdbID), ptrToNullString(imdbID), ptrToNullInt32(tmdbRating),
		ptrToNullInt32(imdbRating), ptrToNullInt32(runtime), poster, ptrToNullInt32(profileID), 
		ptrToNullInt32(pathID), ptrToNullInt32(number), monitoring)

	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}

	// for _, genre := range genres {
	// 	_, err := tx.Exec(`
	// 	INSERT INTO media_genre (media_id, genre_id) VALUES (
	// 		?,
	// 		(SELECT id FROM genre WHERE UPPER(name)=UPPER(?))
	// 	)
	// 	`, id, genre)
	// 	if err != nil {
	// 		return 0, err
	// 	}
	// }
	err = tx.Commit()
	return id, err
}

func GetMediaByImdbID(imdbID string) (media Media, err error) {
	row := db.QueryRow(`
		SELECT id, title, description, released_at, ended_at, content_type,
			parent_media_id, tmdb_id, imdb_id, tmdb_rating, imdb_rating,
			runtime, status, profile_id, path_id, size, item_number, monitoring, path
		FROM media
		WHERE imdb_id = ?
		`, imdbID)

	err = row.Scan(&media.ID, &media.Title, &media.Description, &media.ReleasedAt,
		&media.EndedAt, &media.ContentType, &media.ParentMediaID,
		&media.TmdbID, &media.ImdbID, &media.TmdbRating, &media.ImdbRating,
		&media.Runtime, &media.Status, &media.ProfileID, &media.PathID, &media.Size, 
		&media.Number, &media.Monitoring)

	return media, err
}

func GetPoster(mediaID int) (poster []byte, err error) {
	row := db.QueryRow(`
		SELECT poster
		FROM media
		WHERE id = ?
	`, mediaID)

	err = row.Scan(&poster)

	return poster, err
}

func GetMediaByID(id int) (media Media, err error) {
	row := db.QueryRow(`
		SELECT id, title, description, released_at, ended_at, content_type,
			parent_media_id, tmdb_id, imdb_id, tmdb_rating, imdb_rating,
			runtime, status, profile_id, path_id, size, item_number, monitoring, path
		FROM media
		WHERE id = ?
		`, id)

		err = row.Scan(&media.ID, &media.Title, &media.Description, &media.ReleasedAt,
			&media.EndedAt, &media.ContentType, &media.ParentMediaID,
			&media.TmdbID, &media.ImdbID, &media.TmdbRating, &media.ImdbRating,
			&media.Runtime, &media.Status, &media.ProfileID, &media.PathID, &media.Size, 
			&media.Number, &media.Monitoring, &media.Path)

	return media, err
}

func DeleteMedia(id int) error {
	_, err := db.Exec(`
		DELETE FROM media
		WHERE id = ?
		`, id)
	return err
}

func UpdateMedia(id int, profileID, pathID int) (err error) {
	_, err = db.Exec(`
		UPDATE media SET profile_id = ?, path_id = ? WHERE id = ?
		`, profileID, pathID, id)
	return err
}

func GetMediaReferencing(parentID int) ([]*Media, error) {
	rows, err := db.Query(`
		SELECT id, title, description, released_at, ended_at, content_type,
			parent_media_id, tmdb_id, imdb_id, tmdb_rating, imdb_rating,
			runtime, status, profile_id, path_id, size, item_number, monitoring, path
		FROM media
		WHERE parent_media_id = ?
		`, parentID)

	if err != nil && err != sql.ErrNoRows{
		return nil, err
	}
	defer rows.Close()

	medias := make([]*Media, 0, 10)
	for rows.Next() {
		media := Media{}
		if err := rows.Scan(&media.ID, &media.Title, &media.Description, &media.ReleasedAt,
			&media.EndedAt, &media.ContentType, &media.ParentMediaID,
			&media.TmdbID, &media.ImdbID, &media.TmdbRating, &media.ImdbRating,
			&media.Runtime, &media.Status, &media.ProfileID, &media.PathID, &media.Size, 
			&media.Number, &media.Monitoring, &media.Path)
			err != nil {
			return nil, err
		}
		medias = append(medias, &media)
	}
	return medias, nil
}

func UpdateMediaMonitoring(id int, monitoring bool) error {
	_, err := db.Exec(`
	UPDATE media SET monitoring = ? WHERE id = ?;
	UPDATE media SET monitoring = ? WHERE parent_media_id = ? AND EXISTS (SELECT * FROM media WHERE id = ? AND content_type = 'season')
	`, monitoring, id, monitoring, id, id)
	return err
}

func SetMediaPath(id int, path string) error {
	_, err := db.Exec(`
	UPDATE media SET path = ? WHERE id = ?
	`, path, id)
	return err
}

func GetMonitoringMedia() (*[]Media, error) {
	medias := make([]Media, 0, 100)
	rows, err := db.Query(`
		SELECT id, title, description, released_at, ended_at, content_type,
			parent_media_id, tmdb_id, imdb_id, tmdb_rating, imdb_rating,
			runtime, status, profile_id, path_id, size, item_number, monitoring, path
		FROM media AS m
		WHERE (monitoring = true AND content_type = 'movie')
		OR (monitoring = true AND content_type = 'episode'
			AND (SELECT monitoring FROM media AS g WHERE id = 
				(SELECT parent_media_id FROM media AS p WHERE id = m.parent_media_id)
			)
		   )
		`)
	
	if err != nil && err != sql.ErrNoRows{
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		media := Media{}
		if err := rows.Scan(&media.ID, &media.Title, &media.Description, &media.ReleasedAt,
			&media.EndedAt, &media.ContentType, &media.ParentMediaID,
			&media.TmdbID, &media.ImdbID, &media.TmdbRating, &media.ImdbRating,
			&media.Runtime, &media.Status, &media.ProfileID, &media.PathID, &media.Size, 
			&media.Number, &media.Monitoring, &media.Path)
			err != nil {
			return nil, err
		}
		medias = append(medias, media)
	}
	return &medias, nil
}

func DumpMedias(tx *sql.Tx) ([]*Media, error) {
	rows, err := db.Query(`
		SELECT id, title, description, released_at, ended_at, content_type, poster,
			parent_media_id, tmdb_id, imdb_id, tmdb_rating, imdb_rating,
			runtime, status, profile_id, path_id, size, item_number, monitoring, path
		FROM media
		`)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	defer rows.Close()

	medias := make([]*Media, 0, 10)
	for rows.Next() {
		media := Media{}
		if err := rows.Scan(&media.ID, &media.Title, &media.Description, &media.ReleasedAt,
			&media.EndedAt, &media.ContentType, &media.Poster, &media.ParentMediaID,
			&media.TmdbID, &media.ImdbID, &media.TmdbRating, &media.ImdbRating,
			&media.Runtime, &media.Status, &media.ProfileID, &media.PathID, &media.Size, 
			&media.Number, &media.Monitoring, &media.Path); 
			err != nil {
			return nil, err
		}
		medias = append(medias, &media)
	}
	return medias, nil
}

func RestoreMedias(tx *sql.Tx, medias []*Media) error {
	for _, media := range medias {
		_, err := tx.Exec(`
		INSERT INTO media (id, title, description, released_at, ended_at, content_type,
			parent_media_id, tmdb_id, imdb_id, tmdb_rating, imdb_rating, runtime, 
			poster, profile_id, path_id, item_number, monitoring)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		`, media.ID, media.Title, media.Description, media.ReleasedAt, media.EndedAt,
			media.ContentType, media.ParentMediaID, media.TmdbID, media.ImdbID,
			media.ImdbRating, media.Runtime, media.Poster, media.ProfileID, media.PathID,
			media.Number, media.Monitoring)

		if err != nil {
			return err
		}
	}

	return nil
}
