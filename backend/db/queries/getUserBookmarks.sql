-- name: GetUserBookmarks :many
select * from bookmark where user_id = $1 and deleted is null order by (added) desc;