package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	db "github.com/VinCPR/backend/db/sqlc"
)

type HospitalResponse struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Address     string    `json:"address"`
	CreatedAt   time.Time `json:"created_at"`
}

// listHospitals
// @Summary list created hospital
// @Description list created hospital
// @Tags Hospital
// @Accept	json
// @Produce  json
// @Param pageNumber query string true "page number"
// @Param pageSize query string true "page size"
// @Success 200 {object} []HospitalResponse "ok"
// @Router /hospital/list [get]

func (server *Server) listHospitalsByName(ctx *gin.Context) {
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

	hospitals, err := server.store.ListHospitalsByName(ctx, db.ListHospitalsByNameParams{
		Limit:  p.Limit(),
		Offset: p.Offset(),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	HospitalsResponse := make([]HospitalResponse, 0)
	for _, hospital := range hospitals {
		HospitalsResponse = append(HospitalsResponse, HospitalResponse{
			Name:        hospital.Name,
			Description: hospital.Description,
			Address:     hospital.Address,
			CreatedAt:   hospital.CreatedAt,
		})
	}
	ctx.JSON(http.StatusOK, HospitalsResponse)
}
