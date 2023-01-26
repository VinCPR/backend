package script

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/rs/zerolog/log"

	"github.com/VinCPR/backend/api"
)

//go:embed student_to_group_data.json
var studentToGroupSrc string

func SeedStudentToGroupData(basePath string, server *api.Server) {
	type studentToGroupData struct {
		AcademicYearName string `json:"academic_year_name"`
		StudentID        string `json:"student_id"`
		GroupName        string `json:"group_name"`
	}
	var relations []studentToGroupData
	if err := json.Unmarshal([]byte(studentToGroupSrc), &relations); err != nil {
		log.Fatal().Err(err).Msg("cannot parse service to attending data")
	}
	for _, relation := range relations {
		url := basePath + "/student_to_group"
		recorder := httptest.NewRecorder()
		// Marshal body data to JSON
		data, err := json.Marshal(relation)
		if err != nil {
			log.Fatal().Err(err).Msg("cannot parse create_student_to_group request body")
		}
		request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
		if err != nil {
			log.Fatal().Err(err).Msg("cannot init create_student_to_group request")
		}
		server.ServeHTTP(recorder, request)
		fmt.Println(recorder)
		if recorder.Code != 200 {
			log.Fatal().Msg("failed to call create_student_to_group request")
		}
	}
}
