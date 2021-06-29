BEGIN;

CREATE TABLE downloader(
    id INTEGER PRIMARY KEY,
    name VARCHAR(128) NOT NULL,
    downloader_type VARCHAR(64) NOT NULL,
    config TEXT NOT NULL
);

COMMIT;