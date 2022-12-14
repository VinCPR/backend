package db

import (
	"context"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/VinCPR/backend/util"
)

func createRandomAttending(t *testing.T, user User) Attending {
	arg := CreateAttendingParams{
		UserID:    user.ID,
		FirstName: util.RandomName(),
		LastName:  util.RandomName(),
		Mobile:    util.RandomMobile(),
	}

	attending, err := testQueries.CreateAttending(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, attending)

	require.Equal(t, arg.UserID, attending.UserID)
	require.Equal(t, arg.FirstName, attending.FirstName)
	require.Equal(t, arg.LastName, attending.LastName)
	require.Equal(t, arg.Mobile, attending.Mobile)

	require.NotZero(t, attending.ID)
	require.NotZero(t, attending.CreatedAt)

	return attending
}

func TestCreateAttending(t *testing.T) {
	user := createRandomUser(t)
	createRandomAttending(t, user)
}

func TestGetAttendingByUserId(t *testing.T) {
	user := createRandomUser(t)
	attending1 := createRandomAttending(t, user)
	attending2, err := testQueries.GetAttendingByUserId(context.Background(), user.ID)
	require.NoError(t, err)
	require.NotEmpty(t, attending2)

	require.Equal(t, attending1.ID, attending2.ID)
	require.Equal(t, attending1.UserID, attending2.UserID)
	require.Equal(t, attending1.FirstName, attending2.FirstName)
	require.Equal(t, attending1.LastName, attending2.LastName)
	require.Equal(t, attending1.Mobile, attending2.Mobile)
	require.WithinDuration(t, attending1.CreatedAt, attending2.CreatedAt, time.Second)
}

func TestListAttendingsByName(t *testing.T) {
	var n = 5
	for i := 0; i < n; i++ {
		user := createRandomUser(t)
		createRandomAttending(t, user)

	}
	user := createRandomUser(t)
	lastAttending := createRandomAttending(t, user)
	argtest := ListAttendingsByNameParams{
		Limit:  int32(lastAttending.ID),
		Offset: 0,
	}

	attendingList, err := testQueries.ListAttendingsByName(context.Background(), argtest)
	require.NoError(t, err)

	sort.SliceIsSorted(attendingList, func(i, j int) bool {
		return attendingList[i].FirstName < attendingList[j].FirstName ||
			(attendingList[i].FirstName == attendingList[j].FirstName && attendingList[i].LastName < attendingList[j].LastName)
	})

	arg := ListAttendingsByNameParams{
		Limit:  int32(n),
		Offset: 0,
	}

	attendings, err := testQueries.ListAttendingsByName(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, attendings, n)
	require.EqualValues(t, attendings, attendingList[:n])
}
