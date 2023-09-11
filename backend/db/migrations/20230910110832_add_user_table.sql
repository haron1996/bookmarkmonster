-- +goose Up
create type signup_mode as enum ('email', 'google', 'facebook', 'github', 'reddit');
alter table userr add column signup_mode signup_mode not null default 'email';
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
alter table userr drop column signup_mode;
drop type if exists signup_mode;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd