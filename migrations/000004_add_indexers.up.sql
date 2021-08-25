BEGIN;

CREATE TABLE indexer(
    --sqlite--
    id INTEGER PRIMARY KEY,
    --postgresql--
    id SERIAL PRIMARY KEY,
    --end
    name VARCHAR(128) NOT NULL,
    user_id INTEGER NOT NULL,
    base_url TEXT NOT NULL,
    api_key TEXT NOT NULL,
    for_movies BOOLEAN NOT NULL,
    for_series BOOLEAN NOT NULL,
    download_type VARCHAR(32) NOT NULL
);

COMMIT;