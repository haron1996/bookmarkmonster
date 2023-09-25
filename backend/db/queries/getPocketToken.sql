-- name: GetPocketToken :one
select * from pocket_token where user_id = $1;