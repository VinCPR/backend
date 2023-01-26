package script

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/rs/zerolog/log"

	"github.com/VinCPR/backend/api"
)

//go:embed service_to_attending_data.json
var serviceToAttendingSrc string

func SeedServiceToAttendingData(basePath string, server *api.Server) {
	type serviceToAttendingData struct {
		ServiceID   int64 `json:"service_id"`
		AttendingID int64 `json:"attending_id"`
	}
	var relations []serviceToAttendingData
	if err := json.Unmarshal([]byte(serviceToAttendingSrc), &relations); err != nil {
		log.Fatal().Err(err).Msg("cannot parse service to attending data")
	}
	for _, relation := range relations {
		url := basePath + "/service_to_attending"
		recorder := httptest.NewRecorder()
		// Marshal body data to JSON
		data, err := json.Marshal(relation)
		if err != nil {
			log.Fatal().Err(err).Msg("cannot parse create_service_to_attending request body")
		}
		request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
		if err != nil {
			log.Fatal().Err(err).Msg("cannot init create_service_to_attending request")
		}
		server.ServeHTTP(recorder, request)
		if recorder.Code != 200 {
			log.Fatal().Msg("failed to call create_service_to_attending request")
		}
	}
}
