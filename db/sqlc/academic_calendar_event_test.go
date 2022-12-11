package db

import (
	"context"
	"math/rand"
	"sort"
	"testing"
	_ "time"

	"github.com/stretchr/testify/require"

	"github.com/VinCPR/backend/util"
)

func createRandomEvent(t *testing.T) AcademicCalendarEvent {
	academicYear := createRandomAcademicYear(t)
	randDate := util.RandomDate()
	arg := CreateEventParams{
		AcademicYearID: academicYear.ID,
		Name:           util.RandomName(),
		Type:           util.RandomString(6),
		StartDate:      randDate,
		EndDate:        randDate.AddDate(0, 0, 7*rand.Intn(13)),
	}

	event, err := testQueries.CreateEvent(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, event)

	require.Equal(t, arg.AcademicYearID, event.AcademicYearID)
	require.Equal(t, arg.Name, event.Name)
	require.Equal(t, arg.StartDate, event.StartDate)
	require.Equal(t, arg.EndDate, event.EndDate)
	require.Equal(t, arg.Type, event.Type)

	require.NotZero(t, event.ID)
	require.NotZero(t, event.CreatedAt)
	return event
}

func TestCreateEvent(t *testing.T) {
	createRandomEvent(t)
}

func TestListEventsByAcademicYearID(t *testing.T) {
	var (
		n            = 10
		academicYear = createRandomAcademicYear(t)
		args         = make([]CreateEventsParams, 0)
	)
	{ // Test create multiple events
		for i := 0; i < n; i++ {
			randDate := util.RandomDate()
			args = append(args, CreateEventsParams{
				AcademicYearID: academicYear.ID,
				Name:           util.RandomName(),
				Type:           util.RandomString(6),
				StartDate:      randDate,
				EndDate:        randDate.AddDate(0, 0, 7*rand.Intn(13)),
			})
		}
		eventsLen, err := testQueries.CreateEvents(context.Background(), args)
		require.NoError(t, err)
		require.EqualValues(t, n, eventsLen)
	}
	{
		sort.Slice(args, func(i, j int) bool {
			return args[i].StartDate.Before(args[j].StartDate)
		})
		events, err := testQueries.ListEventsByAcademicYearID(context.Background(), academicYear.ID)
		require.NoError(t, err)
		require.Len(t, events, n)
		for i := 0; i < n; i++ {
			require.Equal(t, args[i].AcademicYearID, events[i].AcademicYearID)
			require.Equal(t, args[i].Name, events[i].Name)
			require.Equal(t, args[i].StartDate, events[i].StartDate)
			require.Equal(t, args[i].EndDate, events[i].EndDate)
			require.Equal(t, args[i].Type, events[i].Type)

			require.NotZero(t, events[i].ID)
			require.NotZero(t, events[i].CreatedAt)
		}
	}
}
