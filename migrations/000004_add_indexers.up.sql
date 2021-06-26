BEGIN;

CREATE TABLE indexer(
    id INTEGER PRIMARY KEY,
    name VARCHAR(128) NOT NULL,
    user_id INTEGER,
    base_url TEXT NOT NULL,
    api_key TEXT NOT NULL,
    for_movies BOOLEAN NOT NULL,
    for_series BOOLEAN NOT NULL,
    download_type VARCHAR(32) NOT NULL
);

COMMIT;