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

//go:embed specialties_data.json
var specialtySrc string

func SeedSpecialtyData(basePath string, server *api.Server) {
	type specialtyData struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	var specialties []specialtyData
	if err := json.Unmarshal([]byte(specialtySrc), &specialties); err != nil {
		log.Fatal().Err(err).Msg("cannot parse specialty data")
	}
	for _, specialty := range specialties {
		url := basePath + "/specialty"
		recorder := httptest.NewRecorder()
		// Marshal body data to JSON
		data, err := json.Marshal(specialty)
		if err != nil {
			log.Fatal().Err(err).Msg("cannot parse create_specialty request body")
		}
		request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
		if err != nil {
			log.Fatal().Err(err).Msg("cannot init create_specialty request")
		}
		server.ServeHTTP(recorder, request)
		if recorder.Code != 200 {
			log.Fatal().Msg("failed to call create_specialty request")
		}
	}
}
