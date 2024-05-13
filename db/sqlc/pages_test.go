package frog_blossom_db

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomPage(t *testing.T) Page {
	// Arrange
	randomUser := createRandomUser(t)

	args := CreatePagesParams{
		Domain:         "example.com",
		PageAuthor:     randomUser.ID,
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

	// Assert
	require.NoError(t, err)
	require.NotEmpty(t, page)
	require.Equal(t, args.Domain, page.Domain)
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
