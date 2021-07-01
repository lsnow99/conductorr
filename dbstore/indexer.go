package dbstore

import "database/sql"

func CreateIndexer(name string, userID *int, baseUrl, apiKey string, forMovies, forSeries bool, downloadType string) error {
	_, err := db.Exec(`INSERT INTO indexer (name, user_id, base_url, api_key, for_movies, for_series, download_type) VALUES (?, ?, ?, ?, ?, ?, ?)`, name, userID, baseUrl, apiKey, forMovies, forSeries, downloadType)
	return err
}

func GetIndexers() (indexers []Indexer, err error) {
	rows, err := db.Query(`SELECT id, name, user_id, base_url, api_key, for_movies, for_series, download_type FROM indexer;`)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	} else if err == sql.ErrNoRows {
		return indexers, nil
	}

	for rows.Next() {
		indexer := Indexer{}
		err := rows.Scan(&indexer.ID, &indexer.Name, &indexer.UserID, &indexer.BaseUrl, &indexer.ApiKey, &indexer.ForMovies, &indexer.ForSeries, &indexer.DownloadType)
		if err != nil {
			return nil, err
		}
		indexers = append(indexers, indexer)
	}

	return indexers, nil
}

func GetIndexerByID(id int) (Indexer, error) {
	indexer := Indexer{}
	row := db.QueryRow(`SELECT id, name, user_id, base_url, api_key, for_movies, for_series, download_type FROM indexer WHERE id = ?`, id)

	err := row.Scan(&indexer.ID, &indexer.Name, &indexer.UserID, &indexer.BaseUrl, &indexer.ApiKey, &indexer.ForMovies, &indexer.ForSeries, &indexer.DownloadType)
	return indexer, err
}

func UpdateIndexer(id int, name string, userID *int, baseUrl, apiKey string, forMovies, forSeries bool) error {
	_, err := db.Exec(`UPDATE indexer SET name = ?, user_id = ?, base_url = ?, api_key = ?, for_movies = ?, for_series = ? WHERE id = ?`, name, userID, baseUrl, apiKey, forMovies, forSeries, id)
	return err
}

func DeleteIndexer(id int) error {
	_, err := db.Exec(`DELETE FROM indexer WHERE id = ?`, id)
	return err
}