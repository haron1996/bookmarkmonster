-- name: JoinWaitList :one
insert into waitlist (id, email, name, company_name, comment)
values ($1, $2, $3, $4, $5)
on conflict (email) do update set id = excluded.id, name = excluded.name, company_name = excluded.company_name, comment = excluded.comment
returning *;