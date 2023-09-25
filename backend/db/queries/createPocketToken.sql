-- name: CreatePocketToken :one
insert into pocket_token (access_token, username, user_id)
values ($1, $2, $3)
on conflict (user_id) do update set access_token = excluded.access_token, username = excluded.username
returning *;