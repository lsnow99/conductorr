BEGIN;

CREATE TABLE path(
    --sqlite--
    id INTEGER PRIMARY KEY,
    --postgresql--
    id SERIAL PRIMARY KEY,
    --end--
    path TEXT NOT NULL,
    movies_default BOOLEAN NOT NULL,
    series_default BOOLEAN NOT NULL
);

COMMIT;