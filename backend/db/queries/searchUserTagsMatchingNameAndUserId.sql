-- name: SearchTagsContainingTagNameAndUserID :many
SELECT * FROM tag
WHERE name ~~* $1 AND user_id = $2 AND deleted IS NULL
ORDER BY (added) DESC;