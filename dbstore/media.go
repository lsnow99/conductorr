package dbstore

import "time"

func SearchMedia(title string, contentType string, page int) ([]*Media, int, error) {
	row := db.QueryRow(`
		SELECT COUNT(id)
		FROM media
		WHERE UPPER(title) LIKE '%?%'
		`, title)

	var count int

	if err := row.Scan(&count); err != nil {
		return nil, 0, err
	}

	rows, err := db.Query(`
		SELECT id, title, description, released_at, ended_at, content_type,
			parent_media_id, tmdb_id, imdb_id, tmdb_rating, imdb_rating,
			runtime
		FROM media
		WHERE UPPER(title) LIKE '%?%' LIMIT 10 OFFSET ?
		`, title, title, (page-1)*10)

	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	medias := make([]*Media, 0, 10)
	for rows.Next() {
		media := Media{}
		if err := rows.Scan(&media.ID, &media.Title, &media.Description, &media.ReleasedAt,
			&media.EndedAt, &media.ContentType, &media.ParentMediaID,
			&media.TmdbID, &media.ImdbID, &media.TmdbRating, &media.ImdbRating,
			&media.Runtime); err != nil {
			return nil, 0, err
		}
		medias = append(medias, &media)
	}

	return medias, 0, nil
}

func AddMedia(title *string, description *string, releasedAt *time.Time, endedAt *time.Time,
	contentType *string, parentMediaID *int, tmdbID *int, imdbID *string,
	tmdbRating *int, imdbRating *int, runtime *int, poster *[]byte, genres []string) (id int, err error) {

	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}

	row := tx.QueryRow(`
		INSERT INTO media (title, description, released_at, ended_at, content_type,
			parent_media_id, tmdb_id, imdb_id, tmdb_rating, imdb_rating, runtime, poster)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		RETURNING id
		`, ptrToNullString(title), ptrToNullString(description), releasedAt, endedAt,
		ptrToNullString(contentType), ptrToNullInt32(parentMediaID), ptrToNullInt32(tmdbID),
		ptrToNullString(imdbID), ptrToNullInt32(tmdbID), ptrToNullInt32(imdbRating),
		ptrToNullInt32(runtime), poster)

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
	return id, tx.Commit()
}

func GetMediaByImdbID(imdbID string) (media Media, err error) {
	row := db.QueryRow(`
		SELECT id, title, description, released_at, ended_at, content_type,
			parent_media_id, tmdb_id, imdb_id, tmdb_rating, imdb_rating,
			runtime
		FROM media
		WHERE imdb_id = ?
		`, imdbID)

	err = row.Scan(&media.ID, &media.Title, &media.Description, &media.ReleasedAt,
		&media.EndedAt, &media.ContentType, &media.ParentMediaID,
		&media.TmdbID, &media.ImdbID, &media.TmdbRating, &media.ImdbRating,
		&media.Runtime)

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