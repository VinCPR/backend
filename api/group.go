package api

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v4"

	db "github.com/VinCPR/backend/db/sqlc"
)

type createGroupRequest struct {
	AcademicYearName string `json:"academic_year_name" binding:"required"`
	Name             string `json:"name" binding:"required"`
}

type groupResponse struct {
	AcademicYearName string    `json:"academic_year_name"`
	Name             string    `json:"name"`
	CreatedAt        time.Time `json:"created_at"`
}

// createGroup
// @Summary create new group
// @Description create new group
// @Tags AcademicCalendar
// @Accept	json
// @Produce  json
// @Param body body createGroupRequest true "input required: academic year name, group name"
// @Success 200 {object} groupResponse "ok"
// @Router /group [post]
func (server *Server) createGroup(ctx *gin.Context) {
	var req createGroupRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	academicYear, err := server.store.GetAcademicYearByName(ctx, req.AcademicYearName)
	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	group, err := server.store.CreateGroup(ctx, db.CreateGroupParams{
		AcademicYearID: academicYear.ID,
		Name:           req.Name,
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
	ctx.JSON(http.StatusOK, groupResponse{
		AcademicYearName: req.AcademicYearName,
		Name:             group.Name,
		CreatedAt:        group.CreatedAt,
	})
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
			AcademicYearName: academicYearName,
			Name:             group.Name,
			CreatedAt:        group.CreatedAt,
		})
	}
	ctx.JSON(http.StatusOK, groupsResponse)
}
