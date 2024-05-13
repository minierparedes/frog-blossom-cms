// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: posts.sql

package frog_blossom_db

import (
	"context"
	"time"
)

const createPosts = `-- name: CreatePosts :one
INSERT INTO posts (
  title,
  content,
  author,
  created_at,
  updated_at,
  status,
  published_at,
  edited_at,
  post_author,
  post_mime_type
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
) RETURNING id, title, content, author, created_at, updated_at, status, published_at, edited_at, post_author, post_mime_type
`

type CreatePostsParams struct {
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	Author       string    `json:"author"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Status       string    `json:"status"`
	PublishedAt  time.Time `json:"published_at"`
	EditedAt     time.Time `json:"edited_at"`
	PostAuthor   int64     `json:"post_author"`
	PostMimeType string    `json:"post_mime_type"`
}

func (q *Queries) CreatePosts(ctx context.Context, arg CreatePostsParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, createPosts,
		arg.Title,
		arg.Content,
		arg.Author,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Status,
		arg.PublishedAt,
		arg.EditedAt,
		arg.PostAuthor,
		arg.PostMimeType,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.Author,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Status,
		&i.PublishedAt,
		&i.EditedAt,
		&i.PostAuthor,
		&i.PostMimeType,
	)
	return i, err
}

const deletePosts = `-- name: DeletePosts :exec
DELETE FROM posts
WHERE id = $1
`

func (q *Queries) DeletePosts(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deletePosts, id)
	return err
}

const getPosts = `-- name: GetPosts :one
SELECT id, title, content, author, created_at, updated_at, status, published_at, edited_at, post_author, post_mime_type FROM posts
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetPosts(ctx context.Context, id int64) (Post, error) {
	row := q.db.QueryRowContext(ctx, getPosts, id)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.Author,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Status,
		&i.PublishedAt,
		&i.EditedAt,
		&i.PostAuthor,
		&i.PostMimeType,
	)
	return i, err
}

const listPosts = `-- name: ListPosts :many
SELECT id, title, content, author, created_at, updated_at, status, published_at, edited_at, post_author, post_mime_type FROM posts
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListPostsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListPosts(ctx context.Context, arg ListPostsParams) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, listPosts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Content,
			&i.Author,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Status,
			&i.PublishedAt,
			&i.EditedAt,
			&i.PostAuthor,
			&i.PostMimeType,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updatePosts = `-- name: UpdatePosts :one
UPDATE posts
  SET title = $2,
  content = $3,
  author = $4,
  created_at = $5,
  updated_at = $6,
  status = $7,
  published_at = $8,
  edited_at = $9,
  post_author = $10,
  post_mime_type = $11
WHERE id = $1
RETURNING id, title, content, author, created_at, updated_at, status, published_at, edited_at, post_author, post_mime_type
`

type UpdatePostsParams struct {
	ID           int64     `json:"id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	Author       string    `json:"author"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Status       string    `json:"status"`
	PublishedAt  time.Time `json:"published_at"`
	EditedAt     time.Time `json:"edited_at"`
	PostAuthor   int64     `json:"post_author"`
	PostMimeType string    `json:"post_mime_type"`
}

func (q *Queries) UpdatePosts(ctx context.Context, arg UpdatePostsParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, updatePosts,
		arg.ID,
		arg.Title,
		arg.Content,
		arg.Author,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Status,
		arg.PublishedAt,
		arg.EditedAt,
		arg.PostAuthor,
		arg.PostMimeType,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.Author,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Status,
		&i.PublishedAt,
		&i.EditedAt,
		&i.PostAuthor,
		&i.PostMimeType,
	)
	return i, err
}