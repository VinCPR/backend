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

//go:embed attendings_data.json
var attendingSrc string

func SeedAttendingData(basePath string, server *api.Server) {
	type attendingData struct {
		AttendingID string `json:"attending_id"`
		FirstName   string `json:"firstname"`
		LastName    string `json:"lastname"`
		Mobile      string `json:"mobile"`
		Email       string `json:"email"`
		Password    string `json:"password"`
	}
	var attendings []attendingData
	if err := json.Unmarshal([]byte(attendingSrc), &attendings); err != nil {
		log.Fatal().Err(err).Msg("cannot parse attending data")
	}
	// create user account
	// userByAttendingID := make(map[string]db.User)
	for _, attending := range attendings {
		url := basePath + "/attending"
		recorder := httptest.NewRecorder()
		// Marshal body data to JSON
		data, err := json.Marshal(attending)
		if err != nil {
			log.Fatal().Err(err).Msg("cannot parse create_attending request body")
		}
		request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
		if err != nil {
			log.Fatal().Err(err).Msg("cannot init create_attending request")
		}
		server.ServeHTTP(recorder, request)
		if recorder.Code != 200 {
			log.Fatal().Msg("failed to call create_user request")
		}
	}
}
