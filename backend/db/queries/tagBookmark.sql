-- name: TagBookmark :one
INSERT INTO bookmark_tag (bookmark_id, tag_id)
VALUES ($1, $2)
RETURNING *;