-- name: CreateBookmarkScreenshot :one
insert into bookmark_screenshots(id, screenshot_location, bookmark_id, user_id, fullpage)
values ($1, $2, $3, $4, $5)
returning *;