-- name: UpdateUserLastLoginAndRefreshToken :one
UPDATE userr
SET last_login = $1, refresh_token = $2
WHERE id = $3
RETURNING *;