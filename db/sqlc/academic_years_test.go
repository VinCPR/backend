package db

import (
	"context"
	"math/rand"
	"testing"

	"github.com/VinCPR/backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomAcademicYear(t *testing.T) AcademicYear {
	RandDate := util.RandomDate()
	arg := CreateAcademicYearParams{
		Name:      util.RandomName(),
		StartDate: RandDate,
		EndDate:   RandDate.AddDate(0, 0, 7*rand.Intn(13)),
	}
	academic_year, err := testQueries.CreateAcademicYear(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, academic_year)

	require.Equal(t, arg.Name, academic_year.Name)
	require.Equal(t, arg.StartDate, academic_year.StartDate)
	require.Equal(t, arg.EndDate, academic_year.EndDate)

	require.NotZero(t, academic_year.ID)
	return academic_year
}

func TestCreateAcademicYear(t *testing.T) {
	createRandomAcademicYear(t)
}

func TestGetAcademicYearByName(t *testing.T) {
	academic_year1 := createRandomAcademicYear(t)

	academic_year2, err := testQueries.GetAcademicYearByName(context.Background(), academic_year1.Name)
	require.NoError(t, err)
	require.NotEmpty(t, academic_year2)

	require.Equal(t, academic_year1.ID, academic_year2.ID)
	require.Equal(t, academic_year1.Name, academic_year2.Name)
	require.Equal(t, academic_year1.StartDate, academic_year2.StartDate)
	require.Equal(t, academic_year1.EndDate, academic_year2.EndDate)
}

// func TestListAcademicYearByName(t *testing.T) {
// 	for i := 0; i < 10; i++ {
// 		createRandomAcademicYear(t)

// 		arg := ListAcademicYearByNameParams{
// 			Limit:  5,
// 			Offset: 0,
// 		}

// 		academic_years, err := testQueries.ListAcademicYearByName(context.Background(), arg)
// 		require.NoError(t, err)
// 		require.NotEmpty(t, academic_years)
// 		academic_years_copy := academic_years
// 		sort.Slice(academic_years, func(i, j int) bool {
// 			return academic_years[i].Name < academic_years[j].Name
// 		})

// 		for i := 0; i < len(academic_years); i++ {
// 			require.NotEmpty(t, academic_years)
// 			require.Equal(t, academic_years[i].ID, academic_years_copy[i].ID)
// 			require.Equal(t, academic_years[i].Email, academic_years_copy[i].Email)
// 		}
// 	}
// }
