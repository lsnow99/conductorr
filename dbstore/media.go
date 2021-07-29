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
			runtime, profile_id, path_id, item_number, monitoring
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
			&media.Runtime, &media.ProfileID, &media.PathID, &media.Number, &media.Monitoring); 
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
			runtime, profile_id, path_id, item_number, monitoring
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
			&media.Runtime, &media.ProfileID, &media.PathID, &media.Number, &media.Monitoring); 
			err != nil {
			return nil, err
		}
		medias = append(medias, &media)
	}
	return medias, nil
}

func AddMedia(title *string, description *string, releasedAt *time.Time, endedAt *time.Time,
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
		RETURNING id
		`, ptrToNullString(title), ptrToNullString(description), released,
		ended, ptrToNullString(contentType), ptrToNullInt32(parentMediaID),
		ptrToNullInt32(tmdbID), ptrToNullString(imdbID), ptrToNullInt32(tmdbID),
		ptrToNullInt32(imdbRating), ptrToNullInt32(runtime), poster, ptrToNullInt32(profileID), 
		ptrToNullInt32(pathID), ptrToNullInt32(number), monitoring)

	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}

	for _, genre := range genres {
		_, err := tx.Exec(`
		INSERT INTO media_genre (media_id, genre_id) VALUES (
			?,
			(SELECT id FROM genre WHERE UPPER(name)=UPPER(?))
		)
		`, id, genre)
		if err != nil {
			return 0, err
		}
	}
	err = tx.Commit()
	return id, err
}

func GetMediaByImdbID(imdbID string) (media Media, err error) {
	row := db.QueryRow(`
		SELECT id, title, description, released_at, ended_at, content_type,
			parent_media_id, tmdb_id, imdb_id, tmdb_rating, imdb_rating,
			runtime, status, path, size, item_number, monitoring
		FROM media
		WHERE imdb_id = ?
		`, imdbID)

	err = row.Scan(&media.ID, &media.Title, &media.Description, &media.ReleasedAt,
		&media.EndedAt, &media.ContentType, &media.ParentMediaID,
		&media.TmdbID, &media.ImdbID, &media.TmdbRating, &media.ImdbRating,
		&media.Runtime, &media.Status, &media.Path, &media.Size, &media.Number,
		&media.Monitoring)

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
			runtime, profile_id, path_id, item_number, monitoring
		FROM media
		WHERE id = ?
		`, id)

	err = row.Scan(&media.ID, &media.Title, &media.Description, &media.ReleasedAt,
		&media.EndedAt, &media.ContentType, &media.ParentMediaID,
		&media.TmdbID, &media.ImdbID, &media.TmdbRating, &media.ImdbRating,
		&media.Runtime, &media.ProfileID, &media.PathID, &media.Number, &media.Monitoring)

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
			runtime, profile_id, path_id, item_number, monitoring
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
			&media.Runtime, &media.ProfileID, &media.PathID, &media.Number, &media.Monitoring); 
			err != nil {
			return nil, err
		}
		medias = append(medias, &media)
	}
	return medias, nil
}

func UpdateMediaMonitoring(mediaID int, monitoring bool) error {
	_, err := db.Exec(`
	UPDATE media SET monitoring = ? WHERE id = ?;
	UPDATE media SET monitoring = ? WHERE parent_media_id = ? AND EXISTS (SELECT * FROM media WHERE id = ? AND content_type = 'season')
	`, monitoring, mediaID, monitoring, mediaID, mediaID)
	return err
}