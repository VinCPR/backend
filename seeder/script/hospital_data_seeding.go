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

//go:embed hospitals_data.json
var hospitalSrc string

func SeedHospitalData(basePath string, server *api.Server) {
	type hospitalData struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description" binding:"required"`
		Address     string `json:"address" binding:"required"`
	}
	var hospitals []hospitalData
	if err := json.Unmarshal([]byte(hospitalSrc), &hospitals); err != nil {
		log.Fatal().Err(err).Msg("cannot parse hospital data")
	}
	for _, hospital := range hospitals {
		url := basePath + "/hospital"
		recorder := httptest.NewRecorder()
		// Marshal body data to JSON
		data, err := json.Marshal(hospital)
		if err != nil {
			log.Fatal().Err(err).Msg("cannot parse create_hospital request body")
		}
		request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
		if err != nil {
			log.Fatal().Err(err).Msg("cannot init create_hospital request")
		}
		server.ServeHTTP(recorder, request)
		if recorder.Code != 200 {
			log.Fatal().Msg("failed to call create_user request")
		}
	}
}
