BEGIN;

ALTER TABLE media ADD COLUMN tmdb_id INTEGER;

COMMIT;
