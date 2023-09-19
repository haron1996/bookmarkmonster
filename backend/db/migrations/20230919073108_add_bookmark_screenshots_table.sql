-- +goose Up
drop table if exists bookmark_screenshots;

create table bookmark_screenshots (
    id text not null primary key,
    screenshot_location text not null,
    bookmark_id text not null references bookmark(id) on delete cascade,
    user_id text not null references userr(id) on delete cascade,
    added timestamptz not null default current_timestamp,
    deleted timestamptz,
    fullpage boolean not null
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
drop table if exists bookmark_screenshots;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
