-- name: GetBookmarkByID :one
select * from bookmark where id = $1 and deleted is null limit 1;