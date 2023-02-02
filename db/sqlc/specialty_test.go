package db

import (
	"context"
	"testing"
	"time"

	"github.com/VinCPR/backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomSpecialty(t *testing.T) Specialty {
	arg := CreateSpecialtyParams{
		Name:        util.RandomString(6),
		Description: util.RandomString(8),
	}

	specialty, err := testQueries.CreateSpecialty(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, specialty)

	require.Equal(t, arg.Name, specialty.Name)
	require.Equal(t, arg.Description, specialty.Description)

	require.NotZero(t, specialty.ID)
	require.NotZero(t, specialty.CreatedAt)

	return specialty
}

func TestCreateSpecialty(t *testing.T) {
	createRandomSpecialty(t)

}

func TestGetSpecialtyByID(t *testing.T) {
	specialty1 := createRandomSpecialty(t)
	specialty2, err := testQueries.GetSpecialtyByID(context.Background(), specialty1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, specialty2)

	require.Equal(t, specialty1.ID, specialty2.ID)
	require.Equal(t, specialty1.Name, specialty2.Name)
	require.Equal(t, specialty1.Description, specialty2.Description)
	require.WithinDuration(t, specialty1.CreatedAt, specialty2.CreatedAt, time.Second)
}

func TestGetSpecialtyByName(t *testing.T) {
	specialty1 := createRandomSpecialty(t)
	specialty2, err := testQueries.GetSpecialtyByName(context.Background(), specialty1.Name)
	require.NoError(t, err)
	require.NotEmpty(t, specialty2)

	require.Equal(t, specialty1.ID, specialty2.ID)
	require.Equal(t, specialty1.Name, specialty2.Name)
	require.Equal(t, specialty1.Description, specialty2.Description)
	require.WithinDuration(t, specialty1.CreatedAt, specialty2.CreatedAt, time.Second)
}

func TestListSpecialtiess(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomSpecialty(t)
	}

	arg := ListSpecialtiesByNameParams{
		Limit:  5,
		Offset: 5,
	}

	specialties, err := testQueries.ListSpecialtiesByName(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, specialties, 5)

	for _, specialty := range specialties {
		require.NotEmpty(t, specialty)
	}
}
