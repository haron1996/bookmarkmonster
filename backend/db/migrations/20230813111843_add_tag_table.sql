-- +goose Up
DROP TABLE IF EXISTS tag;

CREATE TABLE tag (
    id text not null primary key,
    name text not null,
    user_id text not null references userr(id) on delete cascade,
    added timestamptz not null default current_timestamp,
    updated timestamptz,
    deleted timestamptz,
    unique(user_id, name)
);
-- +goose StatementBegin
SELECT 'up SQL query';

create or replace function make_lower_name() 
  returns trigger
as
$$
begin
  new.name := lower(new.name);
  return new;
end;
$$
language plpgsql;

create trigger ensure_lower_name_trg
   before update or insert on tag
   for each row
   execute procedure make_lower_name();
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS tag;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
