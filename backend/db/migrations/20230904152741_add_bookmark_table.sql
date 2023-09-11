-- +goose Up
DROP TABLE IF EXISTS bookmark;

CREATE TABLE bookmark (
    id text not null primary key,
    title text not null,
    bookmark text not null,
    host text not null,
    favicon text,
    thumbnail text,
    notes text,
    user_id text not null references userr(id) on delete cascade,
    added timestamptz not null default current_timestamp,
    updated timestamptz,
    deleted timestamptz,
    folder_id text references folder(folder_id) on delete cascade,
    beautified timestamptz
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS bookmark;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
