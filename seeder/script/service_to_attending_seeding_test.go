package script

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSeedServiceToAttending(t *testing.T) {
	type serviceToAttendingData struct {
		ServiceID   int64 `json:"service_id"`
		AttendingID int64 `json:"attending_id"`
	}
	var relations []serviceToAttendingData

	numberOfAttending := 100
	numberOfService := 25
	for i := 1; i <= numberOfService; i++ {
		n := rand.Int()%5 + 1
		for j := 1; j <= n; j++ {
			relations = append(relations, serviceToAttendingData{
				AttendingID: int64(rand.Int()%numberOfAttending + 1),
				ServiceID:   int64(i),
			})
		}
	}
	res, err := json.Marshal(relations)
	assert.Nil(t, err)
	fmt.Println(string(res))
}
