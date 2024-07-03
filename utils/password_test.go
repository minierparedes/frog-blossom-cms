package utils

import (
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestHashPassword(t *testing.T) {
	// Arrange
	password := RandomString(6)

	// Act
	hashedPassword, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	// Assert
	err = CheckPassword(password, hashedPassword)
	require.NoError(t, err)
}

func TestWrongHashPassword(t *testing.T) {
	// Arrange
	password := RandomString(6)
	wrongPassword := RandomString(6)

	// Act
	hashedPassword, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	// Assert
	err = CheckPassword(wrongPassword, hashedPassword)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}
