// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: getUserRootFolders.sql

package db

import (
	"context"
)

const getUserRootFolders = `-- name: GetUserRootFolders :many
select folder_id, user_id, folder_name, path, label, starred, created_at, updated_at, subfolder_of, deleted_at, folder_description, folder_password, ishidden from folder where user_id = $1 AND subfolder_of is null and deleted_at is null order by created_at desc
`

func (q *Queries) GetUserRootFolders(ctx context.Context, userID string) ([]Folder, error) {
	rows, err := q.db.Query(ctx, getUserRootFolders, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Folder{}
	for rows.Next() {
		var i Folder
		if err := rows.Scan(
			&i.FolderID,
			&i.UserID,
			&i.FolderName,
			&i.Path,
			&i.Label,
			&i.Starred,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.SubfolderOf,
			&i.DeletedAt,
			&i.FolderDescription,
			&i.FolderPassword,
			&i.Ishidden,
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
