-- name: GetUserRootFolderByName :one
select exists (select * from folder where user_id=$1 and deleted_at is null and subfolder_of is null and folder_name = $2);