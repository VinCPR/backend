package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/rs/zerolog/log"

	db "github.com/VinCPR/backend/db/sqlc"
)

type clinicalRotationEventResponse struct {
	EventId       int64     `json:"event_id"`
	GroupName     string    `json:"group_name"`
	SpecialtyName string    `json:"specialty_name"`
	HospitalName  string    `json:"hospital_name"`
	ServiceName   string    `json:"service_name"`
	StartDate     time.Time `json:"start_date"`
	EndDate       time.Time `json:"end_date"`
}

// studentCalendar
// @Summary list calendar events for a student in an academic year
// @Description list calendar events for a student in an academic year
// @Tags ClinicalRotation
// @Accept	json
// @Produce  json
// @Param academicYearName query string true "academic year name"
// @Param studentID query string true "student ID"
// @Success 200 {object} []clinicalRotationEventResponse "ok"
// @Router /rotation/student [get]
func (server *Server) studentCalendar(ctx *gin.Context) {
	academicYearName := ctx.Query("academicYearName")
	academicYear, err := server.store.GetAcademicYearByName(ctx, academicYearName)
	if err != nil {
		if err == pgx.ErrNoRows {
			log.Info().Msgf("cannot find academic year %v", academicYearName)
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	studentID := ctx.Query("studentID")
	student, err := server.store.GetStudentByStudentId(ctx, studentID)
	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	studentToGroups, err := server.store.GetStudentToGroupByAcademicYearIDAndStudentID(ctx,
		db.GetStudentToGroupByAcademicYearIDAndStudentIDParams{
			AcademicYearID: academicYear.ID,
			StudentID:      student.ID,
		})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	groupIDs := make([]int64, 0)
	for _, studentToGroup := range studentToGroups {
		groupIDs = append(groupIDs, studentToGroup.GroupID)
	}
	clinicalRotationEvents, err := server.store.ListRotationEventsByAcademicYearIDAndGroupID(ctx,
		db.ListRotationEventsByAcademicYearIDAndGroupIDParams{
			AcademicYearID: academicYear.ID,
			GroupIds:       groupIDs,
		})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	var (
		service       db.Service
		specialty     db.Specialty
		hospital      db.Hospital
		group         db.Group
		eventResponse []clinicalRotationEventResponse
	)
	for _, event := range clinicalRotationEvents {
		service, err = server.store.GetServiceByID(ctx, event.ServiceID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		specialty, err = server.store.GetSpecialtyByID(ctx, service.SpecialtyID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		hospital, err = server.store.GetHospitalByID(ctx, service.HospitalID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		group, err = server.store.GetGroupByID(ctx, event.GroupID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		eventResponse = append(eventResponse, clinicalRotationEventResponse{
			EventId:       event.ID,
			GroupName:     group.Name,
			SpecialtyName: specialty.Name,
			HospitalName:  hospital.Name,
			ServiceName:   service.Name,
			StartDate:     event.StartDate,
			EndDate:       event.EndDate,
		})
	}
	ctx.JSON(http.StatusOK, eventResponse)
}

// attendingCalendar
// @Summary list calendar events for an attending in an academic year
// @Description list calendar events for an attending in an academic year
// @Tags ClinicalRotation
// @Accept	json
// @Produce  json
// @Param academicYearName query string true "academic year name"
// @Param attendingID query string true "attending ID"
// @Success 200 {object} []clinicalRotationEventResponse "ok"
// @Router /rotation/attending [get]
func (server *Server) attendingCalendar(ctx *gin.Context) {
	academicYearName := ctx.Query("academicYearName")
	academicYear, err := server.store.GetAcademicYearByName(ctx, academicYearName)
	if err != nil {
		if err == pgx.ErrNoRows {
			log.Info().Msgf("cannot find academic year %v", academicYearName)
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	attendingID := ctx.Query("attendingID")
	attending, err := server.store.GetAttendingByAttendingId(ctx, attendingID)
	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	attendingToServices, err := server.store.GetServiceToAttendingByAttendingID(ctx, attending.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	serviceIDs := make([]int64, 0)
	for _, attendingToService := range attendingToServices {
		serviceIDs = append(serviceIDs, attendingToService.ServiceID)
	}
	clinicalRotationEvents, err := server.store.ListRotationEventsByAcademicYearIDAndServiceID(ctx,
		db.ListRotationEventsByAcademicYearIDAndServiceIDParams{
			AcademicYearID: academicYear.ID,
			ServiceIds:     serviceIDs,
		})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	var (
		service       db.Service
		specialty     db.Specialty
		hospital      db.Hospital
		group         db.Group
		eventResponse []clinicalRotationEventResponse
	)
	for _, event := range clinicalRotationEvents {
		service, err = server.store.GetServiceByID(ctx, event.ServiceID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		specialty, err = server.store.GetSpecialtyByID(ctx, service.SpecialtyID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		hospital, err = server.store.GetHospitalByID(ctx, service.HospitalID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		group, err = server.store.GetGroupByID(ctx, event.GroupID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		eventResponse = append(eventResponse, clinicalRotationEventResponse{
			EventId:       event.ID,
			GroupName:     group.Name,
			SpecialtyName: specialty.Name,
			HospitalName:  hospital.Name,
			ServiceName:   service.Name,
			StartDate:     event.StartDate,
			EndDate:       event.EndDate,
		})
	}
	ctx.JSON(http.StatusOK, eventResponse)
}

type clinicalRotationEventDetailResponse struct {
	GroupName     string              `json:"group_name"`
	SpecialtyName string              `json:"specialty_name"`
	HospitalName  string              `json:"hospital_name"`
	ServiceName   string              `json:"service_name"`
	StartDate     time.Time           `json:"start_date"`
	EndDate       time.Time           `json:"end_date"`
	Students      []studentResponse   `json:"students"`
	Attendings    []attendingResponse `json:"attendings"`
}

// clinicalRotationEventDetail
// @Summary provide detail of a clinical rotation event
// @Description provide students and attending info and the clinical info
// @Tags ClinicalRotation
// @Accept	json
// @Produce  json
// @Param eventID query string true "clinical rotation event ID"
// @Success 200 {object} clinicalRotationEventDetailResponse "ok"
// @Router /rotation/detail [get]
func (server *Server) clinicalRotationEventDetail(ctx *gin.Context) {
	eventID := ctx.Query("eventID")
	eventIDInt, err := strconv.Atoi(eventID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	event, err := server.store.GetRotationEventByID(ctx, int64(eventIDInt))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	service, err := server.store.GetServiceByID(ctx, event.ServiceID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	specialty, err := server.store.GetSpecialtyByID(ctx, service.SpecialtyID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	hospital, err := server.store.GetHospitalByID(ctx, service.HospitalID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	group, err := server.store.GetGroupByID(ctx, event.GroupID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	studentToGroups, err := server.store.GetStudentToGroupByGroupID(ctx, group.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	students := make([]studentResponse, 0)
	for _, studentToGroup := range studentToGroups {
		student, err := server.store.GetStudentByID(ctx, studentToGroup.StudentID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		userInfo, err := server.store.GetUserByID(ctx, student.UserID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		students = append(students, studentResponse{
			Email:     userInfo.Email,
			StudentID: student.StudentID,
			FirstName: student.FirstName,
			LastName:  student.LastName,
			Mobile:    student.Mobile,
			CreatedAt: student.CreatedAt,
		})
	}

	serviceToAttendings, err := server.store.GetServiceToAttendingByServiceID(ctx, service.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	attendings := make([]attendingResponse, 0)
	for _, serviceToAttending := range serviceToAttendings {
		attending, err := server.store.GetAttendingByID(ctx, serviceToAttending.AttendingID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		attendings = append(attendings, attendingResponse{
			UserID:      attending.UserID,
			AttendingID: attending.AttendingID,
			FirstName:   attending.FirstName,
			LastName:    attending.LastName,
			Mobile:      attending.Mobile,
			CreatedAt:   attending.CreatedAt,
		})
	}
	ctx.JSON(http.StatusOK, clinicalRotationEventDetailResponse{
		GroupName:     group.Name,
		SpecialtyName: specialty.Name,
		HospitalName:  hospital.Name,
		ServiceName:   service.Name,
		StartDate:     event.StartDate,
		EndDate:       event.EndDate,
		Students:      students,
		Attendings:    attendings,
	})
}
