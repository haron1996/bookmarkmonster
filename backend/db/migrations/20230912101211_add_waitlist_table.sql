-- +goose Up
drop table if exists waitlist;

create table waitlist (
    id text not null primary key,
    email text not null unique,
    name text not null,
    company_name text,
    comment text,
    joined timestamptz not null default current_timestamp
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
drop table if exists waitlist;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
