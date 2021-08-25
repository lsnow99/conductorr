BEGIN;

CREATE TABLE download (
    --sqlite--
    id INTEGER PRIMARY KEY,
    --postgresql--
    id SERIAL PRIMARY KEY,
    --end
    media_id INTEGER REFERENCES media(id) ON DELETE SET NULL,
    downloader_id INTEGER REFERENCES downloader(id) ON DELETE SET NULL,
    status VARCHAR(64) NOT NULL,
    friendly_name TEXT NOT NULL,
    identifier TEXT NOT NULL
);

COMMIT;