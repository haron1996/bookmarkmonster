-- name: GetBookmarkByIdAndUserId :one
select * from bookmark where id =$1 and user_id = $2 limit 1;