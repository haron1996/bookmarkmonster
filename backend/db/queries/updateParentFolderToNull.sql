-- name: UpdateParentFolderToNull :exec
update folder set subfolder_of = null where folder_id = $1;