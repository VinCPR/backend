package script

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/rs/zerolog/log"

	"github.com/VinCPR/backend/api"
)

func SeedGroupData(basePath string, server *api.Server) {
	type groupData struct {
		AcademicYearName string `json:"academic_year_name"`
		Name             string `json:"name"`
	}
	var groups = []groupData{
		{
			AcademicYearName: "2023-2024 MD Program",
			Name:             "Medical Doctor Group 1",
		},
		{
			AcademicYearName: "2023-2024 MD Program",
			Name:             "Medical Doctor Group 2",
		},
		{
			AcademicYearName: "2023-2024 MD Program",
			Name:             "Medical Doctor Group 3",
		},
		{
			AcademicYearName: "2023-2024 MD Program",
			Name:             "Medical Doctor Group 4",
		},
		{
			AcademicYearName: "2023-2024 MD Program",
			Name:             "Medical Doctor Group 5",
		},
		{
			AcademicYearName: "2023-2024 MD Program",
			Name:             "Medical Doctor Group 6",
		},
		{
			AcademicYearName: "2023-2024 MD Program",
			Name:             "Medical Doctor Group 7",
		},
		{
			AcademicYearName: "2023-2024 MD Program",
			Name:             "Medical Doctor Group 8",
		},

		{
			AcademicYearName: "2023-2024 Nursing Program",
			Name:             "Nursing Group 1",
		},
		{
			AcademicYearName: "2023-2024 Nursing Program",
			Name:             "Nursing Group 2",
		},
		{
			AcademicYearName: "2023-2024 Nursing Program",
			Name:             "Nursing Group 3",
		},
		{
			AcademicYearName: "2023-2024 Nursing Program",
			Name:             "Nursing Group 4",
		},
		{
			AcademicYearName: "2023-2024 Nursing Program",
			Name:             "Nursing Group 5",
		},
	}
	for _, group := range groups {
		url := basePath + "/group"
		recorder := httptest.NewRecorder()
		// Marshal body data to JSON
		data, err := json.Marshal(group)
		if err != nil {
			log.Fatal().Err(err).Msg("cannot parse create_group request body")
		}
		request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
		if err != nil {
			log.Fatal().Err(err).Msg("cannot init create_group request")
		}
		server.ServeHTTP(recorder, request)
		if recorder.Code != 200 {
			log.Fatal().Msg("failed to call create_group request")
		}
	}
}
