package db

import (
	"context"
	"testing"
	"time"

	"github.com/VinCPR/backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomHospital(t *testing.T) Hospital {
	arg := CreateHospitalParams{
		Name:        util.RandomString(6),
		Description: util.RandomString(8),
		Address:     util.RandomAddress(),
	}

	hospital, err := testQueries.CreateHospital(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, hospital)

	require.Equal(t, arg.Name, hospital.Name)
	require.Equal(t, arg.Description, hospital.Description)
	require.Equal(t, arg.Address, hospital.Address)

	require.NotZero(t, hospital.ID)
	require.NotZero(t, hospital.CreatedAt)

	return hospital
}

func TestCreateHospital(t *testing.T) {
	createRandomHospital(t)

}

func TestGetHospitalByName(t *testing.T) {
	hospital1 := createRandomHospital(t)
	hospital2, err := testQueries.GetHospitalByName(context.Background(), hospital1.Name)
	require.NoError(t, err)
	require.NotEmpty(t, hospital2)

	require.Equal(t, hospital1.ID, hospital2.ID)
	require.Equal(t, hospital1.Name, hospital2.Name)
	require.Equal(t, hospital1.Description, hospital2.Description)
	require.Equal(t, hospital1.Address, hospital2.Address)
	require.WithinDuration(t, hospital1.CreatedAt, hospital2.CreatedAt, time.Second)
}

func TestGetHospitalByID(t *testing.T) {
	hospital1 := createRandomHospital(t)
	hospital2, err := testQueries.GetHospitalByID(context.Background(), hospital1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, hospital2)

	require.Equal(t, hospital1.ID, hospital2.ID)
	require.Equal(t, hospital1.Name, hospital2.Name)
	require.Equal(t, hospital1.Description, hospital2.Description)
	require.Equal(t, hospital1.Address, hospital2.Address)
	require.WithinDuration(t, hospital1.CreatedAt, hospital2.CreatedAt, time.Second)
}

func TestListHospitals(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomHospital(t)
	}

	arg := ListHospitalsByNameParams{
		Limit:  5,
		Offset: 5,
	}

	hospitals, err := testQueries.ListHospitalsByName(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, hospitals, 5)

	for _, hospital := range hospitals {
		require.NotEmpty(t, hospital)
	}
}
