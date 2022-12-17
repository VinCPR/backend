package db

import (
	"context"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/VinCPR/backend/util"
)

func createRandomService(t *testing.T, hospital Hospital, specialty Specialty) Service {

	arg := CreateServiceParams{
		SpecialtyID: specialty.ID,
		HospitalID:  hospital.ID,
		Name:        util.RandomString(6),
		Description: util.RandomString(8),
	}

	service, err := testQueries.CreateService(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, service)

	require.Equal(t, arg.Name, service.Name)
	require.Equal(t, arg.Description, service.Description)
	require.Equal(t, specialty.ID, service.SpecialtyID)
	require.Equal(t, hospital.ID, service.HospitalID)

	require.NotZero(t, service.ID)
	require.NotZero(t, service.CreatedAt)

	return service
}

func TestCreateService(t *testing.T) {
	hospital := createRandomHospital(t)
	specialty := createRandomSpecialty(t)
	createRandomService(t, hospital, specialty)

}

func TestGetService(t *testing.T) {
	hospital := createRandomHospital(t)
	specialty := createRandomSpecialty(t)
	service1 := createRandomService(t, hospital, specialty)
	service2, err := testQueries.GetServiceByName(context.Background(), service1.Name)
	require.NoError(t, err)
	require.NotEmpty(t, service2)

	require.Equal(t, service1.SpecialtyID, service2.SpecialtyID)
	require.Equal(t, service1.HospitalID, service2.HospitalID)
	require.Equal(t, service1.Name, service2.Name)
	require.Equal(t, service1.Description, service2.Description)
	require.WithinDuration(t, service1.CreatedAt, service2.CreatedAt, time.Second)
}

// TODO refactor test for list queries

func TestListServicesBySpecialID(t *testing.T) {
	for i := 0; i < 10; i++ {
		hospital := createRandomHospital(t)
		specialty := createRandomSpecialty(t)
		createRandomService(t, hospital, specialty)
	}
	hospital := createRandomHospital(t)
	specialty := createRandomSpecialty(t)
	lastService := createRandomService(t, hospital, specialty)
	var numOfServices = lastService.ID

	arg2 := ListServicesBySpecialtyIDParams{
		Limit:  int32(numOfServices),
		Offset: 5,
	}

	service1, err := testQueries.ListServicesBySpecialtyID(context.Background(), arg2)
	require.NoError(t, err)

	var listServicesbySpecialtyID []int64
	for _, service := range service1 {
		listServicesbySpecialtyID = append(listServicesbySpecialtyID, service.SpecialtyID)
	}
	sort.Slice(listServicesbySpecialtyID, func(i, j int) bool { return listServicesbySpecialtyID[i] < listServicesbySpecialtyID[j] })

	arg := ListServicesBySpecialtyIDParams{

		Limit:  5,
		Offset: 5,
	}

	services, err := testQueries.ListServicesBySpecialtyID(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, services, 5)
	var i = 0
	for _, service := range services {
		require.NotEmpty(t, service)
		require.Equal(t, service.SpecialtyID, listServicesbySpecialtyID[i])
		i++
	}
}

func TestListServicesByHospitalID(t *testing.T) {
	for i := 0; i < 10; i++ {
		hospital := createRandomHospital(t)
		specialty := createRandomSpecialty(t)
		createRandomService(t, hospital, specialty)
	}
	hospital := createRandomHospital(t)
	specialty := createRandomSpecialty(t)
	lastService := createRandomService(t, hospital, specialty)
	var numOfServices = lastService.ID

	arg2 := ListServicesByHospitalIDParams{
		Limit:  int32(numOfServices),
		Offset: 5,
	}

	service1, err := testQueries.ListServicesByHospitalID(context.Background(), arg2)
	require.NoError(t, err)

	var listServicesbyHospitalID []int64
	for _, service := range service1 {
		listServicesbyHospitalID = append(listServicesbyHospitalID, service.HospitalID)
	}
	sort.Slice(listServicesbyHospitalID, func(i, j int) bool { return listServicesbyHospitalID[i] < listServicesbyHospitalID[j] })

	arg := ListServicesByHospitalIDParams{

		Limit:  5,
		Offset: 5,
	}

	services, err := testQueries.ListServicesByHospitalID(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, services, 5)
	var i = 0
	for _, service := range services {
		require.NotEmpty(t, service)
		require.Equal(t, service.HospitalID, listServicesbyHospitalID[i])
		i++
	}
}

func TestListServicesBySpecialtyIDAndHospitalID(t *testing.T) {
	for i := 0; i < 10; i++ {
		hospital := createRandomHospital(t)
		specialty := createRandomSpecialty(t)
		createRandomService(t, hospital, specialty)
	}
	hospital := createRandomHospital(t)
	specialty := createRandomSpecialty(t)
	lastService := createRandomService(t, hospital, specialty)
	var numOfServices = lastService.ID

	arg2 := ListServicesBySpecialtyIDAndHospitalIDParams{
		Limit:  int32(numOfServices),
		Offset: 5,
	}

	listServicesbySpecialtyIDAndHospitalID, err := testQueries.ListServicesBySpecialtyIDAndHospitalID(context.Background(), arg2)
	require.NoError(t, err)

	sort.Slice(listServicesbySpecialtyIDAndHospitalID, func(i, j int) bool {
		return listServicesbySpecialtyIDAndHospitalID[i].SpecialtyID < listServicesbySpecialtyIDAndHospitalID[j].SpecialtyID ||
			(listServicesbySpecialtyIDAndHospitalID[i].SpecialtyID == listServicesbySpecialtyIDAndHospitalID[j].SpecialtyID && listServicesbySpecialtyIDAndHospitalID[i].HospitalID < listServicesbySpecialtyIDAndHospitalID[j].HospitalID)
	})
	arg := ListServicesBySpecialtyIDAndHospitalIDParams{

		Limit:  5,
		Offset: 5,
	}

	services, err := testQueries.ListServicesBySpecialtyIDAndHospitalID(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, services, 5)
	var i = 0
	for _, service := range services {
		require.NotEmpty(t, service)
		require.Equal(t, service.SpecialtyID, listServicesbySpecialtyIDAndHospitalID[i].SpecialtyID)
		require.Equal(t, service.HospitalID, listServicesbySpecialtyIDAndHospitalID[i].HospitalID)
		i++
	}
}
