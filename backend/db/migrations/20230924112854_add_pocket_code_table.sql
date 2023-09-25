-- +goose Up
drop table if exists pocket_code cascade;

create table pocket_code(
    user_id text not null unique references userr(id),
    code text not null
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
drop table if exists pocket_code cascade;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
