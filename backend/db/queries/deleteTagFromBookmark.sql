-- name: DeleteTagFromBookmark :exec
delete from bookmark_tag where bookmark_id = $1 and tag_id = $2;