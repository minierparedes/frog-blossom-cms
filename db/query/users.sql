-- name: CreateUsers :one
INSERT INTO Users (
  username,
  email,
  password,
  role,
  first_name,
  last_name,
  user_url,
  description,
  updated_at,
  is_deleted
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
) RETURNING *;

-- name: GetUsers :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateUsers :one
UPDATE users
  SET username = $2,
  email = $3,
  password = $4,
  role = $5,
  first_name = $6,
  last_name = $7,
  user_url = $8,
  description = $9,
  updated_at = $10,
  is_deleted = $11
WHERE id = $1
RETURNING *;

-- name: DeleteUsers :exec
DELETE FROM users
WHERE id = $1;
