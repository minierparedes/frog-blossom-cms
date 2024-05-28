-- name: CreatePages :one
INSERT INTO pages (
  domain,
  author_id,
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
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
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
  author_id = $3,
  page_author = $4,
  title = $5,
  url = $6,
  menu_order = $7,
  component_type = $8,
  component_value = $9,
  page_identifier = $10,
  option_id = $11,
  option_name = $12,
  option_value = $13,
  option_required = $14
WHERE id = $1
RETURNING *;

-- name: DeletePages :exec
DELETE FROM pages
WHERE id = $1;
