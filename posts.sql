-- name: CreatePosts :one
INSERT INTO posts (
  title,
  content,
  author_id,
  url,
  created_at,
  updated_at,
  status,
  published_at,
  edited_at,
  post_author,
  post_mime_type,
  published_by,
  updated_by
) VALUES (
  $1, $2, $3, $4, DEFAULT, $5, $6, $7, $8, $9, $10, $11, $12
) RETURNING *;


-- name: GetPosts :one
SELECT * FROM posts
WHERE id = $1 LIMIT 1;

-- name: ListPosts :many
SELECT * FROM posts
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdatePosts :one
UPDATE posts
  SET title = $2,
  content = $3,
  author_id = $4,
  url = $5,
  updated_at = $6,
  status = $7,
  published_at = $8,
  edited_at = $9,
  post_author = $10,
  post_mime_type = $11,
  published_by = $12,
  updated_by = $13
WHERE id = $1
RETURNING *;

-- name: DeletePosts :exec
DELETE FROM posts
WHERE id = $1;
