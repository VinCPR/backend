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

//go:embed services_data.json
var serviceSrc string

func SeedServiceData(basePath string, server *api.Server) {
	type serviceData struct {
		Hospital    string `json:"hospital"`
		Specialty   string `json:"specialty"`
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	var services []serviceData
	if err := json.Unmarshal([]byte(serviceSrc), &services); err != nil {
		log.Fatal().Err(err).Msg("cannot parse service data")
	}
	for _, service := range services {
		url := basePath + "/service"
		recorder := httptest.NewRecorder()
		// Marshal body data to JSON
		data, err := json.Marshal(service)
		if err != nil {
			log.Fatal().Err(err).Msg("cannot parse create_service request body")
		}
		request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
		if err != nil {
			log.Fatal().Err(err).Msg("cannot init create_service request")
		}
		server.ServeHTTP(recorder, request)
		if recorder.Code != 200 {
			log.Fatal().Msg("failed to call create_service request")
		}
	}
}
