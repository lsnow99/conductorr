
--sqlite--
PRAGMA foreign_keys=off;
--END--

BEGIN;


--sqlite--
UPDATE downloader SET config = JSON_SET(config, '$.baseUrl', JSON_EXTRACT(config, '$.base_url'));
--END--


--postgresql--
UPDATE downloader SET config = JSONB_SET(config, 'baseUrl', JSONB_EXTRACT_PATH(config, 'base_url'));
--END--

COMMIT;

--sqlite--
PRAGMA foreign_keys=on;
--END--

