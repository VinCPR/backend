package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"

	db "github.com/VinCPR/backend/db/sqlc"
)

type createAcademicYearRequest struct {
	Name      string    `json:"name" binding:"required"`
	StartDate time.Time `json:"start_date" form:"start_date" binding:"required,ltefield=EndDate" time_format:"2006-01-02"`
	EndDate   time.Time `json:"end_date" form:"end_date" binding:"required" time_format:"2006-01-02"`
}

type listAcademicYearsRequest struct {
	Limit  int32 `json:"limit" binding:"required"`
	Offset int32 `json:"offset" binding:"required"`
}

type academicYearResponse struct {
	Name      string    `json:"name"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	CreatedAt time.Time `json:"created_at"`
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
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
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
