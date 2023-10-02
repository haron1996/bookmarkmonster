// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: updateParentFolderToNull.sql

package db

import (
	"context"
)

const updateParentFolderToNull = `-- name: UpdateParentFolderToNull :exec
update folder set subfolder_of = null where folder_id = $1
`

func (q *Queries) UpdateParentFolderToNull(ctx context.Context, folderID string) error {
	_, err := q.db.Exec(ctx, updateParentFolderToNull, folderID)
	return err
}
