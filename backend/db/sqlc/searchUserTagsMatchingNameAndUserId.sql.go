// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: searchUserTagsMatchingNameAndUserId.sql

package db

import (
	"context"
)

const searchTagsContainingTagNameAndUserID = `-- name: SearchTagsContainingTagNameAndUserID :many
SELECT id, name, user_id, added, updated, deleted FROM tag
WHERE name ~~* $1 AND user_id = $2 AND deleted IS NULL
ORDER BY (added) DESC
`

type SearchTagsContainingTagNameAndUserIDParams struct {
	Name   string `json:"name"`
	UserID string `json:"user_id"`
}

func (q *Queries) SearchTagsContainingTagNameAndUserID(ctx context.Context, arg SearchTagsContainingTagNameAndUserIDParams) ([]Tag, error) {
	rows, err := q.db.Query(ctx, searchTagsContainingTagNameAndUserID, arg.Name, arg.UserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Tag{}
	for rows.Next() {
		var i Tag
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.UserID,
			&i.Added,
			&i.Updated,
			&i.Deleted,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
