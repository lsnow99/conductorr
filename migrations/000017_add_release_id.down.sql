BEGIN;

ALTER TABLE download DROP COLUMN release_id;

COMMIT;
