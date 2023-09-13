// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: getFolderByNameAndParentID.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const getFolderByNameAndParentID = `-- name: GetFolderByNameAndParentID :one
select exists(select folder_id, user_id, folder_name, path, label, starred, created_at, updated_at, subfolder_of, deleted_at, folder_description, folder_password, ishidden from folder where folder_name = $1 and subfolder_of = $2 limit 1)
`

type GetFolderByNameAndParentIDParams struct {
	FolderName  string      `json:"folder_name"`
	SubfolderOf pgtype.Text `json:"subfolder_of"`
}

func (q *Queries) GetFolderByNameAndParentID(ctx context.Context, arg GetFolderByNameAndParentIDParams) (bool, error) {
	row := q.db.QueryRow(ctx, getFolderByNameAndParentID, arg.FolderName, arg.SubfolderOf)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}