-- name: CreateWebsite :one
INSERT INTO website (
  name,
  domain,
  owner_id,
  selected_template
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetWebsite :one
SELECT * FROM website
WHERE id = $1 LIMIT 1;

-- name: ListWebsite :many
SELECT * FROM website
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateWebsite :one
UPDATE website
  SET name = $2,
      domain = $3,
      owner_id = $4,
      selected_template = $5
WHERE id = $1
RETURNING *;

-- name: DeleteWebsite :exec
DELETE FROM website
WHERE id = $1;
