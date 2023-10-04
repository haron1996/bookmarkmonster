-- +goose Up
alter table bookmark add column if not exists fromChrome boolean not null default 'false';
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
alter table bookmark drop column if exists fromChrome;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
