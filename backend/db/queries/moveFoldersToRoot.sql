-- name: MoveFoldersToRoot :many
update folder set path = subpath(path, nlevel((select path from folder where folder.label = $1))-1) where path <@ (
select path from folder where folder.label = $2
) returning *;