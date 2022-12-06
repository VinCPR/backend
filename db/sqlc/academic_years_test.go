package db

import (
	"context"
	"math/rand"
	"sort"
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

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Email:          util.RandomString(6),
		HashedPassword: util.RandomString(6),
		RoleName:       util.RandomString(6),
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

func TestListUsersByName(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomUser(t)
	}

	arg := ListUsersByNameParams{
		Limit:  5,
		Offset: 0,
	}

	users, err := testQueries.ListUsersByName(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, users)
	users_copy := users
	sort.Slice(users, func(i, j int) bool {
		return users[i].RoleName < users[j].RoleName
	})

	for i := 0; i < len(users); i++ {
		require.NotEmpty(t, users)
		require.Equal(t, users[i].ID, users_copy[i].ID)
		require.Equal(t, users[i].Email, users_copy[i].Email)
	}
}
