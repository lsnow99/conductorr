package dbstore

import (
	"context"
	"database/sql"
)


func GetPaths() ([]Path, error) {
	paths := make([]Path, 0)

	rows, err := db.Query(`SELECT id, path, movies_default, series_default FROM path;`)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	} else if err == sql.ErrNoRows {
		return paths, nil
	}

	for rows.Next() {
		path := Path{}
		err := rows.Scan(&path.ID, &path.Path, &path.MoviesDefault, &path.SeriesDefault)
		if err != nil {
			return nil, err
		}
		paths = append(paths, path)
	}

	return paths, nil
}

func UpdatePaths(ctx context.Context, paths []Path) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	_, err = tx.Exec(`DELETE FROM path WHERE true;`)
	if err != nil {
		tx.Rollback()
		return err
	}
	for _, path := range paths {
		_, err := tx.Exec(`
			INSERT INTO path (path, movies_default, series_default)
			VALUES(?, ?, ?)
			`, path.Path, path.MoviesDefault, path.SeriesDefault)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}

func DeletePath(id int) error {
	_, err := db.Exec(`DELETE FROM path WHERE id = ?`, id)
	return err
}

func GetPath(id int) (Path, error) {
	row := db.QueryRow(`SELECT id, path, movies_default, series_default WHERE id = ?`, id)
	path := Path{}
	return path, row.Scan(&path.ID, &path.Path, &path.MoviesDefault, &path.SeriesDefault)
}