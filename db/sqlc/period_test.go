package db

import (
	"context"
	"math/rand"
	"sort"
	"testing"

	"github.com/VinCPR/backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomPeriod(t *testing.T) Period {
	AcademicYear := createRandomAcademicYear(t)
	RandDate := util.RandomDate()
	arg := CreatePeriodParams{
		AcademicYearID: AcademicYear.ID,
		Name:           util.RandomName(),
		StartDate:      RandDate,
		EndDate:        RandDate.AddDate(0, 0, 7*rand.Intn(13)),
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
	createRandomPeriod(t)
}

func TestGetPeriodByID(t *testing.T) {
	period1 := createRandomPeriod(t)
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
	for i := 0; i < 10; i++ {
		createRandomPeriod(t)
	}
	AcademicYear := util.RandomInt(2022, 2030)
	periods, err := testQueries.ListPeriodsByStartDate(context.Background(), AcademicYear)
	require.NoError(t, err)

	periods_copy := periods
	sort.Slice(periods, func(i, j int) bool {
		return periods[i].StartDate.Before(periods[j].StartDate)
	})

	for i := 0; i < len(periods); i++ {
		require.NotEmpty(t, periods)
		require.Equal(t, periods[i].AcademicYearID, AcademicYear)
		require.Equal(t, periods[i].Name, periods_copy[i].Name)
		require.Equal(t, periods[i].EndDate, periods_copy[i].EndDate)
		require.Equal(t, periods[i].StartDate, periods_copy[i].StartDate)
	}
}
