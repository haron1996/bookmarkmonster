-- name: GetChromeBookmark :one
select * from bookmark where bookmark = $1 and title = $2 and user_id = $3 and fromChrome = 'true' and deleted is null limit 1;