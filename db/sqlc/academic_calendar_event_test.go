package db

import (
	"context"
	"math/rand"
	"sort"
	"testing"
	_ "time"

	"github.com/VinCPR/backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomEvent(t *testing.T) AcademicCalendarEvent {
	AcademicYear := createRandomAcademicYear(t)
	RandDate := util.RandomDate()
	arg := CreateEventParams{
		AcademicYearID: AcademicYear.ID,
		Name:           util.RandomName(),
		Type:           util.RandomString(6),
		StartDate:      RandDate,
		EndDate:        RandDate.AddDate(0, 0, 7*rand.Intn(13)),
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

	for i := 0; i < 10; i++ {
		createRandomEvent(t)
	}
	AcademicYear := util.RandomInt(2022, 2030)
	events, err := testQueries.ListEventsByAcademicYearID(context.Background(), AcademicYear)
	require.NoError(t, err)

	events_copy := events
	sort.Slice(events, func(i, j int) bool {
		return events[i].StartDate.Before(events[j].StartDate)
	})

	for i := 0; i < len(events); i++ {
		require.NotEmpty(t, events)
		require.Equal(t, events[i].AcademicYearID, AcademicYear)
		require.Equal(t, events[i].Name, events_copy[i].Name)
		require.Equal(t, events[i].Type, events_copy[i].Type)
		require.Equal(t, events[i].EndDate, events_copy[i].EndDate)
		require.Equal(t, events[i].StartDate, events_copy[i].StartDate)
		require.Equal(t, events[i].CreatedAt, events_copy[i].CreatedAt)
	}
}
