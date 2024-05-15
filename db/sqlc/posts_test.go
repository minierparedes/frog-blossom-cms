package frog_blossom_db

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomPosts(t *testing.T) Post {
	// Arrange
	now := time.Now().UTC()
	randomUser := createRandomUser(t)

	args := CreatePostsParams{
		Title:        "Lorem ipsum dolor sit amet",
		Content:      "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
		AuthorID:     randomUser.ID,
		Url:          "https://example.com",
		UpdatedAt:    now,
		Status:       "admin",
		PublishedAt:  now,
		EditedAt:     now,
		PostAuthor:   randomUser.Username,
		PostMimeType: "text/plain",
		PublishedBy:  randomUser.Username,
		UpdatedBy:    randomUser.Username,
	}

	// Act
	posts, err := testQueries.CreatePosts(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, posts)

	// Assert
	require.Equal(t, args.Title, posts.Title)
	require.Equal(t, args.Content, posts.Content)
	require.Equal(t, args.AuthorID, posts.AuthorID)
	require.Equal(t, args.Url, posts.Url)
	require.Equal(t, args.UpdatedAt, posts.UpdatedAt.UTC())
	require.Equal(t, args.Status, posts.Status)
	require.Equal(t, args.PublishedAt, posts.PublishedAt.UTC())
	require.Equal(t, args.EditedAt, posts.EditedAt.UTC())
	require.Equal(t, args.PostAuthor, posts.PostAuthor)
	require.Equal(t, args.PostMimeType, posts.PostMimeType)
	require.Equal(t, args.PublishedBy, posts.PublishedBy)
	require.Equal(t, args.UpdatedBy, posts.UpdatedBy)

	return posts
}

func TestCreatePosts(t *testing.T) {
	createRandomPosts(t)
}
