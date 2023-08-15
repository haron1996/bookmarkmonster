-- name: GetUserBookmarksByTagID :many
select * from bookmark_tag where tag_id = $1;