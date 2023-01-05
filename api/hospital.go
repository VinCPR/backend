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

type createHospitalRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Address     string `json:"address" binding:"required"`
}

type hospitalResponse struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Address     string    `json:"address"`
	CreatedAt   time.Time `json:"created_at"`
}

// createHospital
// @Summary create new hospital
// @Description create new hospital
// @Tags Hospital
// @Accept	json
// @Produce  json
// @Param body body createHospitalRequest true "input required: hospital name, description, address"
// @Success 200 {object} HospitalResponse "ok"
// @Router /hospital [post]
func (server *Server) createHospital(ctx *gin.Context) {
	var req createHospitalRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	hospital, err := server.store.CreateHospital(ctx, db.CreateHospitalParams{
		Name:        req.Name,
		Description: req.Description,
		Address:     req.Address,
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
	ctx.JSON(http.StatusOK, hospitalResponse{
		Name:        hospital.Name,
		Description: hospital.Description,
		Address:     hospital.Address,
		CreatedAt:   hospital.CreatedAt,
	})
}

// listHospitals
// @Summary list created hospital
// @Description list created hospital
// @Tags Hospitals
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

	HospitalsResponse := make([]hospitalResponse, 0)
	for _, hospital := range hospitals {
		HospitalsResponse = append(HospitalsResponse, hospitalResponse{
			Name:        hospital.Name,
			Description: hospital.Description,
			Address:     hospital.Address,
			CreatedAt:   hospital.CreatedAt,
		})
	}
	ctx.JSON(http.StatusOK, HospitalsResponse)
}
