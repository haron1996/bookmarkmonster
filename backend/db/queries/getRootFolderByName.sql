-- name: GetRootFolderByName :one
select exists(select * from folder where folder_name = $1 and subfolder_of is null limit 1);