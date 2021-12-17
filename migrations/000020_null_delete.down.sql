--sqlite--
PRAGMA foreign_keys=off;

BEGIN;

ALTER TABLE release_history RENAME TO release_history_old;

CREATE TABLE release_history (
    id TEXT PRIMARY KEY,
    media_id INTEGER NOT NULL REFERENCES media(id) DEFERRABLE INITIALLY DEFERRED,
    last_attempt DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO release_history SELECT * FROM release_history_old;

DROP TABLE release_history_old;

COMMIT;

PRAGMA foreign_keys=on;
--end--