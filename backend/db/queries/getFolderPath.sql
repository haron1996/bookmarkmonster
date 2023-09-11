-- name: GetFolderPath :many
select * from folder where folder.path @> (select path from folder as f where f.label = $1) order by path;