BEGIN;

ALTER TABLE downloader DROP COLUMN file_action;

COMMIT;
