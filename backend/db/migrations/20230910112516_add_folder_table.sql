-- +goose Up
alter table folder add column if not exists folder_description text not null default '';
alter table folder add column if not exists folder_password text;
alter table folder add column if not exists isHidden boolean not null default 'false';
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

create or replace trigger ensure_lower_name_trg
   before update or insert on folder
   for each row
   execute procedure make_lower_name();
-- +goose StatementEnd

-- +goose Down
alter table folder drop column if exists folder_description;
alter table folder drop column if exists folder_password;
alter table folder drop column if exists isHidden;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
