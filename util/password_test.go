package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := RandomString(20)
	passwordEncoded, err := HashPassword(password)

	require.NoError(t, err)
	err = CheckPassword(password, passwordEncoded)
	require.NoError(t, err)

	invalidpassword := RandomString(20)
	err = CheckPassword(invalidpassword, passwordEncoded)
	require.Error(t, err)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}
