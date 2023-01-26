package script

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/rs/zerolog/log"

	"github.com/VinCPR/backend/api"
)

func SeedAcademicYearData(basePath string, server *api.Server) {
	type academicYearData struct {
		Name      string `json:"name"`
		StartDate string `json:"start_date"`
		EndDate   string `json:"end_date"`
	}
	var academicYears = []academicYearData{
		{
			Name:      "2023-2024 MD Program",
			StartDate: "2023-09-30T00:00:00+07:00",
			EndDate:   "2024-06-30T00:00:00+07:00",
		},
		{
			Name:      "2023-2024 Nursing Program",
			StartDate: "2023-09-30T00:00:00+07:00",
			EndDate:   "2024-06-30T00:00:00+07:00",
		},
	}
	for _, academicYear := range academicYears {
		url := basePath + "/academic_year"
		recorder := httptest.NewRecorder()
		// Marshal body data to JSON
		data, err := json.Marshal(academicYear)
		if err != nil {
			log.Fatal().Err(err).Msg("cannot parse create_academic_year request body")
		}
		request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
		if err != nil {
			log.Fatal().Err(err).Msg("cannot init create_academic_year request")
		}
		server.ServeHTTP(recorder, request)
		if recorder.Code != 200 {
			log.Fatal().Msg("failed to call create_academic_year request")
		}
	}
}
