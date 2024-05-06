-- name: CreateTemplate :one
INSERT INTO template (
    name
) VALUES (
    $1
) RETURNING *;

-- name: GetTemplate :one
SELECT * FROM template
WHERE id = $1 LIMIT 1;

-- name: ListTemplate :many
SELECT * FROM template
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateTemplate :one
UPDATE template
SET name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteTemplate :exec
DELETE FROM template
WHERE id = $1;
