BEGIN;

CREATE TABLE system_data (
    --sqlite--
    id INTEGER PRIMARY KEY,
    --postgresql--
    id SERIAL PRIMARY KEY,
    --end--
    jwt_secret VARCHAR(128)
);

INSERT INTO system_data (id) VALUES (1);

COMMIT;