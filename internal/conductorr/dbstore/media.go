package dbstore

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/huandu/go-sqlbuilder"
)

var mediaStruct = sqlbuilder.NewStruct(new(Media))

const fullMediaCols = ` id, title, description, released_at, ended_at, content_type,
parent_media_id, tmdb_id, imdb_id, tvdb_id, tmdb_rating, imdb_rating,
runtime, status, profile_id, path_id, size, item_number, monitoring, path,
added `

type Scannable interface {
	Scan(dest ...interface{}) error
}

func SelectAllCols(sb *sqlbuilder.SelectBuilder) *sqlbuilder.SelectBuilder {
	return sb.Select("id, title, description, released_at, ended_at, content_type, " +
		"parent_media_id, tmdb_id, imdb_id, tvdb_id, tmdb_rating, imdb_rating, " +
		"runtime, status, profile_id, path_id, size, item_number, monitoring, path, " +
		"added")
}

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

	/*

		CASE content_type
			WHEN 'series'
				THEN (SELECT COUNT(id) FROM media WHERE )

		END
	*/

	rows, err := db.Query(`
		SELECT`+fullMediaCols+`
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
		media, err := scanFullMedia(rows)
		if err != nil {
			return nil, 0, err
		}
		medias = append(medias, &media)
	}

	return medias, count, nil
}

func GetAllMedia() ([]*Media, error) {
	rows, err := db.Query(`
		SELECT` + fullMediaCols + `
		FROM media
		WHERE content_type in ('episode', 'movie')
		`)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	defer rows.Close()

	medias := make([]*Media, 0, 10)
	for rows.Next() {
		media, err := scanFullMedia(rows)
		if err != nil {
			return nil, err
		}
		medias = append(medias, &media)
	}
	return medias, nil
}

type DateInterval struct {
	DateFrom time.Time
	DateTo   time.Time
}

// GetMediaInIntervals select an array of media within a date range. Includes all series and season medias
func GetMediaInIntervals(ctx context.Context, dateIntervals []DateInterval) ([]Media, error) {
	sb := mediaStruct.SelectFrom("media")

	dateConditions := make([]string, 0, len(dateIntervals)+1)

	for _, dateInterval := range dateIntervals {
		dateConditions = append(dateConditions, sb.And(sb.GreaterEqualThan("released_at", dateInterval.DateFrom), sb.LessEqualThan("released_at", dateInterval.DateTo)))
	}

	dateConditions = append(dateConditions, sb.In("content_type", "series", "season"))

	sb.Where(
		sb.Or(
			dateConditions...,
		),
	)

	stmt, args := sb.Build()
	rows, err := db.QueryContext(ctx, stmt, args...)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	defer rows.Close()

	fmt.Println(sb.String())

	medias := make([]Media, 0)
	for rows.Next() {
		media := Media{}
		err := rows.Scan(mediaStruct.Addr(&media)...)
		if err != nil {
			return nil, err
		}
		medias = append(medias, media)
	}
	return medias, nil
}

/* UpsertMedia inserts the new media described by the parameters, and on a conflict of the item_number
and parent_media_id, it will only update the metadata
*/
func UpsertMedia(title *string, description *string, releasedAt *time.Time, endedAt *time.Time,
	contentType *string, parentMediaID *int, tmdbID *int, imdbID *string, tvdbID *int,
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
			parent_media_id, tmdb_id, imdb_id, tvdb_id, tmdb_rating, imdb_rating, runtime,
			poster, profile_id, path_id, item_number, monitoring)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		ON CONFLICT DO UPDATE
			SET title=EXCLUDED.title, description=EXCLUDED.description,
				released_at=EXCLUDED.released_at, ended_at=EXCLUDED.ended_at,
				content_type=EXCLUDED.content_type,
				tmdb_id=EXCLUDED.tmdb_id, imdb_id=EXCLUDED.imdb_id, tvdb_id=EXCLUDED.tvdb_id,
				tmdb_rating=EXCLUDED.tmdb_rating, imdb_rating=EXCLUDED.imdb_rating,
				runtime=EXCLUDED.runtime, poster=EXCLUDED.poster
		RETURNING id
		`, ptrToNullString(title), ptrToNullString(description), released,
		ended, ptrToNullString(contentType), ptrToNullInt32(parentMediaID),
		ptrToNullInt32(tmdbID), ptrToNullString(imdbID), ptrToNullInt32(tvdbID), ptrToNullInt32(tmdbRating),
		ptrToNullInt32(imdbRating), ptrToNullInt32(runtime), poster, ptrToNullInt32(profileID),
		ptrToNullInt32(pathID), ptrToNullInt32(number), monitoring)

	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	return id, err
}

