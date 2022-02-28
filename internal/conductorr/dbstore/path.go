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
	defer rows.Close()

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

func NewPath(path string, moviesDefault, seriesDefault bool) error {
	tx, err := db.BeginTx(context.TODO(), nil)
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	if moviesDefault {
		_, err = tx.Exec(`UPDATE path SET movies_default = ? WHERE true`, false)
		if err != nil {
			return err
		}
	}

	if seriesDefault {
		_, err = tx.Exec(`UPDATE path SET series_default = ? WHERE true`, false)
		if err != nil {
			return err
		}
	}

	_, err = tx.Exec(`
	INSERT INTO path (path, movies_default, series_default)
	VALUES(?, ?, ?)`, path, moviesDefault, seriesDefault)
	if err != nil {
		return err
	}

	err = tx.Commit()
	return err
}

func UpdatePath(id int, path string, moviesDefault, seriesDefault bool) error {
	tx, err := db.BeginTx(context.TODO(), nil)
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	if moviesDefault {
		_, err = tx.Exec(`UPDATE path SET movies_default = ? WHERE true`, false)
		if err != nil {
			return err
		}
	}

	if seriesDefault {
		_, err = tx.Exec(`UPDATE path SET series_default = ? WHERE true`, false)
		if err != nil {
			return err
		}
	}

	_, err = tx.Exec(`
	UPDATE path SET path = ?, movies_default = ?, series_default = ?
	WHERE id = ?`, path, moviesDefault, seriesDefault, id)
	if err != nil {
		return err
	}

	err = tx.Commit()
	return err
}

func DeletePath(id int) error {
	_, err := db.Exec(`DELETE FROM path WHERE id = ?`, id)
	return err
}

func GetPath(id int) (Path, error) {
	row := db.QueryRow(`SELECT id, path, movies_default, series_default FROM path WHERE id = ?`, id)
	path := Path{}
	return path, row.Scan(&path.ID, &path.Path, &path.MoviesDefault, &path.SeriesDefault)
}

func DumpPaths(tx *sql.Tx) ([]*Path, error) {
	paths := make([]*Path, 0)

	rows, err := db.Query(`SELECT id, path, movies_default, series_default FROM path;`)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	} else if err == sql.ErrNoRows {
		return paths, nil
	}
	defer rows.Close()

	for rows.Next() {
		path := &Path{}
		err := rows.Scan(&path.ID, &path.Path, &path.MoviesDefault, &path.SeriesDefault)
		if err != nil {
			return nil, err
		}
		paths = append(paths, path)
	}

	return paths, rows.Err()
}

func RestorePaths(tx *sql.Tx, paths []*Path) error {
	for _, path := range paths {
		_, err := tx.Exec(`
		INSERT INTO path (id, path, movies_default, series_default)
		VALUES(?, ?, ?, ?)`, path.ID, path.Path, path.MoviesDefault, path.SeriesDefault)

		if err != nil {
			return err
		}
	}

	return nil
}
