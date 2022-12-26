package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

type periodResponse struct {
	AcademicYearID int64     `json:"academic_year_id"`
	Name           string    `json:"name"`
	StartDate      time.Time `json:"start_date"`
	EndDate        time.Time `json:"end_date"`
	CreatedAt      time.Time `json:"created_at"`
}

// listPeriodsByAcademicYear
// @Summary list periods in that academic year
// @Description list periods in that academic year
// @Tags Periods
// @Accept	json
// @Produce  json
// @Param academicYearName query string true "academic year name"
// @Success 200 {object} []periodResponse "ok"
// @Router /period/list [get]
func (server *Server) listPeriodsByAcademicYear(ctx *gin.Context) {
	academicYearName := ctx.Query("academicYearName")

	academicYear, err := server.store.GetAcademicYearByName(ctx, academicYearName)
	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	periods, err := server.store.ListPeriodsByStartDate(ctx, academicYear.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	periodsResponse := make([]periodResponse, 0)
	for _, period := range periods {
		periodsResponse = append(periodsResponse, periodResponse{
			AcademicYearID: period.AcademicYearID,
			Name:           period.Name,
			StartDate:      period.StartDate,
			EndDate:        period.EndDate,
			CreatedAt:      period.CreatedAt,
		})
	}
	ctx.JSON(http.StatusOK, periodsResponse)
}
