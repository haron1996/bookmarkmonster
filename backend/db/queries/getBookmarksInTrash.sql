-- name: GetBookmarksInTrash :many
select * from bookmark where deleted is not null and user_id = $1 order by deleted desc;