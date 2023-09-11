-- name: GetUserByEmail :one
select * from userr where email = $1;