-- name: UpdateFolderParent :exec
update folder
set subfolder_of = $1
where folder_id = $2;