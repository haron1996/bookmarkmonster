-- name: UpdateEmailVericationToken :exec
update email_verification set expiry = $1 where email = $2;