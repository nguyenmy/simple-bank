-- name: CreateTransfer :one
INSERT INTO transfers(from_account_id,to_account_id,amount) VALUES($1,$2,$3) RETURNING *;

-- name: GetTransfer :one
SELECT * FROM transfers where id=$1 LIMIT 1;

-- name: ListTransfers :many
SELECT * FROM transfers order by id LIMIT $1 OFFSET $2;

-- name: UpdateTransfer :exec
UPDATE transfers set amount=$2 WHERE id=$1;

-- name: DeleteTransfer :exec
DELETE FROM transfers WHERE id=$1;