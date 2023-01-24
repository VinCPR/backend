package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

type groupResponse struct {
	AcademicYearID int64     `json:"academic_year_id"`
	Name           string    `json:"name"`
	CreatedAt      time.Time `json:"created_at"`
}

// listGroupsByAcademicYear
// @Summary list groups in that academic year
// @Description list groups in that academic year
// @Tags Groups
// @Accept	json
// @Produce  json
// @Param academicYearName query string true "academic year name"
// @Success 200 {object} []groupResponse "ok"
// @Router /group/list [get]
func (server *Server) listGroupsByAcademicYear(ctx *gin.Context) {
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

	groups, err := server.store.ListGroupsByName(ctx, academicYear.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	groupsResponse := make([]groupResponse, 0)
	for _, group := range groups {
		groupsResponse = append(groupsResponse, groupResponse{
			AcademicYearID: group.AcademicYearID,
			Name:           group.Name,
			CreatedAt:      group.CreatedAt,
		})
	}
	ctx.JSON(http.StatusOK, groupsResponse)
}
