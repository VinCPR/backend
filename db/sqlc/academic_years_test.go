package db

import (
	"context"
	"math/rand"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/VinCPR/backend/util"
)

func createRandomAcademicYear(t *testing.T) AcademicYear {
	randDate := util.RandomDate()
	arg := CreateAcademicYearParams{
		Name:      util.RandomName(),
		StartDate: randDate,
		EndDate:   randDate.AddDate(0, 0, 7*rand.Intn(13)),
	}
	academicYear, err := testQueries.CreateAcademicYear(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, academicYear)

	require.Equal(t, arg.Name, academicYear.Name)
	require.Equal(t, arg.StartDate, academicYear.StartDate)
	require.Equal(t, arg.EndDate, academicYear.EndDate)

	require.NotZero(t, academicYear.ID)
	return academicYear
}

func TestCreateAcademicYear(t *testing.T) {
	createRandomAcademicYear(t)
}

func TestGetAcademicYearByName(t *testing.T) {
	academicYear1 := createRandomAcademicYear(t)

	academicYear2, err := testQueries.GetAcademicYearByName(context.Background(), academicYear1.Name)
	require.NoError(t, err)
	require.NotEmpty(t, academicYear2)

	require.Equal(t, academicYear1.ID, academicYear2.ID)
	require.Equal(t, academicYear1.Name, academicYear2.Name)
	require.Equal(t, academicYear1.StartDate, academicYear2.StartDate)
	require.Equal(t, academicYear1.EndDate, academicYear2.EndDate)
}

func TestListAcademicYearByName(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAcademicYear(t)
	}

	arg := ListAcademicYearByEndDateParams{
		Limit:  5,
		Offset: 0,
	}

	academicYears, err := testQueries.ListAcademicYearByEndDate(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, academicYears)
	require.Len(t, academicYears, 5)
	require.True(t, sort.SliceIsSorted(academicYears, func(i, j int) bool {
		return academicYears[i].EndDate.After(academicYears[j].EndDate)
	}))
}
