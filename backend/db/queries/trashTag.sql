-- name: TrashTag :one
update tag set deleted = now() where id = $1 and user_id = $2 and deleted is null returning *;