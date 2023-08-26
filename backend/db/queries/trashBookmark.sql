-- name: TrashBookmark :one
update bookmark set deleted = $1 where id = $2 returning *;