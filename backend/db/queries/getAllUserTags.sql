-- name: GetAllUserTags :many
select * from tag where user_id = $1;