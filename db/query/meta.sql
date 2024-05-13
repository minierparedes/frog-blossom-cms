-- name: CreateMeta :one
INSERT INTO meta (
  page_id,
  posts_id,
  meta_title,
  meta_description,
  meta_robots,
  meta_og_image,
  locale,
  page_amount,
  site_language,
  meta_key,
  meta_value
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
) RETURNING *;

-- name: GetMeta :one
SELECT * FROM meta
WHERE id = $1 LIMIT 1;

-- name: ListMeta :many
SELECT * FROM meta
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateMeta :one
UPDATE meta
  SET page_id = $2,
posts_id = $3,
meta_title = $4,
meta_description = $5,
meta_robots = $6,
meta_og_image = $7,
locale = $8,
page_amount = $9,
site_language = $10,
meta_key = $11,
meta_value = $12
WHERE id = $1
RETURNING *;

-- name: DeleteMeta :exec
DELETE FROM meta
WHERE id = $1;
