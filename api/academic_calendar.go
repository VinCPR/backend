package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"

	db "github.com/VinCPR/backend/db/sqlc"
)

type createAcademicYearRequest struct {
	Name      string    `json:"name" binding:"required"`
	StartDate time.Time `form:"start_date" binding:"required,ltefield=EndDate" time_format:"2006-01-02"`
	EndDate   time.Time `form:"end_date" binding:"required" time_format:"2006-01-02"`
}

type createAcademicYearResponse struct {
	Name      string
	StartDate time.Time
	EndDate   time.Time
	CreatedAt time.Time
}

// createAcademicYear
// @Summary create new academic year
// @Description create new academic year
// @Tags AcademicCalendar
// @Accept	json
// @Produce  json
// @Param body body createAcademicYearRequest true "input requires academic year name, start date, end date"
// @Success 200 {object} createAcademicYearResponse "ok"
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
	ctx.JSON(http.StatusOK, createAcademicYearResponse{
		Name:      academicYear.Name,
		StartDate: academicYear.StartDate,
		EndDate:   academicYear.EndDate,
		CreatedAt: academicYear.CreatedAt,
	})
}
