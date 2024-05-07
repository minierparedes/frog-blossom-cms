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

func TestGetUser(t *testing.T) {
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

func TestUpdateUser(t *testing.T) {
	// Arrange
	args := UpdateUsersParams{
		ID:        11,
		Username:  "holyunin8 hello nurse",
		Email:     "holyunin8@si.edu",
		Password:  "gY0_OzLmifL1",
		Role:      sql.NullString{String: "user", Valid: true},
		FirstName: sql.NullString{String: "Hillery", Valid: true},
		LastName:  sql.NullString{String: "Olyunin", Valid: true},
		AvatarUrl: sql.NullString{String: "https://robohash.org/veritatisquaeratnemo.png?size=50x50&set=set1", Valid: true},
		Bio:       sql.NullString{String: "Aenean fermentum. Donec ut mauris eget massa tempor convallis. Nulla neque libero, convallis eget, eleifend luctus, ultricies eu, nibh.", Valid: true},
		UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
	}
	// Act
	user, err := testQueries.UpdateUsers(context.Background(), args)
	// Assert
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, args.ID, user.ID)
	require.Equal(t, args.Username, user.Username)
	require.Equal(t, args.Email, user.Email)
	require.Equal(t, args.Password, user.Password)
	require.Equal(t, args.Role, user.Role)
	require.Equal(t, args.FirstName, user.FirstName)
	require.Equal(t, args.LastName, user.LastName)
	require.Equal(t, args.AvatarUrl, user.AvatarUrl)
	require.Equal(t, args.Bio, user.Bio)
	require.WithinDuration(t, args.UpdatedAt.Time, user.UpdatedAt.Time, time.Second)
}

func TestDeleteUser(t *testing.T) {
	// Arrange
	randomUser := createRandomUser(t)

	err := testQueries.DeleteUsers(context.Background(), randomUser.ID)
	require.NoError(t, err)
	// Act
	user, err := testQueries.GetUsers(context.Background(), randomUser.ID)
	// Assert
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user)
}

func TestListUsers(t *testing.T) {
	// Arrange
	for i := 0; i < 10; i++ {
		createRandomUser(t)
	}

	args := ListUsersParams{
		Limit:  5,
		Offset: 5,
	}
	// Act
	users, err := testQueries.ListUsers(context.Background(), args)
	require.NoError(t, err)
	// Assert
	require.Len(t, users, 5)

	for _, user := range users {
		require.NotEmpty(t, user)
	}
}
