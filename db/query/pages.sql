-- name: CreatePages :one
INSERT INTO pages (
  domain,
  page_author,
  title,
  url,
  menu_order,
  component_type,
  component_value,
  page_identifier,
  option_id,
  option_name,
  option_value,
  option_required
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
) RETURNING *;

-- name: GetPages :one
SELECT * FROM pages
WHERE id = $1 LIMIT 1;

-- name: ListPages :many
SELECT * FROM pages
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdatePages :one
UPDATE pages
  SET domain = $2,
  page_author = $3,
  title = $4,
  url = $5,
  menu_order = $6,
  component_type = $7,
  component_value = $8,
  page_identifier = $9,
  option_id = $10,
  option_name = $11,
  option_value = $12,
  option_required = $13
WHERE id = $1
RETURNING *;

-- name: DeletePages :exec
DELETE FROM pages
WHERE id = $1;
