-- name: RenameTag :one
update tag set name = $1 where id = $2 and user_id = $3 and deleted is null returning *;