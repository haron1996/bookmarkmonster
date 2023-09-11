-- name: GetEmailConfirmationLink :one
select * from email_verification where email = $1;