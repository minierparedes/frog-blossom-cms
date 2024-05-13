-- name: CreatePosts :one
INSERT INTO posts (
  title,
  content,
  author_id,
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
  $1, $2, $3, DEFAULT, $4, $5, $6, $7, $8, $9, $10, $11
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
  updated_at = $5,
  status = $6,
  published_at = $7,
  edited_at = $8,
  post_author = $9,
  post_mime_type = $10,
  published_by = $11,
  updated_by = $12
WHERE id = $1
RETURNING *;

-- name: DeletePosts :exec
DELETE FROM posts
WHERE id = $1;
