-- name: GetWaitList :many
select * from waitlist order by joined desc;