// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: getWaitList.sql

package db

import (
	"context"
)

const getWaitList = `-- name: GetWaitList :many
select id, email, name, company_name, comment, joined from waitlist order by joined desc
`

func (q *Queries) GetWaitList(ctx context.Context) ([]Waitlist, error) {
	rows, err := q.db.Query(ctx, getWaitList)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Waitlist{}
	for rows.Next() {
		var i Waitlist
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.Name,
			&i.CompanyName,
			&i.Comment,
			&i.Joined,
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