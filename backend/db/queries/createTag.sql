-- name: CreateTag :one
INSERT INTO tag (id, name, user_id)
VALUES ($1, $2, $3)
RETURNING *;