-- name: CreateFolder :one
insert into folder (folder_id, folder_name, subfolder_of, user_id, path, label, folder_description)
values ($1, $2, $3, $4, $5, $6, $7)
returning *;