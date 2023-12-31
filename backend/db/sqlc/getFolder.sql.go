// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: getFolder.sql

package db

import (
	"context"
)

const getFolder = `-- name: GetFolder :one
select folder_id, user_id, folder_name, path, label, starred, created_at, updated_at, subfolder_of, deleted_at, folder_description, folder_password, ishidden from folder where folder_id = $1 and user_id = $2 and deleted_at is null limit 1
`

type GetFolderParams struct {
	FolderID string `json:"folder_id"`
	UserID   string `json:"user_id"`
}

func (q *Queries) GetFolder(ctx context.Context, arg GetFolderParams) (Folder, error) {
	row := q.db.QueryRow(ctx, getFolder, arg.FolderID, arg.UserID)
	var i Folder
	err := row.Scan(
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
	)
	return i, err
}
