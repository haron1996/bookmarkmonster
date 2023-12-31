// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: getUserBookmarksByTagID.sql

package db

import (
	"context"
)

const getUserBookmarksByTagID = `-- name: GetUserBookmarksByTagID :many
select bookmark_id, tag_id from bookmark_tag where tag_id = $1
`

func (q *Queries) GetUserBookmarksByTagID(ctx context.Context, tagID string) ([]BookmarkTag, error) {
	rows, err := q.db.Query(ctx, getUserBookmarksByTagID, tagID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []BookmarkTag{}
	for rows.Next() {
		var i BookmarkTag
		if err := rows.Scan(&i.BookmarkID, &i.TagID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
