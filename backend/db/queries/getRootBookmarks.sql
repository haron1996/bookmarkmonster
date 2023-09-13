-- name: GetUserRootBookmarks :many
select * from bookmark where user_id = $1 and folder_id is null and deleted is null order by (added) desc;