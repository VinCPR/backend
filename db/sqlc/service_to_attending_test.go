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

	serviceToAttending, err := testQueries.CreateServiceToAttending(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, serviceToAttending)

	require.Equal(t, arg.ServiceID, serviceToAttending.ServiceID)
	require.Equal(t, arg.AttendingID, serviceToAttending.AttendingID)

	require.NotZero(t, serviceToAttending.ID)
	require.NotZero(t, serviceToAttending.CreatedAt)

	return serviceToAttending
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

	serviceToAttending1 := createRandomServiceToAttending(t, service, attending)

	serviceToAttending2, err := testQueries.GetServiceToAttendingByServiceID(context.Background(), serviceToAttending1.ServiceID)
	require.NoError(t, err)
	require.NotEmpty(t, serviceToAttending2)

	require.Equal(t, serviceToAttending1.AttendingID, serviceToAttending2.AttendingID)
	require.Equal(t, serviceToAttending1.ServiceID, serviceToAttending2.ServiceID)
	require.WithinDuration(t, serviceToAttending1.CreatedAt, serviceToAttending2.CreatedAt, time.Second)
}

// Test for GetServiceToAttendingByAttendingID
func TestGetServiceToAttendingByAttendingID(t *testing.T) {
	hospital := createRandomHospital(t)
	specialty := createRandomSpecialty(t)
	service := createRandomService(t, hospital, specialty)

	user := createRandomUser(t)
	attending := createRandomAttending(t, user)

	serviceToAttending1 := createRandomServiceToAttending(t, service, attending)

	serviceToAttending2, err := testQueries.GetServiceToAttendingByAttendingID(context.Background(), serviceToAttending1.AttendingID)
	require.NoError(t, err)
	require.NotEmpty(t, serviceToAttending2)

	require.Equal(t, serviceToAttending1.AttendingID, serviceToAttending2.AttendingID)
	require.Equal(t, serviceToAttending1.ServiceID, serviceToAttending2.ServiceID)
	require.WithinDuration(t, serviceToAttending1.CreatedAt, serviceToAttending2.CreatedAt, time.Second)
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

	allServiceToAttending, err := testQueries.ListServicesToAttendingsByServiceID(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, allServiceToAttending)
	require.GreaterOrEqual(t, len(allServiceToAttending), 10)

	sort.SliceIsSorted(allServiceToAttending, func(i, j int) bool {
		return allServiceToAttending[i].ServiceID < allServiceToAttending[j].ServiceID
	})

	for _, serviceToAttending := range allServiceToAttending {
		require.NotEmpty(t, serviceToAttending)
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

	allServiceToAttending, err := testQueries.ListServicesToAttendingsByAttendingID(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, allServiceToAttending)
	require.GreaterOrEqual(t, len(allServiceToAttending), 10)

	sort.SliceIsSorted(allServiceToAttending, func(i, j int) bool {
		return allServiceToAttending[i].AttendingID < allServiceToAttending[j].AttendingID
	})

	for _, serviceToAttending := range allServiceToAttending {
		require.NotEmpty(t, serviceToAttending)
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
	}
}
