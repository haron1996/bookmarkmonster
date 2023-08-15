-- name: GetUserBookmarks :many
select * from bookmark where user_id = $1 order by (added) desc;