package frog_blossom_db

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateWebsites(t *testing.T) {
	// Arrange
	arg := CreateWebsitesParams{
		Name:           "felix-fe",
		Domain:         "https://fix-contactform-type-error.felix-fe.pages.dev",
		Password:       sql.NullString{String: "mypassword", Valid: false},
		TemplateID:     sql.NullInt64{Int64: 1, Valid: true},
		BuilderEnabled: sql.NullBool{Bool: false, Valid: false},
	}
	// Act
	websites, err := testQueries.CreateWebsites(context.Background(), arg)
	// Assert
	require.NoError(t, err)
	require.NotEmpty(t, websites)

	require.Equal(t, arg.Name, websites.Name)
	require.Equal(t, arg.Domain, websites.Domain)
	require.Equal(t, arg.Password, websites.Password)
	require.Equal(t, arg.TemplateID, websites.TemplateID)
	require.Equal(t, arg.BuilderEnabled, websites.BuilderEnabled)

	require.NotZero(t, websites.ID)
}
