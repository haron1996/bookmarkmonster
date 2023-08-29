-- name: TrashTag :one
update tag set deleted = now() where id = $1 returning *;