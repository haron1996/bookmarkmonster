-- name: GetPocketCode :one
select * from pocket_code where user_id = $1;