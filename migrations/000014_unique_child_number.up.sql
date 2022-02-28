--SQLITE--
PRAGMA foreign_keys=off;
--END--

BEGIN;

CREATE TABLE new_media(
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
);

INSERT INTO new_media SELECT * FROM media;

DROP TABLE media;

ALTER TABLE new_media RENAME TO media;

COMMIT;

--SQLITE--
PRAGMA foreign_keys=on;
--END--
