-- name: CreateBookmark :one
INSERT INTO bookmark (id, title, bookmark, host, user_id, notes, folder_id)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;