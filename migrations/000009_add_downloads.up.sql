BEGIN;

CREATE TABLE download (
    id INTEGER PRIMARY KEY,
    media_id INTEGER REFERENCES media(id) ON DELETE SET NULL,
    downloader_id INTEGER REFERENCES downloader(id) ON DELETE SET NULL,
    status VARCHAR(64) NOT NULL,
    friendly_name TEXT NOT NULL,
    identifier TEXT NOT NULL
);

COMMIT;