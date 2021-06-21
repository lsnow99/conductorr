BEGIN;

CREATE TABLE genre (
    id INTEGER PRIMARY KEY,
    name VARCHAR(256) NOT NULL UNIQUE
);

CREATE TABLE media_genre (
    id INTEGER PRIMARY KEY,
    media_id INTEGER NOT NULL,
    genre_id INTEGER NOT NULL,
    FOREIGN KEY(media_id) REFERENCES media(id),
    FOREIGN KEY(genre_id) REFERENCES genre(id) 
);

COMMIT;