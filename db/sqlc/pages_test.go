package frog_blossom_db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomPage(t *testing.T) Page {
	// Arrange
	randomUser := createRandomUser(t)

	args := CreatePagesParams{
		Domain:         "example.com",
		AuthorID:       randomUser.ID,
		PageAuthor:     randomUser.Username,
		Title:          "Homepage",
		Url:            "/home",
		MenuOrder:      1,
		ComponentType:  "Text",
		ComponentValue: "Welcome to our website!",
		PageIdentifier: "home",
		OptionID:       98765,
		OptionName:     "site_title",
		OptionValue:    "My Website",
		OptionRequired: true,
	}

	// Act
	page, err := testQueries.CreatePages(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, page)

	// Assert

	require.Equal(t, args.Domain, page.Domain)
	require.Equal(t, args.AuthorID, page.AuthorID)
	require.Equal(t, args.PageAuthor, page.PageAuthor)
	require.Equal(t, args.Title, page.Title)
	require.Equal(t, args.Url, page.Url)
	require.Equal(t, args.MenuOrder, page.MenuOrder)
	require.Equal(t, args.ComponentType, page.ComponentType)
	require.Equal(t, args.ComponentValue, page.ComponentValue)
	require.Equal(t, args.PageIdentifier, page.PageIdentifier)
	require.Equal(t, args.OptionID, page.OptionID)
	require.Equal(t, args.OptionName, page.OptionName)
	require.Equal(t, args.OptionValue, page.OptionValue)
	require.Equal(t, args.OptionRequired, page.OptionRequired)

	return page
}

func TestCreatePages(t *testing.T) {
	createRandomPage(t)
}

func TestGetPages(t *testing.T) {
	// Arrange
	randomPage := createRandomPage(t)

	// Act
	page, err := testQueries.GetPages(context.Background(), randomPage.ID)
	require.NoError(t, err)
	require.NotEmpty(t, page)

	// Assert
	require.Equal(t, randomPage.ID, page.ID)
	require.Equal(t, page.Domain, page.Domain)
	require.Equal(t, page.AuthorID, page.AuthorID)
	require.Equal(t, page.PageAuthor, page.PageAuthor)
	require.Equal(t, page.Title, page.Title)
	require.Equal(t, page.Url, page.Url)
	require.Equal(t, page.MenuOrder, page.MenuOrder)
	require.Equal(t, page.ComponentType, page.ComponentType)
	require.Equal(t, page.ComponentValue, page.ComponentValue)
	require.Equal(t, page.PageIdentifier, page.PageIdentifier)
	require.Equal(t, page.OptionID, page.OptionID)
	require.Equal(t, page.OptionName, page.OptionName)
	require.Equal(t, page.OptionValue, page.OptionValue)
	require.Equal(t, page.OptionRequired, page.OptionRequired)
}

func TestUpdatePages(t *testing.T) {
	// Arrange
	randomUser := createRandomUser(t)
	randomPage := createRandomPage(t)

	user, err := testQueries.GetUsers(context.Background(), randomUser.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	pages, err := testQueries.GetPages(context.Background(), randomPage.ID)

	args := UpdatePagesParams{
		ID:             pages.ID,
		Domain:         "example.com",
		AuthorID:       user.ID,
		PageAuthor:     user.Username,
		Title:          "Homepage",
		Url:            "/home",
		MenuOrder:      1,
		ComponentType:  "Text",
		ComponentValue: "Welcome to our website!",
		PageIdentifier: "contact",
		OptionID:       98765,
		OptionName:     "site_title",
		OptionValue:    "My Website",
		OptionRequired: true,
	}

	// Act
	page, err := testQueries.UpdatePages(context.Background(), args)

	// Assert
	require.NoError(t, err)
	require.NotEmpty(t, page)
	require.Equal(t, args.ID, page.ID)
	require.Equal(t, args.Domain, page.Domain)
	require.Equal(t, args.AuthorID, page.AuthorID)
	require.Equal(t, args.PageAuthor, page.PageAuthor)
	require.Equal(t, args.Title, page.Title)
	require.Equal(t, args.Url, page.Url)
	require.Equal(t, args.MenuOrder, page.MenuOrder)
	require.Equal(t, args.ComponentType, page.ComponentType)
	require.Equal(t, args.ComponentValue, page.ComponentValue)
	require.Equal(t, args.PageIdentifier, page.PageIdentifier)
	require.Equal(t, args.OptionID, page.OptionID)
	require.Equal(t, args.OptionName, page.OptionName)
	require.Equal(t, args.OptionValue, page.OptionValue)
	require.Equal(t, args.OptionRequired, page.OptionRequired)
}

func TestDeletePages(t *testing.T) {
	// Arrange
	randomPage := createRandomPage(t)

	err := testQueries.DeletePages(context.Background(), randomPage.ID)
	require.NoError(t, err)

	// Act
	page, err := testQueries.GetPages(context.Background(), randomPage.ID)

	// Assert
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, page)
}

func TestListPages(t *testing.T) {
	// Arrange
	for i := 0; i < 10; i++ {
		createRandomPage(t)
	}

	args := ListPagesParams{
		Limit:  5,
		Offset: 5,
	}

	// Act
	page, err := testQueries.ListPages(context.Background(), args)
	require.NoError(t, err)

	// Assert
	require.Len(t, page, 5)

	for _, page := range page {
		require.NotEmpty(t, page)
	}
}
