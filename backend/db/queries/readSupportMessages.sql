-- name: ReadSupportMessages :many
select * from support order by created desc;