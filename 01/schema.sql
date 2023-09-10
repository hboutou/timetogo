DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    hashed_password BLOB NOT NULL,
    created DATETIME NOT NULL
);


DROP TABLE IF EXISTS sessions;
CREATE TABLE sessions (
    token TEXT PRIMARY KEY,
    data BLOB NOT NULL,
    expiry DATETIME NOT NULL
);

DROP TABLE IF EXISTS snippets;
CREATE TABLE snippets (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created DATETIME NOT NULL,
    expires DATETIME NOT NULL
);

CREATE INDEX idx_snippets_created ON snippets(created);

INSERT INTO snippets (title, content, created, expires) VALUES (
    'An old silent pond',
    'An old silent pond...
A frog jumps into the pond,
splash! Silence again.

– Matsuo Bashō',
    DATETIME(),
    DATETIME('NOW', '+1 YEAR')
);
INSERT INTO snippets (title, content, created, expires) VALUES (
    'Over the wintry forest',
    'Over the wintry
forest, winds howl in rage
with no leaves to blow.

– Natsume Soseki',
    DATETIME(),
    DATETIME('NOW', '+1 YEAR')
);
INSERT INTO snippets (title, content, created, expires) VALUES (
    'First autumn morning',
    'First autumn morning
the mirror I stare into
shows my father''s face.

– Murakami Kijo',
    DATETIME(),
    DATETIME('NOW', '+1 DAY')
);

