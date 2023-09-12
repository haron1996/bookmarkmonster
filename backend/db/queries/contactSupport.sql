-- name: ContactSupport :one
insert into support (id, email, title, comment)
values ($1, $2, $3, $4)
returning *;