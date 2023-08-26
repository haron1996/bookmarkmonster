-- name: RenameBookmark :one
update bookmark set title = $1 where id = $2 returning *;