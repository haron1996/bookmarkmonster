-- name: CreatePocketCode :one
insert into pocket_code (user_id, code)
values ($1, $2)
on conflict(user_id) do update set code = excluded.code
returning *;