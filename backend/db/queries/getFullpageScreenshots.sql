-- name: GetFullpageScreenshots :many
select * from bookmark_screenshots where user_id = $1 and fullpage = 'true' and deleted is null order by added desc;