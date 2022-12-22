--sqlite--
PRAGMA foreign_keys=off;
--END--

BEGIN;


--sqlite--
UPDATE downloader SET config = JSON_SET(config, '$.base_url', JSON_EXTRACT(config, '$.baseUrl'));
--END--


--postgresql--
UPDATE downloader SET config = JSONB_SET(config, 'base_url', JSONB_EXTRACT_PATH(config, 'baseUrl'));
--END--

COMMIT;

--sqlite--
PRAGMA foreign_keys=on;
--END--

