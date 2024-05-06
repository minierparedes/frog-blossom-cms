package frog_blossom_db

import (
	"context"
	"testing"

	"github.com/reflection/frog_blossom_db/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateWebsite(t *testing.T) {
	// Arrange
	user, err := testQueries.GetUsers(context.Background(), utils.RandomID())
	require.NoError(t, err)
	require.NotEmpty(t, user)

	template, err := testQueries.GetTemplate(context.Background(), utils.RandomID())
	require.NoError(t, err)
	require.NotEmpty(t, template)

	args := CreateWebsiteParams{
		Name:             "felix-fe",
		Domain:           "https://fix-contactform-type-error.felix-fe.pages.dev",
		OwnerID:          user.ID,
		SelectedTemplate: template.ID,
	}
	// Act
	website, err := testQueries.CreateWebsite(context.Background(), args)
	// Assert
	require.NoError(t, err)
	require.NotEmpty(t, website)

	require.Equal(t, args.Name, website.Name)
	require.Equal(t, args.Domain, website.Domain)
	require.Equal(t, args.OwnerID, website.OwnerID)
	require.Equal(t, args.SelectedTemplate, website.SelectedTemplate)

	require.NotZero(t, website.ID)
}

func TestGetWebsite(t *testing.T) {
	// Arrange
	user, err := testQueries.GetUsers(context.Background(), 1)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	template, err := testQueries.GetTemplate(context.Background(), 1)
	require.NoError(t, err)
	require.NotEmpty(t, template)
	// Act
	website, err := testQueries.GetWebsite(context.Background(), 1)
	require.NoError(t, err)
	require.NotEmpty(t, website)

	// Assert
	require.Equal(t, user.ID, website.OwnerID)
	require.Equal(t, template.ID, website.SelectedTemplate)
}
