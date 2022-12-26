package db

import (
	"context"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomServiceToAttending(t *testing.T, service Service, attending Attending) ServiceToAttending {

	arg := CreateServiceToAttendingParams{
		ServiceID:   service.ID,
		AttendingID: attending.ID,
	}

<<<<<<< HEAD
	serviceToAttending, err := testQueries.CreateServiceToAttending(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, serviceToAttending)

	require.Equal(t, arg.ServiceID, serviceToAttending.ServiceID)
	require.Equal(t, arg.AttendingID, serviceToAttending.AttendingID)

	require.NotZero(t, serviceToAttending.ID)
	require.NotZero(t, serviceToAttending.CreatedAt)

	return serviceToAttending
=======
	service_to_attending, err := testQueries.CreateServiceToAttending(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, service_to_attending)

	require.Equal(t, arg.ServiceID, service_to_attending.ServiceID)
	require.Equal(t, arg.AttendingID, service_to_attending.AttendingID)

	require.NotZero(t, service_to_attending.ID)
	require.NotZero(t, service_to_attending.CreatedAt)

	return service_to_attending
>>>>>>> cfc0062 (add sql and test for hospital, specialty, service and service to attending)
}

// Test for CreateServiceToAttending
func TestCreateServiceToAttending(t *testing.T) {
	hospital := createRandomHospital(t)
	specialty := createRandomSpecialty(t)
	service := createRandomService(t, hospital, specialty)

	user := createRandomUser(t)
	attending := createRandomAttending(t, user)
	createRandomServiceToAttending(t, service, attending)

}

// Test for GetServiceToAttendingByServiceID
func TestGetServiceToAttendingByServiceID(t *testing.T) {
	hospital := createRandomHospital(t)
	specialty := createRandomSpecialty(t)
	service := createRandomService(t, hospital, specialty)

	user := createRandomUser(t)
	attending := createRandomAttending(t, user)

<<<<<<< HEAD
	serviceToAttending1 := createRandomServiceToAttending(t, service, attending)

	serviceToAttending2, err := testQueries.GetServiceToAttendingByServiceID(context.Background(), serviceToAttending1.ServiceID)
	require.NoError(t, err)
	require.NotEmpty(t, serviceToAttending2)

	require.Equal(t, serviceToAttending1.AttendingID, serviceToAttending2.AttendingID)
	require.Equal(t, serviceToAttending1.ServiceID, serviceToAttending2.ServiceID)
	require.WithinDuration(t, serviceToAttending1.CreatedAt, serviceToAttending2.CreatedAt, time.Second)
=======
	service_to_attending1 := createRandomServiceToAttending(t, service, attending)

	service_to_attending2, err := testQueries.GetServiceToAttendingByServiceID(context.Background(), service_to_attending1.ServiceID)
	require.NoError(t, err)
	require.NotEmpty(t, service_to_attending2)

	require.Equal(t, service_to_attending1.AttendingID, service_to_attending2.AttendingID)
	require.Equal(t, service_to_attending1.ServiceID, service_to_attending2.ServiceID)
	require.WithinDuration(t, service_to_attending1.CreatedAt, service_to_attending2.CreatedAt, time.Second)
>>>>>>> cfc0062 (add sql and test for hospital, specialty, service and service to attending)
}

// Test for GetServiceToAttendingByAttendingID
func TestGetServiceToAttendingByAttendingID(t *testing.T) {
	hospital := createRandomHospital(t)
	specialty := createRandomSpecialty(t)
	service := createRandomService(t, hospital, specialty)

	user := createRandomUser(t)
	attending := createRandomAttending(t, user)

<<<<<<< HEAD
	serviceToAttending1 := createRandomServiceToAttending(t, service, attending)

	serviceToAttending2, err := testQueries.GetServiceToAttendingByAttendingID(context.Background(), serviceToAttending1.AttendingID)
	require.NoError(t, err)
	require.NotEmpty(t, serviceToAttending2)

	require.Equal(t, serviceToAttending1.AttendingID, serviceToAttending2.AttendingID)
	require.Equal(t, serviceToAttending1.ServiceID, serviceToAttending2.ServiceID)
	require.WithinDuration(t, serviceToAttending1.CreatedAt, serviceToAttending2.CreatedAt, time.Second)
=======
	service_to_attending1 := createRandomServiceToAttending(t, service, attending)

	service_to_attending2, err := testQueries.GetServiceToAttendingByAttendingID(context.Background(), service_to_attending1.AttendingID)
	require.NoError(t, err)
	require.NotEmpty(t, service_to_attending2)

	require.Equal(t, service_to_attending1.AttendingID, service_to_attending2.AttendingID)
	require.Equal(t, service_to_attending1.ServiceID, service_to_attending2.ServiceID)
	require.WithinDuration(t, service_to_attending1.CreatedAt, service_to_attending2.CreatedAt, time.Second)
>>>>>>> cfc0062 (add sql and test for hospital, specialty, service and service to attending)
}

// Test for ListServicesToAttendingsByServiceID
func TestListServicesToAttendingsByServiceID(t *testing.T) {
	for i := 0; i < 10; i++ {
		hospital := createRandomHospital(t)
		specialty := createRandomSpecialty(t)
		service := createRandomService(t, hospital, specialty)

		user := createRandomUser(t)
		attending := createRandomAttending(t, user)
		createRandomServiceToAttending(t, service, attending)
	}
<<<<<<< HEAD

	allServiceToAttending, err := testQueries.ListServicesToAttendingsByServiceID(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, allServiceToAttending)
	require.GreaterOrEqual(t, len(allServiceToAttending), 10)

	sort.SliceIsSorted(allServiceToAttending, func(i, j int) bool {
		return allServiceToAttending[i].ServiceID < allServiceToAttending[j].ServiceID
	})

	for _, serviceToAttending := range allServiceToAttending {
		require.NotEmpty(t, serviceToAttending)
=======
	hospital := createRandomHospital(t)
	specialty := createRandomSpecialty(t)
	service := createRandomService(t, hospital, specialty)

	user := createRandomUser(t)
	attending := createRandomAttending(t, user)
	lastServiceToAttending := createRandomServiceToAttending(t, service, attending)
	var numOfServiceToAttendings = lastServiceToAttending.ID

	arg2 := ListServicesToAttendingsByServiceIDParams{
		Limit:  int32(numOfServiceToAttendings),
		Offset: 5,
	}

	service_to_attending1, err := testQueries.ListServicesToAttendingsByServiceID(context.Background(), arg2)
	require.NoError(t, err)

	var listServicesToAttendingsbyServiceID []int64
	for _, service_to_attending := range service_to_attending1 {
		listServicesToAttendingsbyServiceID = append(listServicesToAttendingsbyServiceID, service_to_attending.ServiceID)
	}
	sort.Slice(listServicesToAttendingsbyServiceID, func(i, j int) bool {
		return listServicesToAttendingsbyServiceID[i] < listServicesToAttendingsbyServiceID[j]
	})

	arg := ListServicesToAttendingsByServiceIDParams{

		Limit:  5,
		Offset: 5,
	}

	service_to_attendings, err := testQueries.ListServicesToAttendingsByServiceID(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, service_to_attendings, 5)
	var i = 0
	for _, service_to_attending := range service_to_attendings {
		require.NotEmpty(t, service_to_attendings)
		require.Equal(t, service_to_attending.ServiceID, listServicesToAttendingsbyServiceID[i])
		i++
>>>>>>> cfc0062 (add sql and test for hospital, specialty, service and service to attending)
	}
}

// Test for ListServicesToAttendingsByAttendingID
func TestListServicesToAttendingsByAttendingID(t *testing.T) {
	for i := 0; i < 10; i++ {
		hospital := createRandomHospital(t)
		specialty := createRandomSpecialty(t)
		service := createRandomService(t, hospital, specialty)

		user := createRandomUser(t)
		attending := createRandomAttending(t, user)
		createRandomServiceToAttending(t, service, attending)
	}
<<<<<<< HEAD

	allServiceToAttending, err := testQueries.ListServicesToAttendingsByAttendingID(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, allServiceToAttending)
	require.GreaterOrEqual(t, len(allServiceToAttending), 10)

	sort.SliceIsSorted(allServiceToAttending, func(i, j int) bool {
		return allServiceToAttending[i].AttendingID < allServiceToAttending[j].AttendingID
	})

	for _, serviceToAttending := range allServiceToAttending {
		require.NotEmpty(t, serviceToAttending)
=======
	hospital := createRandomHospital(t)
	specialty := createRandomSpecialty(t)
	service := createRandomService(t, hospital, specialty)

	user := createRandomUser(t)
	attending := createRandomAttending(t, user)
	lastServiceToAttending := createRandomServiceToAttending(t, service, attending)
	var numOfServiceToAttendings = lastServiceToAttending.ID

	arg2 := ListServicesToAttendingsByAttendingIDParams{
		Limit:  int32(numOfServiceToAttendings),
		Offset: 5,
	}

	service_to_attending1, err := testQueries.ListServicesToAttendingsByAttendingID(context.Background(), arg2)
	require.NoError(t, err)

	var listServicesToAttendingsbyAttendingID []int64
	for _, service_to_attending := range service_to_attending1 {
		listServicesToAttendingsbyAttendingID = append(listServicesToAttendingsbyAttendingID, service_to_attending.AttendingID)
	}
	sort.Slice(listServicesToAttendingsbyAttendingID, func(i, j int) bool {
		return listServicesToAttendingsbyAttendingID[i] < listServicesToAttendingsbyAttendingID[j]
	})

	arg := ListServicesToAttendingsByAttendingIDParams{

		Limit:  5,
		Offset: 5,
	}

	service_to_attendings, err := testQueries.ListServicesToAttendingsByAttendingID(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, service_to_attendings, 5)
	var i = 0
	for _, service_to_attending := range service_to_attendings {
		require.NotEmpty(t, service_to_attendings)
		require.Equal(t, service_to_attending.AttendingID, listServicesToAttendingsbyAttendingID[i])
		i++
>>>>>>> cfc0062 (add sql and test for hospital, specialty, service and service to attending)
	}
}

// Test for ListServicesToAttendingsByAll
func TestListServicesToAttendingsByAll(t *testing.T) {
	for i := 0; i < 10; i++ {
		hospital := createRandomHospital(t)
		specialty := createRandomSpecialty(t)
		service := createRandomService(t, hospital, specialty)

		user := createRandomUser(t)
		attending := createRandomAttending(t, user)
		createRandomServiceToAttending(t, service, attending)
	}
<<<<<<< HEAD

	allServiceToAttending, err := testQueries.ListServicesToAttendingsByServiceIDAndAttendingID(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, allServiceToAttending)
	require.GreaterOrEqual(t, len(allServiceToAttending), 10)

	sort.SliceIsSorted(allServiceToAttending, func(i, j int) bool {
		return allServiceToAttending[i].ServiceID < allServiceToAttending[j].ServiceID ||
			(allServiceToAttending[i].ServiceID == allServiceToAttending[j].ServiceID && allServiceToAttending[i].AttendingID < allServiceToAttending[j].AttendingID)

	})

	for _, serviceToAttending := range allServiceToAttending {
		require.NotEmpty(t, serviceToAttending)
=======
	hospital := createRandomHospital(t)
	specialty := createRandomSpecialty(t)
	service := createRandomService(t, hospital, specialty)

	user := createRandomUser(t)
	attending := createRandomAttending(t, user)
	lastServiceToAttending := createRandomServiceToAttending(t, service, attending)
	var numOfServiceToAttendings = lastServiceToAttending.ID

	arg2 := ListServicesToAttendingsByAllParams{
		Limit:  int32(numOfServiceToAttendings),
		Offset: 5,
	}

	listServicesToAttendingsbyAll, err := testQueries.ListServicesToAttendingsByAll(context.Background(), arg2)
	require.NoError(t, err)

	sort.Slice(listServicesToAttendingsbyAll, func(i, j int) bool {
		return listServicesToAttendingsbyAll[i].ServiceID < listServicesToAttendingsbyAll[j].ServiceID ||
			(listServicesToAttendingsbyAll[i].ServiceID == listServicesToAttendingsbyAll[j].ServiceID && listServicesToAttendingsbyAll[i].AttendingID < listServicesToAttendingsbyAll[j].AttendingID)
	})
	arg := ListServicesToAttendingsByAllParams{

		Limit:  5,
		Offset: 5,
	}

	service_to_attendings, err := testQueries.ListServicesToAttendingsByAll(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, service_to_attendings, 5)
	var i = 0
	for _, service_to_attending := range service_to_attendings {
		require.NotEmpty(t, service_to_attendings)
		require.Equal(t, service_to_attending.ServiceID, listServicesToAttendingsbyAll[i].ServiceID)
		require.Equal(t, service_to_attending.AttendingID, listServicesToAttendingsbyAll[i].AttendingID)
		i++
>>>>>>> cfc0062 (add sql and test for hospital, specialty, service and service to attending)
	}
}
