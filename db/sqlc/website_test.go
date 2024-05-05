package frog_blossom_db

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"testing"
)

func MockUserData() *User {
	user := &User{
		ID:       1,
		Username: "",
		Email:    "",
		Password: "",
		Role: sql.NullString{
			String: "admin",
		},
	}
	return user
}

func TestCreateWebsite(t *testing.T) {
	// Arrange
	user, err := testQueries.GetUsers(context.Background(), 5)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	args := CreateWebsiteParams{
		Name:             "felix-fe",
		Domain:           "https://fix-contactform-type-error.felix-fe.pages.dev",
		OwnerID:          user.ID,
		SelectedTemplate: sql.NullInt64{Int64: 123, Valid: true},
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
