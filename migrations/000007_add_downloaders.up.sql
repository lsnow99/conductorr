BEGIN;

CREATE TABLE downloader(
    id INTEGER PRIMARY KEY,
    downloader_type VARCHAR(64) NOT NULL,
    config TEXT NOT NULL
);

COMMIT;