-- +goose Up
drop table if exists email_verification;

create table email_verification (
    email text not null unique,
    code text not null,
    expiry timestamptz not null,
    user_password text not null
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
drop table email_verification;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
