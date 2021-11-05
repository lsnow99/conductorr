package dbstore

import (
	"strings"
	"testing"
)

func TestOnlySQLite(t *testing.T) {
	sql := `
	--sqlite--
	SELECT * FROM fake_table;
	`
	sqlMap := ConvertSQL(sql)

	expectedMap := map[string]string{
		sqliteKey: `SELECT * FROM fake_table;`,
		pgKey: ``,
	}

	compareSQL(t, sqlMap, expectedMap)
}

func TestSQLiteExplicitEnd(t *testing.T) {
	sql := `
	--sqlite--
	SELECT * FROM fake_table;
	--end--

	`
	sqlMap := ConvertSQL(sql)

	expectedMap := map[string]string{
		sqliteKey: `SELECT * FROM fake_table;`,
		pgKey: ``,
	}

	compareSQL(t, sqlMap, expectedMap)
}

func TestSQLiteAndPG1(t *testing.T) {
	sql := `
	--sqlite--
	SELECT * FROM fake_table;
	--postgresql--
	INSERT INTO fake_table (id, bar) VALUES (7, 'foo');
	`
	sqlMap := ConvertSQL(sql)

	expectedMap := map[string]string{
		sqliteKey: `SELECT * FROM fake_table;`,
		pgKey: `INSERT INTO fake_table (id, bar) VALUES (7, 'foo');`,
	}

	compareSQL(t, sqlMap, expectedMap)
}

func TestSQLiteAndPG2(t *testing.T) {
	sql := `
	--sqlite--
	SELECT * FROM fake_table;
	--end--
	--postgresql--
	INSERT INTO fake_table (id, bar) VALUES (7, 'foo');
	--end--
	`
	sqlMap := ConvertSQL(sql)

	expectedMap := map[string]string{
		sqliteKey: `SELECT * FROM fake_table;`,
		pgKey: `INSERT INTO fake_table (id, bar) VALUES (7, 'foo');`,
	}

	compareSQL(t, sqlMap, expectedMap)
}

func TestSQLiteAndPG3(t *testing.T) {
	sql := `
	--sqlite--
	SELECT * FROM fake_table;
	--end--
	SELECT * FROM users;
	--postgresql--
	INSERT INTO fake_table (id, bar) VALUES (7, 'foo');
	--end--
	`
	sqlMap := ConvertSQL(sql)

	expectedMap := map[string]string{
		sqliteKey: `SELECT * FROM fake_table;
	SELECT * FROM users;`,
		pgKey: `SELECT * FROM users;
	INSERT INTO fake_table (id, bar) VALUES (7, 'foo');`,
	}

	compareSQL(t, sqlMap, expectedMap)
}

func TestSQLiteAndPG4(t *testing.T) {
	sql := `
	--sqlite--
	SELECT * FROM fake_table;
	--end--
	SELECT * FROM users;
	--postgresql--
	INSERT INTO fake_table (id, bar) VALUES (7, 'foo');
	--end--
	SELECT *;
	`
	sqlMap := ConvertSQL(sql)

	expectedMap := map[string]string{
		sqliteKey: `SELECT * FROM fake_table;
	SELECT * FROM users;
	SELECT *;`,
		pgKey: `SELECT * FROM users;
	INSERT INTO fake_table (id, bar) VALUES (7, 'foo');
	SELECT *;`,
	}

	compareSQL(t, sqlMap, expectedMap)
}

func TestRealSQLiteAndPG(t *testing.T) {
	sql := `CREATE TABLE media(
		--sqlite--
		id INTEGER PRIMARY KEY,
		--postgresql--
		id SERIAL PRIMARY KEY,
		--end--
		title VARCHAR(256),
		description VARCHAR(2048),
		released_at DATETIME,
		ended_at DATETIME,
		content_type VARCHAR(32),
		--sqlite--
		poster BLOB,
		--postgresql--
		poster BYTEA,
		--end--
		parent_media_id INTEGER,
		tmdb_id INTEGER,
		imdb_id VARCHAR(10),
		tmdb_rating INTEGER,
		imdb_rating INTEGER,
		runtime INTEGER,
		status VARCHAR(32) NOT NULL DEFAULT 'missing',
		path VARCHAR(2048),
		size BIGINT,
		profile_id INTEGER REFERENCES profile(id),
		path_id INTEGER REFERENCES path(id),
		item_number INTEGER,
		monitoring BOOLEAN NOT NULL DEFAULT true,
		FOREIGN KEY(parent_media_id) REFERENCES media(id) ON DELETE CASCADE,
		CONSTRAINT uq_imdb_id UNIQUE (imdb_id)
		CONSTRAINT uq_child_num UNIQUE (item_number, parent_media_id)
	);`

	sqlMap := ConvertSQL(sql)

	expectedMap := map[string]string{
		sqliteKey: `CREATE TABLE media(
		id INTEGER PRIMARY KEY,
		title VARCHAR(256),
		description VARCHAR(2048),
		released_at DATETIME,
		ended_at DATETIME,
		content_type VARCHAR(32),
		poster BLOB,
		parent_media_id INTEGER,
		tmdb_id INTEGER,
		imdb_id VARCHAR(10),
		tmdb_rating INTEGER,
		imdb_rating INTEGER,
		runtime INTEGER,
		status VARCHAR(32) NOT NULL DEFAULT 'missing',
		path VARCHAR(2048),
		size BIGINT,
		profile_id INTEGER REFERENCES profile(id),
		path_id INTEGER REFERENCES path(id),
		item_number INTEGER,
		monitoring BOOLEAN NOT NULL DEFAULT true,
		FOREIGN KEY(parent_media_id) REFERENCES media(id) ON DELETE CASCADE,
		CONSTRAINT uq_imdb_id UNIQUE (imdb_id)
		CONSTRAINT uq_child_num UNIQUE (item_number, parent_media_id)
	);`,
		pgKey: `CREATE TABLE media(
		id SERIAL PRIMARY KEY,
		title VARCHAR(256),
		description VARCHAR(2048),
		released_at DATETIME,
		ended_at DATETIME,
		content_type VARCHAR(32),
		poster BYTEA,
		parent_media_id INTEGER,
		tmdb_id INTEGER,
		imdb_id VARCHAR(10),
		tmdb_rating INTEGER,
		imdb_rating INTEGER,
		runtime INTEGER,
		status VARCHAR(32) NOT NULL DEFAULT 'missing',
		path VARCHAR(2048),
		size BIGINT,
		profile_id INTEGER REFERENCES profile(id),
		path_id INTEGER REFERENCES path(id),
		item_number INTEGER,
		monitoring BOOLEAN NOT NULL DEFAULT true,
		FOREIGN KEY(parent_media_id) REFERENCES media(id) ON DELETE CASCADE,
		CONSTRAINT uq_imdb_id UNIQUE (imdb_id)
		CONSTRAINT uq_child_num UNIQUE (item_number, parent_media_id)
	);`,
	}

	compareSQL(t, sqlMap, expectedMap)
}

func compareSQL(t *testing.T, sqlMap, expectedMap map[string]string) {
	if len(sqlMap) != len(expectedMap) {
		t.Fatal("returned sql map number of keys differs from the expected map")
	}
	for k, v1 := range expectedMap {
		v2, ok := sqlMap[k]
		if !ok {
			t.Fatalf("expected key %s in result map, but was not found", k)
		}
		tv1 := strings.TrimSpace(v1)
		tv2 := strings.TrimSpace(v2)
		if tv1 != tv2 {
			t.Fatalf("expected value %s for key %s but got %s", tv1, k, tv2)
		}
	}
}
