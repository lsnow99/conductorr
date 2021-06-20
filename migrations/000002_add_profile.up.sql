BEGIN;

CREATE TABLE profile(
    id INTEGER PRIMARY KEY,
    name VARCHAR(128),
    filter TEXT,
    sorter TEXT
);

COMMIT;