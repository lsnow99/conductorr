BEGIN;

CREATE TABLE media_server(
    id INTEGER PRIMARY KEY,
    name VARCHAR(128) NOT NULL,
    media_server_type VARCHAR(64) NOT NULL,
    config TEXT NOT NULL
);

COMMIT;