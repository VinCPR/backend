package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

type blockResponse struct {
	AcademicYearID int64     `json:"academic_year_id"`
	Name           string    `json:"name"`
	Period         int64     `json:"period"`
	CreatedAt      time.Time `json:"created_at"`
}

// listBlocksByAcademicYear
// @Summary list blocks in that academic year
// @Description list blocks in that academic year
// @Tags Blocks
// @Accept	json
// @Produce  json
// @Param academicYearName query string true "academic year name"
// @Success 200 {object} []blockResponse "ok"
// @Router /block/list [get]
func (server *Server) listBlocksByAcademicYear(ctx *gin.Context) {
	academicYearName := ctx.Query("academicYearName")

	academicYear, err := server.store.GetAcademicYearByName(ctx, academicYearName)
	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	blocks, err := server.store.ListBlocksByAcademicYear(ctx, academicYear.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	blocksResponse := make([]blockResponse, 0)
	for _, block := range blocks {
		blocksResponse = append(blocksResponse, blockResponse{
			AcademicYearID: block.AcademicYearID,
			Name:           block.Name,
			Period:         block.Period,
			CreatedAt:      block.CreatedAt,
		})
	}
	ctx.JSON(http.StatusOK, blocksResponse)
}
