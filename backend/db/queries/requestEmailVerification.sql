-- name: RequestEmailVerification :one
insert into email_verification(email, code, expiry, user_password)
values($1, $2, $3, $4)
on conflict (email) do update set code = excluded.code, expiry = excluded.expiry, user_password = excluded.user_password
returning *;