func GetMediaByImdbID(imdbID string) (Media, error) {
	row := db.QueryRow(`
		SELECT`+fullMediaCols+`
		FROM media
		WHERE imdb_id = ?
		`, imdbID)

	return scanFullMedia(row)
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

func GetMediaByID(id int) (Media, error) {
	row := db.QueryRow(`
		SELECT`+fullMediaCols+`
		FROM media
		WHERE id = ?
		`, id)

	return scanFullMedia(row)
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
		SELECT`+fullMediaCols+`
		FROM media
		WHERE parent_media_id = ?
		`, parentID)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	defer rows.Close()

	medias := make([]*Media, 0, 10)
	for rows.Next() {
		media, err := scanFullMedia(rows)
		if err != nil {
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
		SELECT` + fullMediaCols + `
		FROM media AS m
		WHERE (monitoring = true AND content_type = 'movie')
		OR (monitoring = true AND content_type = 'episode'
			AND (SELECT monitoring FROM media AS g WHERE id =
				(SELECT parent_media_id FROM media AS p WHERE id = m.parent_media_id)
			)
		   )
		`)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		media, err := scanFullMedia(rows)
		if err != nil {
			return nil, err
		}
		medias = append(medias, media)
	}
	return &medias, nil
}

func GetRecentlyAddedMedia(max int) ([]Media, error) {
	medias := make([]Media, 0, max)

	rows, err := db.Query(`
		SELECT`+fullMediaCols+`
		FROM media AS m
		WHERE content_type = 'series'
		OR content_type = 'movie'
		ORDER BY added DESC
		LIMIT ?;
	`, max)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		media, err := scanFullMedia(rows)
		if err != nil {
			return nil, err
		}
		medias = append(medias, media)
	}
	return medias, nil
}

func DumpMedias(tx *sql.Tx) ([]*Media, error) {
	rows, err := tx.Query(`
		SELECT id, title, description, released_at, ended_at, content_type, poster,
			parent_media_id, tmdb_id, imdb_id, tvdb_id, tmdb_rating, imdb_rating,
			runtime, status, profile_id, path_id, size, item_number, monitoring, path,
			added
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
			&media.TmdbID, &media.ImdbID, &media.TvdbID, &media.TmdbRating, &media.ImdbRating,
			&media.Runtime, &media.Status, &media.ProfileID, &media.PathID, &media.Size,
			&media.ItemNumber, &media.Monitoring, &media.Path, &media.Added); err != nil {
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
			parent_media_id, tmdb_id, imdb_id, tvdb_id, tmdb_rating, imdb_rating, runtime,
			poster, profile_id, path_id, item_number, monitoring, added)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		`, media.ID, media.Title, media.Description, media.ReleasedAt, media.EndedAt,
			media.ContentType, media.ParentMediaID, media.TmdbID, media.ImdbID, media.TvdbID,
			media.ImdbRating, media.Runtime, media.Poster, media.ProfileID, media.PathID,
			media.ItemNumber, media.Monitoring, media.Added)

		if err != nil {
			return err
		}
	}

	return nil
}

func scanFullMedia(rows Scannable) (Media, error) {
	media := Media{}
	if err := rows.Scan(&media.ID, &media.Title, &media.Description, &media.ReleasedAt,
		&media.EndedAt, &media.ContentType, &media.ParentMediaID,
		&media.TmdbID, &media.ImdbID, &media.TvdbID, &media.TmdbRating, &media.ImdbRating,
		&media.Runtime, &media.Status, &media.ProfileID, &media.PathID, &media.Size,
		&media.ItemNumber, &media.Monitoring, &media.Path, &media.Added); err != nil {
		return media, err
	}
	return media, nil
}
