-- name: MoveBookmark :one
update bookmark set folder_id = $1 where id = $2 returning *;