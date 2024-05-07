package frog_blossom_db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomTemplate(t *testing.T) Template {
	// Arrange
	name := "Twenty Twenty"
	// Act
	template, err := testQueries.CreateTemplate(context.Background(), "Twenty Twenty")
	// Assert
	require.NoError(t, err)
	require.NotEmpty(t, template)
	require.Equal(t, name, template.Name)

	return template
}

func TestCreateTemplate(t *testing.T) {
	createRandomTemplate(t)
}

func TestGetTemplate(t *testing.T) {
	// Arrange
	randomTemplate := createRandomTemplate(t)

	// Act
	template, err := testQueries.GetTemplate(context.Background(), randomTemplate.ID)
	require.NoError(t, err)
	require.NotEmpty(t, template)

	// Assert
	require.Equal(t, randomTemplate.ID, template.ID)
	require.Equal(t, randomTemplate.Name, template.Name)
}
