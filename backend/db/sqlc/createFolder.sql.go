// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: createFolder.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createFolder = `-- name: CreateFolder :one
insert into folder (folder_id, folder_name, subfolder_of, user_id, path, label, folder_description)
values ($1, $2, $3, $4, $5, $6, $7)
returning folder_id, user_id, folder_name, path, label, starred, created_at, updated_at, subfolder_of, deleted_at, folder_description, folder_password, ishidden
`

type CreateFolderParams struct {
	FolderID          string      `json:"folder_id"`
	FolderName        string      `json:"folder_name"`
	SubfolderOf       pgtype.Text `json:"subfolder_of"`
	UserID            string      `json:"user_id"`
	Path              string      `json:"path"`
	Label             string      `json:"label"`
	FolderDescription string      `json:"folder_description"`
}

func (q *Queries) CreateFolder(ctx context.Context, arg CreateFolderParams) (Folder, error) {
	row := q.db.QueryRow(ctx, createFolder,
		arg.FolderID,
		arg.FolderName,
		arg.SubfolderOf,
		arg.UserID,
		arg.Path,
		arg.Label,
		arg.FolderDescription,
	)
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
