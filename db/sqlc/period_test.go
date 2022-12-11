package db

import (
	"context"
	"math/rand"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/VinCPR/backend/util"
)

func createRandomPeriod(t *testing.T, academicYearID int64) Period {
	randDate := util.RandomDate()
	arg := CreatePeriodParams{
		AcademicYearID: academicYearID,
		Name:           util.RandomName(),
		StartDate:      randDate,
		EndDate:        randDate.AddDate(0, 0, 7*rand.Intn(13)),
	}
	period, err := testQueries.CreatePeriod(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, period)

	require.Equal(t, arg.AcademicYearID, period.AcademicYearID)
	require.Equal(t, arg.Name, period.Name)
	require.Equal(t, arg.StartDate, period.StartDate)
	require.Equal(t, arg.EndDate, period.EndDate)

	require.NotZero(t, period.ID)
	return period
}

func TestCreatePeriod(t *testing.T) {
	academicYear := createRandomAcademicYear(t)
	createRandomPeriod(t, academicYear.ID)
}

func TestGetPeriodByID(t *testing.T) {
	academicYear := createRandomAcademicYear(t)
	period1 := createRandomPeriod(t, academicYear.ID)
	period2, err := testQueries.GetPeriodByID(context.Background(), period1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, period2)

	require.Equal(t, period1.ID, period2.ID)
	require.Equal(t, period1.AcademicYearID, period2.AcademicYearID)
	require.Equal(t, period1.Name, period2.Name)
	require.Equal(t, period1.StartDate, period2.StartDate)
	require.Equal(t, period1.EndDate, period2.EndDate)
}

func TestListPeriodsByStartDate(t *testing.T) {
	var (
		academicYear   = createRandomAcademicYear(t)
		createdPeriods = make([]Period, 0)
		n              = 10
	)
	for i := 0; i < n; i++ {
		createdPeriods = append(createdPeriods, createRandomPeriod(t, academicYear.ID))
	}

	periods, err := testQueries.ListPeriodsByStartDate(context.Background(), academicYear.ID)
	require.NoError(t, err)
	require.NotEmpty(t, periods)

	sort.Slice(createdPeriods, func(i, j int) bool {
		return createdPeriods[i].StartDate.Before(createdPeriods[j].StartDate)
	})

	for i := 0; i < n; i++ {
		require.Equal(t, periods[i].AcademicYearID, academicYear.ID)
		require.Equal(t, periods[i].Name, createdPeriods[i].Name)
		require.Equal(t, periods[i].EndDate, createdPeriods[i].EndDate)
		require.Equal(t, periods[i].StartDate, createdPeriods[i].StartDate)
	}
}
