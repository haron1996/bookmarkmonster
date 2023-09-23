-- +goose Up
drop trigger if exists ensure_lower_name_trg on tag cascade;
-- +goose StatementBegin
SELECT 'up SQL query';

create or replace function make_tag_name_lower() 
  returns trigger
as
$$
begin
  new.name := lower(new.name);
  return new;
end;
$$
language plpgsql;

create or replace trigger ensure_lower_tag_name_trg
   before update or insert on tag
   for each row
   execute procedure make_tag_name_lower();
-- +goose StatementEnd

-- +goose Down
drop trigger if exists ensure_lower_tag_name_trg on tag cascade;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
