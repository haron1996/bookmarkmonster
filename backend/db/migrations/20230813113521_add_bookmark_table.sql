-- +goose Up
DROP TABLE IF EXISTS bookmark;

CREATE TABLE bookmark (
    id text not null primary key,
    title text not null constraint bookmark_title_must_be_unique unique,
    bookmark text not null constraint bookmark_must_be_unique unique,
    host text not null,
    favicon text not null,
    thumbnail text not null,
    notes text,
    tag_id text not null references tag(id) on delete cascade,
    user_id text not null references userr(id) on delete cascade,
    added timestamptz not null default current_timestamp,
    updated timestamptz,
    deleted timestamptz
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS bookmark;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
