package api

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"

	db "github.com/VinCPR/backend/db/sqlc"
)

type createSpecialtyRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type specialtyResponse struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

// createSpecialty
// @Summary create new specialty
// @Description create new specialty
// @Tags Specialty
// @Accept	json
// @Produce  json
// @Param body body createSpecialtyRequest true "input required: specialty name, description, address"
// @Success 200 {object} specialtyResponse "ok"
// @Router /specialty [post]
func (server *Server) createSpecialty(ctx *gin.Context) {
	var req createSpecialtyRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	specialty, err := server.store.CreateSpecialty(ctx, db.CreateSpecialtyParams{
		Name:        req.Name,
		Description: req.Description,
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
	ctx.JSON(http.StatusOK, specialtyResponse{
		Name:        specialty.Name,
		Description: specialty.Description,
		CreatedAt:   specialty.CreatedAt,
	})
}

// listSpecialties
// @Summary list created specialty
// @Description list created specialty
// @Tags Specialties
// @Accept	json
// @Produce  json
// @Param pageNumber query string true "page number"
// @Param pageSize query string true "page size"
// @Success 200 {object} []specialtyResponse "ok"
// @Router /specialty/list [get]
func (server *Server) listSpecialtiesByName(ctx *gin.Context) {
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

	specialties, err := server.store.ListSpecialtiesByName(ctx, db.ListSpecialtiesByNameParams{
		Limit:  p.Limit(),
		Offset: p.Offset(),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	SpecialtiesResponse := make([]specialtyResponse, 0)
	for _, specialty := range specialties {
		SpecialtiesResponse = append(SpecialtiesResponse, specialtyResponse{
			Name:        specialty.Name,
			Description: specialty.Description,
			CreatedAt:   specialty.CreatedAt,
		})
	}
	ctx.JSON(http.StatusOK, SpecialtiesResponse)
}
