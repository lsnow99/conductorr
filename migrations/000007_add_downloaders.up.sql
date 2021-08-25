BEGIN;

CREATE TABLE downloader(
    --sqlite--
    id INTEGER PRIMARY KEY,
    --postgresql--
    id SERIAL PRIMARY KEY,
    --end
    name VARCHAR(128) NOT NULL,
    downloader_type VARCHAR(64) NOT NULL,
    config TEXT NOT NULL
);

COMMIT;