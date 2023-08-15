-- name: GetBookmarkByID :one
select * from bookmark where id = $1 limit 1;