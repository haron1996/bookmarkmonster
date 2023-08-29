-- name: TrashBookmark :one
update bookmark set deleted = now() where id = $1 returning *;