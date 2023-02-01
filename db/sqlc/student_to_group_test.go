package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomStudentToGroup(t *testing.T, academic_year AcademicYear, student Student, group Group) StudentToGroup {

	arg := CreateStudentToGroupParams{
		AcademicYearID: academic_year.ID,
		StudentID:      student.ID,
		GroupID:        group.ID,
	}

	student_to_group, err := testQueries.CreateStudentToGroup(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, student_to_group)

	require.Equal(t, arg.GroupID, student_to_group.GroupID)
	require.Equal(t, arg.AcademicYearID, student_to_group.AcademicYearID)
	require.Equal(t, arg.StudentID, student_to_group.StudentID)

	require.NotZero(t, student_to_group.ID)
	require.NotZero(t, student_to_group.CreatedAt)

	return student_to_group
}

func TestCreateStudentToGroup(t *testing.T) {
	academic_year := createRandomAcademicYear(t)
	user := createRandomUser(t)
	student := createRandomStudent(t, user)
	group := createRandomGroup(t, academic_year)
	createRandomStudentToGroup(t, academic_year, student, group)

}

func TestGetStudentToGroupByAcademicYearID(t *testing.T) {
	academic_year := createRandomAcademicYear(t)
	user := createRandomUser(t)
	student := createRandomStudent(t, user)
	group := createRandomGroup(t, academic_year)
	student_to_group1 := createRandomStudentToGroup(t, academic_year, student, group)
	student_to_group2, err := testQueries.GetStudentToGroupByAcademicYearID(context.Background(), student_to_group1.AcademicYearID)
	require.NoError(t, err)
	require.NotEmpty(t, student_to_group2)

	for i := 0; i < len(student_to_group2); i++ {
		require.Equal(t, student_to_group1.ID, student_to_group2[i].ID)
		require.Equal(t, student_to_group1.AcademicYearID, student_to_group2[i].AcademicYearID)
		require.Equal(t, student_to_group1.StudentID, student_to_group2[i].StudentID)
		require.Equal(t, student_to_group1.GroupID, student_to_group2[i].GroupID)
		require.WithinDuration(t, student_to_group1.CreatedAt, student_to_group2[i].CreatedAt, time.Second)
	}
}

func TestGetStudentToGroupByStudentID(t *testing.T) {
	academic_year := createRandomAcademicYear(t)
	user := createRandomUser(t)
	student := createRandomStudent(t, user)
	group := createRandomGroup(t, academic_year)
	student_to_group1 := createRandomStudentToGroup(t, academic_year, student, group)
	student_to_group2, err := testQueries.GetStudentToGroupByStudentID(context.Background(), student_to_group1.StudentID)
	require.NoError(t, err)
	require.NotEmpty(t, student_to_group2)

	for i := 0; i < len(student_to_group2); i++ {
		require.Equal(t, student_to_group1.ID, student_to_group2[i].ID)
		require.Equal(t, student_to_group1.AcademicYearID, student_to_group2[i].AcademicYearID)
		require.Equal(t, student_to_group1.StudentID, student_to_group2[i].StudentID)
		require.Equal(t, student_to_group1.GroupID, student_to_group2[i].GroupID)
		require.WithinDuration(t, student_to_group1.CreatedAt, student_to_group2[i].CreatedAt, time.Second)
	}
}

func TestGetStudentToGroupByGroupID(t *testing.T) {
	academic_year := createRandomAcademicYear(t)
	user := createRandomUser(t)
	student := createRandomStudent(t, user)
	group := createRandomGroup(t, academic_year)
	student_to_group1 := createRandomStudentToGroup(t, academic_year, student, group)
	student_to_group2, err := testQueries.GetStudentToGroupByGroupID(context.Background(), student_to_group1.GroupID)
	require.NoError(t, err)
	require.NotEmpty(t, student_to_group2)

	for i := 0; i < len(student_to_group2); i++ {
		require.Equal(t, student_to_group1.ID, student_to_group2[i].ID)
		require.Equal(t, student_to_group1.AcademicYearID, student_to_group2[i].AcademicYearID)
		require.Equal(t, student_to_group1.StudentID, student_to_group2[i].StudentID)
		require.Equal(t, student_to_group1.GroupID, student_to_group2[i].GroupID)
		require.WithinDuration(t, student_to_group1.CreatedAt, student_to_group2[i].CreatedAt, time.Second)
	}
}

func TestGetStudentToGroupBycademicYearIDAndStudentID(t *testing.T) {
	academic_year := createRandomAcademicYear(t)
	user := createRandomUser(t)
	student := createRandomStudent(t, user)
	group := createRandomGroup(t, academic_year)
	student_to_group1 := createRandomStudentToGroup(t, academic_year, student, group)
	student_to_group2, err := testQueries.GetStudentToGroupByAcademicYearIDAndStudentID(context.Background(), GetStudentToGroupByAcademicYearIDAndStudentIDParams{student_to_group1.AcademicYearID, student_to_group1.StudentID})
	require.NoError(t, err)
	require.NotEmpty(t, student_to_group2)

	for i := 0; i < len(student_to_group2); i++ {
		require.Equal(t, student_to_group1.ID, student_to_group2[i].ID)
		require.Equal(t, student_to_group1.AcademicYearID, student_to_group2[i].AcademicYearID)
		require.Equal(t, student_to_group1.StudentID, student_to_group2[i].StudentID)
		require.Equal(t, student_to_group1.GroupID, student_to_group2[i].GroupID)
		require.WithinDuration(t, student_to_group1.CreatedAt, student_to_group2[i].CreatedAt, time.Second)
	}
}
