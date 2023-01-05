package api

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v4"

	db "github.com/VinCPR/backend/db/sqlc"
)

type createServiceRequest struct {
	Hospital    string `json:"hospital" binding:"required"`
	Specialty   string `json:"specialty" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type serviceResponse struct {
	Hospital    string    `json:"hospital"`
	Specialty   string    `json:"specialty"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

// createService
// @Summary create new service
// @Description create new service
// @Tags Service
// @Accept	json
// @Produce  json
// @Param body body createServiceRequest true "input required: service hospitalID, specialtyID, name, description"
// @Success 200 {object} serviceResponse "ok"
// @Router /service [post]
func (server *Server) createService(ctx *gin.Context) {
	var req createServiceRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	hospital, err := server.store.GetHospitalByName(ctx, req.Hospital)
	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	specialty, err := server.store.GetSpecialtyByName(ctx, req.Specialty)
	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	service, err := server.store.CreateService(ctx, db.CreateServiceParams{
		HospitalID:  hospital.ID,
		SpecialtyID: specialty.ID,
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
	ctx.JSON(http.StatusOK, serviceResponse{
		Hospital:    req.Hospital,
		Specialty:   req.Specialty,
		Name:        service.Name,
		Description: service.Description,
		CreatedAt:   service.CreatedAt,
	})
}

// listServicesByHospitalID
// @Summary list created service
// @Description list created service
// @Tags Services
// @Accept	json
// @Produce  json
// @Param pageNumber query string true "page number"
// @Param pageSize query string true "page size"
// @Success 200 {object} []serviceResponse "ok"
// @Router /service/list/hospital [get]
func (server *Server) listServicesbyHospital(ctx *gin.Context) {
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

	services, err := server.store.ListServicesByHospitalID(ctx, db.ListServicesByHospitalIDParams{
		Limit:  p.Limit(),
		Offset: p.Offset(),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := make([]serviceResponse, 0)
	for _, service := range services {
		var hospital db.Hospital
		var specialty db.Specialty

		hospital, err = server.store.GetHospitalByID(ctx, service.ID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		specialty, err = server.store.GetSpecialtyByID(ctx, service.SpecialtyID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		response = append(response, serviceResponse{
			Hospital:    hospital.Name,
			Specialty:   specialty.Name,
			Name:        service.Name,
			Description: service.Description,
			CreatedAt:   service.CreatedAt,
		})
	}
	ctx.JSON(http.StatusOK, response)
}

// listServicesBySpecialtyID
// @Summary list created service
// @Description list created service
// @Tags Services
// @Accept	json
// @Produce  json
// @Param pageNumber query string true "page number"
// @Param pageSize query string true "page size"
// @Success 200 {object} []serviceResponse "ok"
// @Router /service/list/specialty [get]
func (server *Server) listServicesbySpecialty(ctx *gin.Context) {
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

	services, err := server.store.ListServicesBySpecialtyID(ctx, db.ListServicesBySpecialtyIDParams{
		Limit:  p.Limit(),
		Offset: p.Offset(),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := make([]serviceResponse, 0)
	for _, service := range services {
		var hospital db.Hospital
		var specialty db.Specialty

		hospital, err = server.store.GetHospitalByID(ctx, service.ID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		specialty, err = server.store.GetSpecialtyByID(ctx, service.SpecialtyID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		response = append(response, serviceResponse{
			Hospital:    hospital.Name,
			Specialty:   specialty.Name,
			Name:        service.Name,
			Description: service.Description,
			CreatedAt:   service.CreatedAt,
		})
	}
	ctx.JSON(http.StatusOK, response)
}

// listServicesBySpecialtyIDAndHospitalID
// @Summary list created service
// @Description list created service
// @Tags Services
// @Accept	json
// @Produce  json
// @Param pageNumber query string true "page number"
// @Param pageSize query string true "page size"
// @Success 200 {object} []serviceResponse "ok"
// @Router /service/list/specialty_and_hospital [get]
func (server *Server) listServicesBySpecialtyAndHospital(ctx *gin.Context) {
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

	services, err := server.store.ListServicesBySpecialtyIDAndHospitalID(ctx, db.ListServicesBySpecialtyIDAndHospitalIDParams{
		Limit:  p.Limit(),
		Offset: p.Offset(),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := make([]serviceResponse, 0)
	for _, service := range services {
		var hospital db.Hospital
		var specialty db.Specialty

		hospital, err = server.store.GetHospitalByID(ctx, service.ID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		specialty, err = server.store.GetSpecialtyByID(ctx, service.SpecialtyID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		response = append(response, serviceResponse{
			Hospital:    hospital.Name,
			Specialty:   specialty.Name,
			Name:        service.Name,
			Description: service.Description,
			CreatedAt:   service.CreatedAt,
		})
	}
	ctx.JSON(http.StatusOK, response)
}
