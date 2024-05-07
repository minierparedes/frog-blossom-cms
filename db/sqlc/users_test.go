package frog_blossom_db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	// Arrange
	args := CreateUsersParams{
		Username: "cshopcott6",
		Email:    "cshopcott6@friendfeed.com",
		Password: "pP7<8jrQbwS",
		Role: sql.NullString{
			String: "admin", Valid: true,
		},
		FirstName: sql.NullString{
			String: "Cointon", Valid: true,
		},
		LastName: sql.NullString{
			String: "Shopcott", Valid: true,
		},
		AvatarUrl: sql.NullString{
			String: "https://robohash.org/nihildelectussed.png?size=50x50&set=set1", Valid: true,
		},
		Bio: sql.NullString{
			String: "Morbi porttitor lorem id ligula. Suspendisse ornare consequat lectus. In est risus, auctor sed, tristique in, tempus sit amet, sem. Fusce consequat. Nulla nisl. Nunc nisl.", Valid: true,
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

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)

}

func TestGetUsers(t *testing.T) {
	// Arrange
	randomUser := createRandomUser(t)
	user, err := testQueries.GetUsers(context.Background(), randomUser.ID)
	// Act
	require.NoError(t, err)
	require.NotEmpty(t, user)

	// Assert
	require.Equal(t, randomUser.ID, user.ID)
	require.Equal(t, randomUser.Username, user.Username)
	require.Equal(t, randomUser.Email, user.Email)
	require.Equal(t, randomUser.Password, user.Password)
	require.Equal(t, randomUser.Role, user.Role)
	require.Equal(t, randomUser.FirstName, user.FirstName)
	require.Equal(t, randomUser.LastName, user.LastName)
	require.Equal(t, randomUser.AvatarUrl, user.AvatarUrl)
	require.Equal(t, randomUser.Bio, user.Bio)
	require.WithinDuration(t, randomUser.CreatedAt.Time, user.CreatedAt.Time, time.Second)
}
