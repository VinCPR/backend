package db

import (
	"context"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/VinCPR/backend/util"
)

func createRandomStudent(t *testing.T, user User) Student {
	arg := CreateStudentParams{
		UserID:    user.ID,
		StudentID: util.RandomStudentID(),
		FirstName: util.RandomName(),
		LastName:  util.RandomName(),
		Mobile:    util.RandomMobile(),
	}

	student, err := testQueries.CreateStudent(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, student)

	require.Equal(t, arg.UserID, student.UserID)
	require.Equal(t, arg.StudentID, student.StudentID)
	require.Equal(t, arg.FirstName, student.FirstName)
	require.Equal(t, arg.LastName, student.LastName)
	require.Equal(t, arg.Mobile, student.Mobile)

	require.NotZero(t, student.ID)
	require.NotZero(t, student.CreatedAt)

	return student
}

func TestCreateStudent(t *testing.T) {
	user := createRandomUser(t)
	createRandomStudent(t, user)
}

func TestGetStudentByStudentId(t *testing.T) {
	user := createRandomUser(t)
	student1 := createRandomStudent(t, user)
	student2, err := testQueries.GetStudentByUserId(context.Background(), user.ID)
	require.NoError(t, err)
	require.NotEmpty(t, student2)

	require.Equal(t, student1.ID, student2.ID)
	require.Equal(t, student1.UserID, student2.UserID)
	require.Equal(t, student1.StudentID, student2.StudentID)
	require.Equal(t, student1.FirstName, student2.FirstName)
	require.Equal(t, student1.LastName, student2.LastName)
	require.Equal(t, student1.Mobile, student2.Mobile)
	require.WithinDuration(t, student1.CreatedAt, student2.CreatedAt, time.Second)
}

func TestGetStudentByUserId(t *testing.T) {
	user := createRandomUser(t)
	student1 := createRandomStudent(t, user)
	student2, err := testQueries.GetStudentByStudentId(context.Background(), student1.StudentID)
	require.NoError(t, err)
	require.NotEmpty(t, student2)

	require.Equal(t, student1.ID, student2.ID)
	require.Equal(t, student1.UserID, student2.UserID)
	require.Equal(t, student1.StudentID, student2.StudentID)
	require.Equal(t, student1.FirstName, student2.FirstName)
	require.Equal(t, student1.LastName, student2.LastName)
	require.Equal(t, student1.Mobile, student2.Mobile)
	require.WithinDuration(t, student1.CreatedAt, student2.CreatedAt, time.Second)
}

func TestListStudentsByName(t *testing.T) {
	var n = 5
	for i := 0; i < n; i++ {
		user := createRandomUser(t)
		createRandomStudent(t, user)

	}
	allStudents, err := testQueries.ListStudentsByName(context.Background(), ListStudentsByNameParams{
		Limit:  100,
		Offset: 0,
	})
	require.NoError(t, err)

	sort.SliceIsSorted(allStudents, func(i, j int) bool {
		return allStudents[i].FirstName < allStudents[j].FirstName ||
			(allStudents[i].FirstName == allStudents[j].FirstName && allStudents[i].LastName < allStudents[j].LastName)
	})

	students, err := testQueries.ListStudentsByName(context.Background(), ListStudentsByNameParams{
		Limit:  int32(n),
		Offset: 0,
	})
	require.NoError(t, err)

	require.Len(t, students, n)
	require.EqualValues(t, students, allStudents[:n])
}

func TestListStudentsByStudentID(t *testing.T) {
	var n = 5
	for i := 0; i < n; i++ {
		user := createRandomUser(t)
		createRandomStudent(t, user)

	}
	allStudents, err := testQueries.ListStudentsByStudentID(context.Background(), ListStudentsByStudentIDParams{
		Limit:  100,
		Offset: 0,
	})
	require.NoError(t, err)

	sort.SliceIsSorted(allStudents, func(i, j int) bool {
		return allStudents[i].FirstName < allStudents[j].FirstName ||
			(allStudents[i].FirstName == allStudents[j].FirstName && allStudents[i].LastName < allStudents[j].LastName)
	})

	students, err := testQueries.ListStudentsByStudentID(context.Background(), ListStudentsByStudentIDParams{
		Limit:  int32(n),
		Offset: 0,
	})
	require.NoError(t, err)

	require.Len(t, students, n)
	require.EqualValues(t, students, allStudents[:n])
}
