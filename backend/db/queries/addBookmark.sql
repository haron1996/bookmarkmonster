-- name: AddBookmark :one
INSERT INTO bookmark (id, title, bookmark, host, favicon, thumbnail, user_id)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;