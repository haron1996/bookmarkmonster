-- +goose Up
DROP TABLE IF EXISTS userr cascade;

CREATE TABLE userr (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL CONSTRAINT email_must_be_unique UNIQUE,
    email_verified BOOLEAN NOT NULL DEFAULT 'false',
    picture TEXT,
    account_password TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    last_login TIMESTAMPTZ,
    refresh_token TEXT,
    deleted timestamptz
);

CREATE INDEX email_idx ON userr(LOWER(email));
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS userr cascade;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd