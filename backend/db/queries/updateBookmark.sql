-- name: UpdateBookmark :one
update bookmark set title = $1, notes = $2 where id = $3 and user_id = $4 returning *;