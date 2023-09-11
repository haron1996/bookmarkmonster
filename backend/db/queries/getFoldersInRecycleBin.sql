-- name: GetFoldersInRecycleBin :many
select * from folder where deleted_at is not null and user_id = $1 order by deleted_at desc;