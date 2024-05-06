package frog_blossom_db

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateUser(t *testing.T) {
	// Arrange
	args := CreateUsersParams{
		Username: "John Doe",
		Email:    "sdsdsdsadsasadadas@example.com",
		Password: "12345",
		Role: sql.NullString{
			String: "admin", Valid: true,
		},
		FirstName: sql.NullString{
			String: "John", Valid: true,
		},
		LastName: sql.NullString{
			String: "Doe", Valid: true,
		},
		AvatarUrl: sql.NullString{
			String: "avatar", Valid: true,
		},
		Bio: sql.NullString{
			String: "bio", Valid: true,
		},
	}

	// Act

	user, err := testQueries.CreateUsers(context.Background(), args)
	// Assert
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, args.Username, user.Username)
	require.Equal(t, args.Email, user.Email)
	require.Equal(t, args.Password, user.Password)
	require.Equal(t, args.Role, user.Role)
	require.Equal(t, args.Bio, user.Bio)

}
