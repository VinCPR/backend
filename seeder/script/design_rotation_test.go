package script

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rs/zerolog/log"
)

//go:embed design_rotation_data.json
var designRotationRequestSrc string

func TestDesignRotation(t *testing.T) {
	ClearDataDBMigration("file://../../db/migration", config.DBUrl)
	SeedStudentData(config.BasePath, server)
	SeedAttendingData(config.BasePath, server)
	SeedHospitalData(config.BasePath, server)
	SeedSpecialtyData(config.BasePath, server)
	SeedServiceData(config.BasePath, server)
	SeedServiceToAttendingData(config.BasePath, server)
	SeedAcademicYearData(config.BasePath, server)
	SeedGroupData(config.BasePath, server)
	SeedStudentToGroupData(config.BasePath, server)

	type serviceInfoRequest struct {
		ServiceName    string `json:"service_name"`
		DurationInWeek int    `json:"duration_in_week"`
	}

	type hospitalInfoRequest struct {
		HospitalName string               `json:"hospital_name"`
		Services     []serviceInfoRequest `json:"services"`
	}

	type specialtyInfoRequest struct {
		SpecialtyName string                `json:"specialty_name"`
		Hospitals     []hospitalInfoRequest `json:"hospitals"`
	}

	type blockInfoRequest struct {
		BlockName     string                   `json:"block_name"`
		GroupCalendar [][]specialtyInfoRequest `json:"group_calendar"`
	}

	type periodInfoRequest struct {
		PeriodName string `json:"period_name" binding:"required"`
		StartDate  string `json:"start_date" binding:"required"`
		// EndDate = StartDate + WeeksPerPeriod
		// EndDate    time.Time `json:"end_date" form:"end_date" binding:"required" time_format:"2006-01-02"`
	}

	type designRotationRequest struct {
		AcademicYearName string `json:"academic_year_name" binding:"required"`
		GroupsPerBlock   int    `json:"groups_per_block" binding:"required"`
		NumberOfPeriod   int    `json:"number_of_period" binding:"required"`
		WeeksPerPeriod   int    `json:"weeks_per_period" binding:"required"`

		Periods []periodInfoRequest `json:"periods" binding:"required"`
		Blocks  []blockInfoRequest  `json:"blocks" binding:"required"`
	}

	var requests []designRotationRequest
	if err := json.Unmarshal([]byte(designRotationRequestSrc), &requests); err != nil {
		log.Fatal().Err(err).Msg("cannot parse design rotation request data")
	}
	for _, request := range requests {
		url := config.BasePath + "/rotation/design"
		recorder := httptest.NewRecorder()
		// Marshal body data to JSON
		data, err := json.Marshal(request)
		if err != nil {
			log.Fatal().Err(err).Msg("cannot parse design_rotation request body")
		}
		httpRequest, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
		if err != nil {
			log.Fatal().Err(err).Msg("cannot init design_rotation request")
		}
		server.ServeHTTP(recorder, httpRequest)
		fmt.Println(recorder)
		if recorder.Code != 200 {
			log.Fatal().Msg("failed to call design_rotation request")
		}
	}
}
