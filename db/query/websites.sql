-- name: CreateWebsites :one
INSERT INTO websites (
  name,
  domain,
  owner_id,
  password,
  template_id,
  builder_enabled
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetWebsites :one
SELECT * FROM websites
WHERE id = $1 LIMIT 1;

-- name: ListWebsites :many
SELECT * FROM websites
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateWebsites :one
UPDATE websites
  SET domain = $2,
  password = $3,
  builder_enabled = $4
WHERE id = $1
RETURNING *;

-- name: DeleteWebsites :exec
DELETE FROM websites
WHERE id = $1;
