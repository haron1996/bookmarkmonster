// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: getFolderByNameUserIdAndParentFolderId.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const getFolderByNameUserIdAndParentFolderId = `-- name: GetFolderByNameUserIdAndParentFolderId :one
select exists (select folder_id, user_id, folder_name, path, label, starred, created_at, updated_at, subfolder_of, deleted_at, folder_description, folder_password, ishidden from folder where user_id=$1 and deleted_at is null and subfolder_of = $2 and folder_name = $3)
`

type GetFolderByNameUserIdAndParentFolderIdParams struct {
	UserID      string      `json:"user_id"`
	SubfolderOf pgtype.Text `json:"subfolder_of"`
	FolderName  string      `json:"folder_name"`
}

func (q *Queries) GetFolderByNameUserIdAndParentFolderId(ctx context.Context, arg GetFolderByNameUserIdAndParentFolderIdParams) (bool, error) {
	row := q.db.QueryRow(ctx, getFolderByNameUserIdAndParentFolderId, arg.UserID, arg.SubfolderOf, arg.FolderName)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}