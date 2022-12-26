package db

import (
	"context"
	"fmt"
	"testing"
	"time"

<<<<<<< HEAD
	"github.com/stretchr/testify/require"

	"github.com/VinCPR/backend/util"
=======
	"github.com/VinCPR/backend/util"
	"github.com/stretchr/testify/require"
>>>>>>> cfc0062 (add sql and test for hospital, specialty, service and service to attending)
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Email:          util.RandomEmail(),
		HashedPassword: util.RandomString(8),
		RoleName:       util.RandomName(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.RoleName, user.RoleName)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)

}

func TestGetUserByEmail(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUserByEmail(context.Background(), user1.Email)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.RoleName, user2.RoleName)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestListUsersByID(t *testing.T) {
<<<<<<< HEAD
	for i := 0; i < 15; i++ {
=======
	for i := 0; i < 5; i++ {
>>>>>>> cfc0062 (add sql and test for hospital, specialty, service and service to attending)
		createRandomUser(t)
	}

	arg := ListUsersByIDParams{
		Limit:  5,
		Offset: 10,
	}

	users, err := testQueries.ListUsersByID(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, users, 5)

	for _, user := range users {
		require.NotEmpty(t, user)
		fmt.Println(user.ID)
	}
}
