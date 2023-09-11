-- name: UpdateFolder :one
update folder set folder_name = $1, folder_description = $2 where folder_id = $3 and user_id = $4 returning *;