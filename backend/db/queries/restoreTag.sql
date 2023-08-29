-- name: RestoreTag :one
update tag set deleted = null where id = $1 returning *;