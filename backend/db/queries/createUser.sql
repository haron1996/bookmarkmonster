-- name: CreateUser :one
INSERT INTO userr (id, email, name, email_verified, picture, signup_mode, account_password)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;