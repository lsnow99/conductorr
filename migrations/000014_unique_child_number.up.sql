--SQLITE--
PRAGMA defer_foreign_keys = true;
--END--

BEGIN;


ALTER TABLE media RENAME TO old_media;

CREATE TABLE media(
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

INSERT INTO media SELECT * FROM old_media;

COMMIT;

--SQLITE--
PRAGMA defer_foreign_keys = false;
--END--
