-- name: GetTagByUserAndName :one
select * from tag where user_id = $1 and name = $2 limit 1;