-- name: GetAllUserTags :many
select * from tag where user_id = $1 AND deleted is null order by (added) desc;