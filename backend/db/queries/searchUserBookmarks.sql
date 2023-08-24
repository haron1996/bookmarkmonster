-- name: SearchUserBookmarks :many
SELECT * FROM bookmark
WHERE title ~~* $1 AND user_id = $2 AND deleted IS NULL;