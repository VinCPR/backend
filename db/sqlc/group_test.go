package db

import (
	"context"
	"sort"
	"testing"
	"time"

	"github.com/VinCPR/backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomGroup(t *testing.T, academic_year AcademicYear) Group {

	arg := CreateGroupParams{
		AcademicYearID: academic_year.ID,
		Name:           util.RandomString(10),
	}

	group, err := testQueries.CreateGroup(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, group)

	require.Equal(t, arg.Name, group.Name)
	require.Equal(t, arg.AcademicYearID, group.AcademicYearID)

	require.NotZero(t, group.ID)
	require.NotZero(t, group.CreatedAt)

	return group
}

func TestCreateGroup(t *testing.T) {
	academic_year := createRandomAcademicYear(t)
	createRandomGroup(t, academic_year)

}

func TestGetGroupByID(t *testing.T) {
	academic_year := createRandomAcademicYear(t)
	group1 := createRandomGroup(t, academic_year)
	group2, err := testQueries.GetGroupByID(context.Background(), group1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, group2)

	require.Equal(t, group1.ID, group2.ID)
	require.Equal(t, group1.AcademicYearID, group2.AcademicYearID)
	require.Equal(t, group1.Name, group2.Name)
	require.WithinDuration(t, group2.CreatedAt, group2.CreatedAt, time.Second)
}

func TestGetGroupByIndex(t *testing.T) {
	academic_year := createRandomAcademicYear(t)
	group1 := createRandomGroup(t, academic_year)
	group2, err := testQueries.GetGroupByIndex(context.Background(), GetGroupByIndexParams{group1.AcademicYearID, group1.Name})
	require.NoError(t, err)
	require.NotEmpty(t, group2)

	require.Equal(t, group1.ID, group2.ID)
	require.Equal(t, group1.AcademicYearID, group2.AcademicYearID)
	require.Equal(t, group1.Name, group2.Name)
	require.WithinDuration(t, group2.CreatedAt, group2.CreatedAt, time.Second)
}

// TODO refactor test for list queries
func TestListGroupsByName(t *testing.T) {
	var (
		academicYear  = createRandomAcademicYear(t)
		createdGroups = make([]Group, 0)
		n             = 5
	)
	for i := 0; i < n; i++ {
		createdGroups = append(createdGroups, createRandomGroup(t, academicYear))
	}

	groups, err := testQueries.ListGroupsByName(context.Background(), academicYear.ID)
	require.NoError(t, err)
	require.NotEmpty(t, groups)

	sort.Slice(createdGroups, func(i, j int) bool {
		return createdGroups[i].Name < createdGroups[j].Name ||
			(createdGroups[i].Name == createdGroups[j].Name && createdGroups[i].ID < createdGroups[j].ID)
	})

	for i := 0; i < n; i++ {
		require.Equal(t, groups[i].AcademicYearID, academicYear.ID)
		require.Equal(t, groups[i].Name, groups[i].Name)
	}
}
