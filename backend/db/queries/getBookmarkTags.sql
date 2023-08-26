-- name: GetBookmarkTags :many
select tag_id from bookmark_tag where bookmark_id = $1;