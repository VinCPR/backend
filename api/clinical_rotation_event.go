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
// @Param studentUserID query string true "student user ID"
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

	studentUserID := ctx.Query("studentUserID")
	studentUserIDInt, err := strconv.Atoi(studentUserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	student, err := server.store.GetStudentByUserId(ctx, int64(studentUserIDInt))
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
// @Param attendingUserID query string true "attending user ID"
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

	attendingUserID := ctx.Query("attendingUserID")
	attendingUserIDInt, err := strconv.Atoi(attendingUserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	attending, err := server.store.GetAttendingByUserId(ctx, int64(attendingUserIDInt))
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

func (server *Server) eventDetail(ctx *gin.Context) {

}