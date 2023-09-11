-- name: GetFolder :one
select * from folder where folder_id = $1 and user_id = $2 and deleted_at is null limit 1;