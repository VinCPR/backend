package api

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v4"

	db "github.com/VinCPR/backend/db/sqlc"
)

type createAcademicYearRequest struct {
	Name      string    `json:"name" binding:"required"`
	StartDate time.Time `json:"start_date" form:"start_date" binding:"required,ltefield=EndDate" time_format:"2006-01-02"`
	EndDate   time.Time `json:"end_date" form:"end_date" binding:"required" time_format:"2006-01-02"`
}

type academicYearResponse struct {
	Name      string    `json:"name"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	CreatedAt time.Time `json:"created_at"`
}

type academicCalendarEventResponse struct {
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	CreatedAt time.Time `json:"created_at"`
}

type getAcademicCalendarResponse struct {
	AcademicYear           academicYearResponse            `json:"academic_year"`
	AcademicCalendarEvents []academicCalendarEventResponse `json:"academic_calendar_events"`
}

// createAcademicYear
// @Summary create new academic year
// @Description create new academic year
// @Tags AcademicCalendar
// @Accept	json
// @Produce  json
// @Param body body createAcademicYearRequest true "input required: academic year name, start date, end date"
// @Success 200 {object} academicYearResponse "ok"
// @Router /academic_year [post]
func (server *Server) createAcademicYear(ctx *gin.Context) {
	var req createAcademicYearRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	academicYear, err := server.store.CreateAcademicYear(ctx, db.CreateAcademicYearParams{
		Name:      req.Name,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
	})
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			ctx.JSON(http.StatusForbidden, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, academicYearResponse{
		Name:      academicYear.Name,
		StartDate: academicYear.StartDate,
		EndDate:   academicYear.EndDate,
		CreatedAt: academicYear.CreatedAt,
	})
}

// listAcademicYears
// @Summary list created academic year
// @Description list created academic year
// @Tags AcademicCalendar
// @Accept	json
// @Produce  json
// @Param pageNumber query string true "page number"
// @Param pageSize query string true "page size"
// @Success 200 {object} []academicYearResponse "ok"
// @Router /academic_year/list [get]
func (server *Server) listAcademicYears(ctx *gin.Context) {
	pageNumber := ctx.Query("pageNumber")
	pageSize := ctx.Query("pageSize")

	// init pagination
	pageNumberInt, err := strconv.Atoi(pageNumber)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	p := InitPagination(int32(pageNumberInt), int32(pageSizeInt))

	academicYears, err := server.store.ListAcademicYearByEndDate(ctx, db.ListAcademicYearByEndDateParams{
		Limit:  p.Limit(),
		Offset: p.Offset(),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	academicYearsResponse := make([]academicYearResponse, 0)
	for _, academicYear := range academicYears {
		academicYearsResponse = append(academicYearsResponse, academicYearResponse{
			Name:      academicYear.Name,
			StartDate: academicYear.StartDate,
			EndDate:   academicYear.EndDate,
			CreatedAt: academicYear.CreatedAt,
		})
	}
	ctx.JSON(http.StatusOK, academicYearsResponse)
}

// getAcademicCalendar
// @Summary return the list of events of an academic year
// @Description return the list of events of an academic year
// @Tags AcademicCalendar
// @Accept	json
// @Produce  json
// @Param academicYearName query string true "academic year name"
// @Success 200 {object} []getAcademicCalendarResponse "ok"
// @Router /academic_year/calendar [get]
func (server *Server) getAcademicCalendar(ctx *gin.Context) {
	academicYearName := ctx.Query("academicYearName")

	academicYear, err := server.store.GetAcademicYearByName(ctx, academicYearName)
	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	academicCalendarEvents, err := server.store.ListEventsByAcademicYearID(ctx, academicYear.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	eventResponse := make([]academicCalendarEventResponse, 0)
	for _, event := range academicCalendarEvents {
		eventResponse = append(eventResponse, academicCalendarEventResponse{
			Name:      event.Name,
			Type:      event.Type,
			StartDate: event.StartDate,
			EndDate:   event.EndDate,
			CreatedAt: event.CreatedAt,
		})
	}
	ctx.JSON(http.StatusOK, getAcademicCalendarResponse{
		AcademicYear: academicYearResponse{
			Name:      academicYear.Name,
			StartDate: academicYear.StartDate,
			EndDate:   academicYear.EndDate,
			CreatedAt: academicYear.CreatedAt,
		},
		AcademicCalendarEvents: eventResponse,
	})
}

// TODO insert academic calendar events API (only for admin)
