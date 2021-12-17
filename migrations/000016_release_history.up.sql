BEGIN;

CREATE TABLE release_history (
    id TEXT PRIMARY KEY,
    media_id INTEGER NOT NULL REFERENCES media(id) ON DELETE SET NULL DEFERRABLE INITIALLY DEFERRED,
    --postgresql--
    last_attempt DATETIME NOT NULL DEFAULT NOW()
    --sqlite--
    last_attempt DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
    --end--
);

COMMIT;