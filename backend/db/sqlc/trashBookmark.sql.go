// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: trashBookmark.sql

package db

import (
	"context"
)

const trashBookmark = `-- name: TrashBookmark :one
update bookmark set deleted = now() where id = $1 returning id, title, bookmark, host, favicon, thumbnail, notes, user_id, added, updated, deleted, folder_id, beautified
`

func (q *Queries) TrashBookmark(ctx context.Context, id string) (Bookmark, error) {
	row := q.db.QueryRow(ctx, trashBookmark, id)
	var i Bookmark
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Bookmark,
		&i.Host,
		&i.Favicon,
		&i.Thumbnail,
		&i.Notes,
		&i.UserID,
		&i.Added,
		&i.Updated,
		&i.Deleted,
		&i.FolderID,
		&i.Beautified,
	)
	return i, err
}
