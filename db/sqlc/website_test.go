package frog_blossom_db

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateWebsite(t *testing.T) {
	// Arrange
	user, err := testQueries.GetUsers(context.Background(), 1)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	user64 := int64(user.ID)

	template, err := testQueries.GetTemplate(context.Background(), 1)
	require.NoError(t, err)
	require.NotEmpty(t, template)

	args := CreateWebsiteParams{
		Name:             "felix-fe",
		Domain:           "https://fix-contactform-type-error.felix-fe.pages.dev",
		OwnerID:          user64,
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
