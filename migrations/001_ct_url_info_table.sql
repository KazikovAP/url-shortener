CREATE TABLE IF NOT EXISTS url_info (
    id          BIGSERIAL   PRIMARY KEY,
    url         TEXT        UNIQUE NOT NULL,
    alias       TEXT        NOT NULL,
    created_at  TIMESTAMP   WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);