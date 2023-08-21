-- name: GetUsers :many
select * from userr order by last_login desc;