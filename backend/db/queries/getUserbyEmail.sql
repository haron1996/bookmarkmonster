-- name: GetUserByEmailAndID :one
SELECT * FROM userr
WHERE email = $1 AND id = $2
LIMIT 1;