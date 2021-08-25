BEGIN;

CREATE TABLE user(
    --sqlite--
    id INTEGER PRIMARY KEY,
    --postgresql--
    id SERIAL PRIMARY KEY,
    --end
    username VARCHAR(128) UNIQUE,
    password VARBINARY(128)
);

CREATE TABLE media(
    --sqlite--
    id INTEGER PRIMARY KEY,
    --postgresql--
    id SERIAL PRIMARY KEY,
    --end
    title VARCHAR(256),
    description VARCHAR(2048),
    released_at DATETIME,
    ended_at DATETIME,
    content_type VARCHAR(32),
    --sqlite--
    poster BLOB,
    --postgresql--
    poster BYTEA,
    --end
    parent_media_id INTEGER,
    tmdb_id INTEGER,
    imdb_id VARCHAR(10),
    tmdb_rating INTEGER,
    imdb_rating INTEGER,
    runtime INTEGER,
    status VARCHAR(32) NOT NULL DEFAULT 'missing',
    path VARCHAR(2048),
    size BIGINT,
    FOREIGN KEY(parent_media_id) REFERENCES media(id) ON DELETE CASCADE,
    CONSTRAINT uq_imdb_id UNIQUE (imdb_id)
);

COMMIT;