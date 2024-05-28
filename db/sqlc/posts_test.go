package frog_blossom_db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
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
	require.WithinDuration(t, args.UpdatedAt, posts.UpdatedAt.UTC(), time.Millisecond)
	require.Equal(t, args.Status, posts.Status)
	require.WithinDuration(t, args.PublishedAt, posts.PublishedAt.UTC(), time.Millisecond)
	require.WithinDuration(t, args.EditedAt, posts.EditedAt.UTC(), time.Millisecond)
	require.Equal(t, args.PostAuthor, posts.PostAuthor)
	require.Equal(t, args.PostMimeType, posts.PostMimeType)
	require.Equal(t, args.PublishedBy, posts.PublishedBy)
	require.Equal(t, args.UpdatedBy, posts.UpdatedBy)

	return posts
}

func TestCreatePosts(t *testing.T) {
	createRandomPosts(t)
}

func TestGetPosts(t *testing.T) {
	// Arrange
	randomPosts := createRandomPosts(t)

	// Act
	posts, err := testQueries.GetPosts(context.Background(), randomPosts.ID)
	require.NoError(t, err)
	require.NotEmpty(t, posts)

	// Assert
	require.Equal(t, randomPosts.ID, posts.ID)
	require.Equal(t, randomPosts.Title, posts.Title)
	require.Equal(t, randomPosts.Content, posts.Content)
	require.Equal(t, randomPosts.AuthorID, posts.AuthorID)
	require.Equal(t, randomPosts.Url, posts.Url)
	require.WithinDuration(t, randomPosts.UpdatedAt, posts.UpdatedAt, time.Millisecond)
	require.Equal(t, randomPosts.Status, posts.Status)
	require.WithinDuration(t, randomPosts.PublishedAt, posts.PublishedAt, time.Millisecond)
	require.WithinDuration(t, randomPosts.EditedAt, posts.EditedAt, time.Millisecond)
	require.Equal(t, randomPosts.PostAuthor, posts.PostAuthor)
	require.Equal(t, randomPosts.PostMimeType, posts.PostMimeType)
	require.Equal(t, randomPosts.PublishedBy, posts.PublishedBy)
	require.Equal(t, randomPosts.UpdatedBy, posts.UpdatedBy)
}

func TestUpdatePosts(t *testing.T) {
	// Arrange
	now := time.Now().UTC()
	randomPosts := createRandomPosts(t)
	randomUser := createRandomUser(t)

	posts, err := testQueries.GetPosts(context.Background(), randomPosts.ID)
	require.NoError(t, err)
	require.NotEmpty(t, posts)

	args := UpdatePostsParams{
		ID:           posts.ID,
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
	post, err := testQueries.UpdatePosts(context.Background(), args)

	// Assert
	require.NoError(t, err)
	require.NotEmpty(t, post)
	require.Equal(t, args.ID, post.ID)
	require.Equal(t, args.Title, post.Title)
	require.Equal(t, args.Content, post.Content)
	require.Equal(t, args.AuthorID, post.AuthorID)
	require.Equal(t, args.Url, post.Url)
	require.WithinDuration(t, args.UpdatedAt, post.UpdatedAt.UTC(), time.Millisecond)
	require.Equal(t, args.Status, post.Status)
	require.WithinDuration(t, args.PublishedAt, post.PublishedAt.UTC(), time.Millisecond)
	require.WithinDuration(t, args.EditedAt, post.EditedAt.UTC(), time.Millisecond)
	require.Equal(t, args.PostAuthor, post.PostAuthor)
	require.Equal(t, args.PostMimeType, post.PostMimeType)
	require.Equal(t, args.PublishedBy, post.PublishedBy)
	require.Equal(t, args.UpdatedBy, post.UpdatedBy)
}

func TestDeletePosts(t *testing.T) {
	// Arrange
	randomPosts := createRandomPosts(t)

	err := testQueries.DeletePosts(context.Background(), randomPosts.ID)
	require.NoError(t, err)

	// Act
	posts, err := testQueries.GetPosts(context.Background(), randomPosts.ID)

	// Assert
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, posts)
}

func TestListPosts(t *testing.T) {
	// Arrange
	for i := 0; i < 10; i++ {
		createRandomPosts(t)
	}

	args := ListPostsParams{
		Limit:  5,
		Offset: 5,
	}

	// Act
	posts, err := testQueries.ListPosts(context.Background(), args)
	require.NoError(t, err)

	// Assert
	require.Len(t, posts, 5)

	for _, post := range posts {
		require.NotEmpty(t, post)
	}
}
