CREATE TABLE IF NOT EXISTS albums (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    artist VARCHAR(255) NOT NULL,
    cover VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS tracks (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    album_id INTEGER REFERENCES albums(id),
    file_name VARCHAR(255) NOT NULL
);

\set app_user `echo "$DB_USER"`
\set app_password `echo "$DB_PASSWORD"`

CREATE USER :app_user WITH PASSWORD :'app_password';
GRANT SELECT, INSERT, UPDATE ON ALL TABLES IN SCHEMA public TO :app_user;