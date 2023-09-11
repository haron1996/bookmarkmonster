-- name: GetUserRootFolders :many
select * from folder where user_id = $1 AND subfolder_of is null and deleted_at is null order by created_at desc;