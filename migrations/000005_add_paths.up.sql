BEGIN;

CREATE TABLE path(
    id INTEGER PRIMARY KEY,
    path TEXT NOT NULL,
    movies_default BOOLEAN NOT NULL,
    series_default BOOLEAN NOT NULL
);

COMMIT;