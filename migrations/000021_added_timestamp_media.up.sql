
--SQLITE--
PRAGMA foreign_keys=off;
--END--

BEGIN;

DROP TABLE IF EXISTS old_media;

CREATE TABLE old_media(
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
    tvdb_id INTEGER,
    --postgresql--
    added DATETIME NOT NULL DEFAULT NOW(),
    --sqlite--
    added DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    --end--
    FOREIGN KEY(parent_media_id) REFERENCES media(id) ON DELETE CASCADE,
    CONSTRAINT uq_imdb_id UNIQUE (imdb_id)
    CONSTRAINT uq_child_num UNIQUE (item_number, parent_media_id)
);

INSERT INTO old_media (id, title, description, released_at, ended_at, content_type, 
    poster, parent_media_id, tmdb_id, imdb_id, tmdb_rating, imdb_rating,
    runtime, status, path, size, profile_id, path_id, item_number, monitoring,
    tvdb_id) 
SELECT id, title, description, released_at, ended_at, content_type, 
    poster, parent_media_id, tmdb_id, imdb_id, tmdb_rating, imdb_rating,
    runtime, status, path, size, profile_id, path_id, item_number, monitoring,
    tvdb_id 
FROM media;

DROP TABLE media;

ALTER TABLE old_media RENAME TO media;

COMMIT;

--SQLITE--
PRAGMA foreign_keys=on;
--END--
