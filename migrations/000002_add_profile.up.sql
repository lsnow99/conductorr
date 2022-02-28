BEGIN;

CREATE TABLE profile(
    --sqlite--
    id INTEGER PRIMARY KEY,
    --postgresql--
    id SERIAL PRIMARY KEY,
    --end--
    name VARCHAR(128),
    filter TEXT,
    sorter TEXT
);

COMMIT;