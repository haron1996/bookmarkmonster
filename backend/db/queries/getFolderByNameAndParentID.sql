-- name: GetFolderByNameAndParentID :one
select exists(select * from folder where folder_name = $1 and subfolder_of = $2 limit 1);