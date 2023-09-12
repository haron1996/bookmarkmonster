-- +goose Up
drop table if exists support;

create table support (
    id text not null primary key,
    email text not null,
    title text not null,
    comment text not null,
    created timestamptz not null default current_timestamp
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
drop table if exists support;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
