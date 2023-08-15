// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: getAllUserTags.sql

package db

import (
	"context"
)

const getAllUserTags = `-- name: GetAllUserTags :many
select id, name, user_id, added, updated, deleted from tag where user_id = $1
`

func (q *Queries) GetAllUserTags(ctx context.Context, userID string) ([]Tag, error) {
	rows, err := q.db.Query(ctx, getAllUserTags, userID)
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
