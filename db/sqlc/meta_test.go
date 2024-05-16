package frog_blossom_db

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomMeta(t *testing.T) Meta {
	// Arrange
	randomPage := createRandomPage(t)
	randomPosts := createRandomPosts(t)

	args := CreateMetaParams{
		PageID:          sql.NullInt64{Int64: randomPage.ID, Valid: true},
		PostsID:         sql.NullInt64{Int64: randomPosts.ID, Valid: true},
		MetaTitle:       sql.NullString{String: "Sample Meta Title", Valid: true},
		MetaDescription: sql.NullString{String: "Sample Meta Description", Valid: true},
		MetaRobots:      sql.NullString{String: "index, follow", Valid: true},
		MetaOgImage:     sql.NullString{String: "https://example.com/image.jpg", Valid: true},
		Locale:          sql.NullString{String: "ja_JP", Valid: true},
		PageAmount:      3,
		SiteLanguage: sql.NullString{
			String: "ja", Valid: true,
		},
		MetaKey:   "_thumbnail_id",
		MetaValue: "12345",
	}

	// Act
	meta, err := testQueries.CreateMeta(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, meta)

	// Assert
	require.Equal(t, args.PageID, meta.PageID)
	require.Equal(t, args.PostsID, meta.PostsID)
	require.Equal(t, args.MetaTitle, meta.MetaTitle)
	require.Equal(t, args.MetaDescription, meta.MetaDescription)
	require.Equal(t, args.MetaRobots, meta.MetaRobots)
	require.Equal(t, args.MetaOgImage, meta.MetaOgImage)
	require.Equal(t, args.Locale, meta.Locale)
	require.Equal(t, args.PageAmount, meta.PageAmount)
	require.Equal(t, args.SiteLanguage, meta.SiteLanguage)
	require.Equal(t, args.MetaKey, meta.MetaKey)
	require.Equal(t, args.MetaValue, meta.MetaValue)

	return meta
}

func TestCreateMeta(t *testing.T) {
	createRandomMeta(t)
}

func TestGetMeta(t *testing.T) {
	// Arrange
	randomMeta := createRandomMeta(t)

	// Act
	meta, err := testQueries.GetMeta(context.Background(), randomMeta.ID)
	require.NoError(t, err)
	require.NotEmpty(t, meta)

	// Assert
	require.Equal(t, randomMeta.ID, meta.ID)
	require.Equal(t, randomMeta.MetaTitle, meta.MetaTitle)
	require.Equal(t, randomMeta.MetaDescription, meta.MetaDescription)
	require.Equal(t, randomMeta.MetaRobots, meta.MetaRobots)
	require.Equal(t, randomMeta.MetaOgImage, meta.MetaOgImage)
	require.Equal(t, randomMeta.Locale, meta.Locale)
	require.Equal(t, randomMeta.PageAmount, meta.PageAmount)
	require.Equal(t, randomMeta.SiteLanguage, meta.SiteLanguage)
	require.Equal(t, randomMeta.MetaKey, meta.MetaKey)
	require.Equal(t, randomMeta.MetaValue, meta.MetaValue)
}

func TestUpdateMeta(t *testing.T) {
	// Arrange
	randomPage := createRandomPage(t)
	randomPosts := createRandomPosts(t)
	randomMeta := createRandomMeta(t)

	pages, err := testQueries.GetPages(context.Background(), randomPage.ID)
	require.NoError(t, err)
	require.NotEmpty(t, pages)

	posts, err := testQueries.GetPosts(context.Background(), randomPosts.ID)
	require.NoError(t, err)
	require.NotEmpty(t, posts)

	metas, err := testQueries.GetMeta(context.Background(), randomMeta.ID)
	require.NoError(t, err)
	require.NotEmpty(t, metas)

	args := UpdateMetaParams{
		ID:              metas.ID,
		PageID:          sql.NullInt64{Int64: randomPage.ID, Valid: true},
		PostsID:         sql.NullInt64{Int64: randomPosts.ID, Valid: true},
		MetaTitle:       sql.NullString{String: "Sample Meta Title", Valid: true},
		MetaDescription: sql.NullString{String: "Sample Meta Description", Valid: true},
		MetaRobots:      sql.NullString{String: "index, follow", Valid: true},
		MetaOgImage:     sql.NullString{String: "https://example.com/image.jpg", Valid: true},
		Locale:          sql.NullString{String: "ja_JP", Valid: true},
		PageAmount:      3,
		SiteLanguage: sql.NullString{
			String: "ja", Valid: true,
		},
		MetaKey:   "_thumbnail_id",
		MetaValue: "12345",
	}

	// Act
	meta, err := testQueries.UpdateMeta(context.Background(), args)

	// Assert
	require.NoError(t, err)
	require.NotEmpty(t, meta)
	require.Equal(t, args.ID, meta.ID)
	require.Equal(t, args.PageID, meta.PageID)
	require.Equal(t, args.PostsID, meta.PostsID)
	require.Equal(t, args.MetaTitle, meta.MetaTitle)
	require.Equal(t, args.MetaDescription, meta.MetaDescription)
	require.Equal(t, args.MetaRobots, meta.MetaRobots)
	require.Equal(t, args.MetaOgImage, meta.MetaOgImage)
	require.Equal(t, args.Locale, meta.Locale)
	require.Equal(t, args.PageAmount, meta.PageAmount)
	require.Equal(t, args.SiteLanguage, meta.SiteLanguage)
	require.Equal(t, args.MetaKey, meta.MetaKey)
	require.Equal(t, args.MetaValue, meta.MetaValue)
}

func TestDeleteMeta(t *testing.T) {
	// Arrange
	randomMeta := createRandomMeta(t)

	err := testQueries.DeleteMeta(context.Background(), randomMeta.ID)
	require.NoError(t, err)

	// Act
	meta, err := testQueries.GetMeta(context.Background(), randomMeta.ID)

	// Assert
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, meta)
}

func TestListMeta(t *testing.T) {
	// Arrange
	for i := 0; i < 10; i++ {
		createRandomMeta(t)
	}

	args := ListMetaParams{
		Limit:  5,
		Offset: 5,
	}

	// Act
	meta, err := testQueries.ListMeta(context.Background(), args)
	require.NoError(t, err)

	// Assert
	require.Len(t, meta, 5)

	for _, meta := range meta {
		require.NotEmpty(t, meta)
	}
}
