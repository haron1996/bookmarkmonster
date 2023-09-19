-- name: GetAboveFoldScreenshots :many
select * from bookmark_screenshots where user_id = $1 and fullpage = 'false' and deleted is null order by added desc;