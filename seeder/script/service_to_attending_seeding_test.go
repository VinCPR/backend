package script

import (
	"testing"
)

func TestSeedServiceToAttending(t *testing.T) {
	// type serviceToAttendingData struct {
	// 	ServiceID   int64 `json:"service_id"`
	// 	AttendingID int64 `json:"attending_id"`
	// }
	// var relations []serviceToAttendingData
	//
	// numberOfAttending := 100
	// numberOfService := 25
	// for i := 1; i <= numberOfService; i++ {
	// 	n := rand.Int()%5 + 1
	// 	checked := make(map[int64]struct{})
	// 	for j := 1; j <= n; j++ {
	// 		attendingID := int64(rand.Int()%numberOfAttending + 1)
	// 		if _, ok := checked[attendingID]; ok {
	// 			j--
	// 			continue
	// 		}
	// 		checked[attendingID] = struct{}{}
	// 		relations = append(relations, serviceToAttendingData{
	// 			AttendingID: attendingID,
	// 			ServiceID:   int64(i),
	// 		})
	// 	}
	// }
	// res, err := json.Marshal(relations)
	// assert.Nil(t, err)
	// fmt.Println(string(res))
	ClearDataDBMigration("file://../../db/migration", config.DBUrl)
	SeedAttendingData(config.BasePath, server)
	SeedHospitalData(config.BasePath, server)
	SeedSpecialtyData(config.BasePath, server)
	SeedServiceData(config.BasePath, server)
	SeedServiceToAttendingData(config.BasePath, server)
}
