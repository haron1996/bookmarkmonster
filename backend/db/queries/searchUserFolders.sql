-- name: SearchUserFolders :many
SELECT * FROM folder
WHERE folder_name ~~* $1 AND user_id = $2 AND deleted_at IS NULL
ORDER BY (created_at) DESC;