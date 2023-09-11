-- name: CheckIfEmailIsInUse :one
select exists(select * from userr where email = $1);