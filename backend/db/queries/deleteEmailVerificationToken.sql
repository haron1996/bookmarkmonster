-- name: DeleteEmailVerificationToken :exec
delete from email_verification where email = $1;