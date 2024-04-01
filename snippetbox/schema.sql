DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id              INTEGER  PRIMARY KEY AUTOINCREMENT,
    name            TEXT     NOT NULL,
    email           TEXT     NOT NULL UNIQUE,
    hashed_password BLOB     NOT NULL,
    created         DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX index__users__name ON users(name);
CREATE INDEX index__users__email ON users(email);


DROP TABLE IF EXISTS sessions;
CREATE TABLE sessions (
    token   TEXT    PRIMARY KEY,
    data    BLOB    NOT NULL,
    expiry  REAL    NOT NULL
);
CREATE INDEX index__sessions__expiry ON sessions(expiry);


DROP TABLE IF EXISTS snippets;
CREATE TABLE snippets (
    id      INTEGER  PRIMARY KEY AUTOINCREMENT,
    title   TEXT     NOT NULL,
    content TEXT     NOT NULL,
    expires DATETIME NOT NULL,
    created DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX index__snippets__created ON snippets(created);
CREATE INDEX index__snippets__expires ON snippets(expires);


INSERT INTO snippets (title, created, expires, content) VALUES
(
    'An old silent pond', DATETIME(), DATETIME('NOW', '+1 YEAR'),
    'An old silent pond...
A frog jumps into the pond,
splash! Silence again.

– Matsuo Bashō'
),
(
    'Over the wintry forest', DATETIME(), DATETIME('NOW', '+1 YEAR'),
    'Over the wintry
forest, winds howl in rage
with no leaves to blow.

– Natsume Soseki'
),
(
    'First autumn morning', DATETIME(), DATETIME('NOW', '+1 DAY'),
    'First autumn morning
the mirror I stare into
shows my father''s face.

– Murakami Kijo'
);

