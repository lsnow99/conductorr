BEGIN;

CREATE TABLE download (
    id INTEGER PRIMARY KEY,
    media_id INTEGER NOT NULL REFERENCES media(id),
    downloader_id INTEGER NOT NULL REFERENCES downloader(id),
    status VARCHAR(64) NOT NULL,
    friendly_name TEXT NOT NULL,
    identifier TEXT NOT NULL
);

COMMIT;