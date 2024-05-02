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
