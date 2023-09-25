-- +goose Up
drop table if exists pocket_token cascade;

create table pocket_token (
    access_token text not null,
    username text not null,
    user_id text not null unique references userr(id)
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
drop table if exists pocket_code cascade;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
