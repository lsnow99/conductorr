BEGIN;

CREATE TABLE media_server(
    --sqlite--
    id INTEGER PRIMARY KEY,
    --postgresql--
    id SERIAL PRIMARY KEY,
    --end
    name VARCHAR(128) NOT NULL,
    media_server_type VARCHAR(64) NOT NULL,
    config TEXT NOT NULL
);

COMMIT;