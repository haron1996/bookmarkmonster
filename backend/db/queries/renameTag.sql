-- name: RenameTag :one
update tag set name = $1 where id = $2 returning *;