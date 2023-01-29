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

//go:embed students_data.json
var studentSrc string

func SeedStudentData(basePath string, server *api.Server) {
	type studentData struct {
		StudentID string `json:"student_id"`
		FirstName string `json:"firstname"`
		LastName  string `json:"lastname"`
		Mobile    string `json:"mobile"`
		Email     string `json:"email"`
		Password  string `json:"password"`
		Biography string `json:"biography"`
		Image     string `json:"image"`
	}
	var students []studentData
	if err := json.Unmarshal([]byte(studentSrc), &students); err != nil {
		log.Fatal().Err(err).Msg("cannot parse student data")
	}
	// create user account
	// userByStudentID := make(map[string]db.User)
	for _, student := range students {
		url := basePath + "/student"
		recorder := httptest.NewRecorder()
		// Marshal body data to JSON
		data, err := json.Marshal(student)
		if err != nil {
			log.Fatal().Err(err).Msg("cannot parse create_student request body")
		}
		request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
		if err != nil {
			log.Fatal().Err(err).Msg("cannot init create_student request")
		}
		server.ServeHTTP(recorder, request)
		if recorder.Code != 200 {
			log.Fatal().Msg("failed to call create_student request")
		}
	}
}
