-- name: GetRootBookmarks :many
select * from bookmark where user_id = $1 and folder_id is null and deleted is null and beautified is null order by added desc;