-- name: GetFolderBookmarks :many
select * from bookmark where user_id = $1 and folder_id = $2 and deleted is null and beautified is null order by added desc;