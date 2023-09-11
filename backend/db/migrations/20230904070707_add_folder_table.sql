-- +goose Up
CREATE EXTENSION IF NOT EXISTS ltree;
DROP TABLE IF EXISTS folder;

CREATE TABLE folder (
    folder_id text not null primary key,
    user_id text not null references userr (id) on delete cascade,
    folder_name text not null,
    path ltree not null,
    label text check (label ~* '^[A-Za-z0-9_]{1,255}$') not null unique,
    starred boolean not null default 'false',
    created_at timestamptz not null default current_timestamp,
    updated_at timestamptz not null default current_timestamp,
    subfolder_of text references folder (folder_id) on delete cascade,
    deleted_at timestamptz,
    unique(subfolder_of, folder_name)
);

-- hidden folders shoudln't appear during search.. its bookmarks should not appear during search

-- +goose StatementBegin
SELECT 'up SQL query';

create or replace function make_lower_name() 
  returns trigger
as
$$
begin
  new.folder_name := lower(new.folder_name);
  return new;
end;
$$
language plpgsql;

create trigger ensure_lower_name_trg
   before update or insert on folder
   for each row
   execute procedure make_lower_name();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
DROP TABLE IF EXISTS folder;