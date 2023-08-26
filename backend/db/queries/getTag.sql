-- name: GetTag :one
select * from tag where id = $1 limit 1;