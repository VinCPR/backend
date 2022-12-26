package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	db "github.com/VinCPR/backend/db/sqlc"
)

type SpecialtyResponse struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

// listSpecialties
// @Summary list created specialty
// @Description list created specialty
// @Tags Specialty
// @Accept	json
// @Produce  json
// @Param pageNumber query string true "page number"
// @Param pageSize query string true "page size"
// @Success 200 {object} []SpecialtyResponse "ok"
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

	SpecialtiesResponse := make([]SpecialtyResponse, 0)
	for _, specialty := range specialties {
		SpecialtiesResponse = append(SpecialtiesResponse, SpecialtyResponse{
			Name:        specialty.Name,
			Description: specialty.Description,
			CreatedAt:   specialty.CreatedAt,
		})
	}
	ctx.JSON(http.StatusOK, SpecialtiesResponse)
}
