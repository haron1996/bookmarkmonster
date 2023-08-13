-- name: CreateUser :one
INSERT INTO userr (id, email, name, email_verified, picture)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;