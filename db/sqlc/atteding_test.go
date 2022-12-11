package db

import (
	"context"
	"sort"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
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
	allAttendings, err := testQueries.ListAttendingsByName(context.Background(), ListAttendingsByNameParams{
		Limit:  100,
		Offset: 0,
	})
	require.NoError(t, err)

	sort.SliceIsSorted(allAttendings, func(i, j int) bool {
		return allAttendings[i].FirstName < allAttendings[j].FirstName ||
			(allAttendings[i].FirstName == allAttendings[j].FirstName && allAttendings[i].LastName < allAttendings[j].LastName)
	})

	attendings, err := testQueries.ListAttendingsByName(context.Background(), ListAttendingsByNameParams{
		Limit:  int32(n),
		Offset: 0,
	})
	require.NoError(t, err)

	require.Len(t, attendings, n)
	require.EqualValues(t, attendings, allAttendings[:n])
	spew.Dump(attendings)
}
