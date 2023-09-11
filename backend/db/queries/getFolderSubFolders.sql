-- name: GetFolderSubFolders :many
select * from folder where user_id = $1 AND subfolder_of = $2 and deleted_at is null order by created_at desc;