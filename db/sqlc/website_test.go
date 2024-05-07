package frog_blossom_db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomWebsite(t *testing.T) Website {
	// Arrange
	randomUser := createRandomUser(t)
	randomTemplate := createRandomTemplate(t)

	user, err := testQueries.GetUsers(context.Background(), randomUser.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	template, err := testQueries.GetTemplate(context.Background(), randomTemplate.ID)
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

	return website
}

func TestCreateWebsite(t *testing.T) {
	createRandomWebsite(t)
}

func TestUpdateWebsite(t *testing.T) {
	// Arrange
	user, err := testQueries.GetUsers(context.Background(), 9)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	template, err := testQueries.GetTemplate(context.Background(), 9)
	require.NoError(t, err)
	require.NotEmpty(t, template)

	args := UpdateWebsiteParams{
		ID:               9,
		Name:             "Mystery on Monster Island",
		Domain:           "berkeley.edu",
		OwnerID:          9,
		SelectedTemplate: 9,
	}
	// Act
	website, err := testQueries.UpdateWebsite(context.Background(), args)
	// Assert
	require.NoError(t, err)
	require.NotEmpty(t, website)
	require.Equal(t, args.ID, website.ID)
	require.Equal(t, args.Name, website.Name)
	require.Equal(t, args.Domain, website.Domain)
	require.Equal(t, args.OwnerID, website.OwnerID)
	require.Equal(t, args.SelectedTemplate, website.SelectedTemplate)
}

func TestGetWebsite(t *testing.T) {
	// Arrange
	randomWebsite := createRandomWebsite(t)

	user, err := testQueries.GetUsers(context.Background(), 9)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	template, err := testQueries.GetTemplate(context.Background(), 9)
	require.NoError(t, err)
	require.NotEmpty(t, template)
	// Act
	website, err := testQueries.GetWebsite(context.Background(), randomWebsite.ID)
	require.NoError(t, err)
	require.NotEmpty(t, website)

	// Assert
	require.Equal(t, randomWebsite.ID, website.ID)
	require.Equal(t, randomWebsite.OwnerID, website.OwnerID)
	require.Equal(t, randomWebsite.SelectedTemplate, website.SelectedTemplate)
}

func TestDeleteWebsite(t *testing.T) {
	// Arrange
	err := testQueries.DeleteWebsite(context.Background(), 9)
	require.NoError(t, err)
	// Act
	website, err := testQueries.GetWebsite(context.Background(), 9)
	// Assert
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, website)
}
