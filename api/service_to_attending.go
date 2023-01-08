package api

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"

	db "github.com/VinCPR/backend/db/sqlc"
)

type createServiceToAttendingRequest struct {
	ServiceID   int64 `json:"service_id" binding:"required"`
	AttendingID int64 `json:"attending_id" binding:"required"`
}

type ServiceToAttendingResponse struct {
	ServiceID   int64     `json:"service_id"`
	AttendingID int64     `json:"attending_id"`
	CreatedAt   time.Time `json:"created_at"`
}

// createServiceToAttending
// @Summary create new service to attending
// @Description create new service to attending
// @Tags ServiceToAttending
// @Accept	json
// @Produce  json
// @Param body body createServiceToAttendingRequest true "input required: serviceID, attendingID"
// @Success 200 {object} ServiceToAttendingResponse "ok"
// @Router /service_to_attending [post]
func (server *Server) createServiceToAttending(ctx *gin.Context) {
	var req createServiceToAttendingRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	service_to_attending, err := server.store.CreateServiceToAttending(ctx, db.CreateServiceToAttendingParams{
		ServiceID:   req.ServiceID,
		AttendingID: req.AttendingID,
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
	ctx.JSON(http.StatusOK, ServiceToAttendingResponse{
		ServiceID:   service_to_attending.ServiceID,
		AttendingID: service_to_attending.AttendingID,
		CreatedAt:   service_to_attending.CreatedAt,
	})
}

// listServicesToAttendingsbyServiceID
// @Summary list created services to attendings
// @Description list created services to attendings
// @Tags ServicesToAttendings
// @Accept	json
// @Produce  json
// @Param pageNumber query string true "page number"
// @Param pageSize query string true "page size"
// @Success 200 {object} []ServiceToAttendingResponse "ok"
// @Router /service_to_attending/list/service_id [get]

func (server *Server) listServicesToAttendingsbyServiceID(ctx *gin.Context) {

	services_to_attendings, err := server.store.ListServicesToAttendingsByServiceID(context.Background())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ServicesToAttendingsResponse := make([]ServiceToAttendingResponse, 0)
	for _, service_to_attending := range services_to_attendings {
		ServicesToAttendingsResponse = append(ServicesToAttendingsResponse, ServiceToAttendingResponse{
			ServiceID:   service_to_attending.ServiceID,
			AttendingID: service_to_attending.AttendingID,
			CreatedAt:   service_to_attending.CreatedAt,
		})
	}
	ctx.JSON(http.StatusOK, ServicesToAttendingsResponse)
}

// listServicesToAttendingsbyAttendingID
// @Summary list created services to attendings
// @Description list created services to attendings
// @Tags ServicesToAttendings
// @Accept	json
// @Produce  json
// @Param pageNumber query string true "page number"
// @Param pageSize query string true "page size"
// @Success 200 {object} []ServiceToAttendingResponse "ok"
// @Router /service_to_attending/list/attending_id [get]

func (server *Server) listServicesToAttendingsbyAttendingID(ctx *gin.Context) {

	services_to_attendings, err := server.store.ListServicesToAttendingsByAttendingID(context.Background())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ServicesToAttendingsResponse := make([]ServiceToAttendingResponse, 0)
	for _, service_to_attending := range services_to_attendings {
		ServicesToAttendingsResponse = append(ServicesToAttendingsResponse, ServiceToAttendingResponse{
			ServiceID:   service_to_attending.ServiceID,
			AttendingID: service_to_attending.AttendingID,
			CreatedAt:   service_to_attending.CreatedAt,
		})
	}
	ctx.JSON(http.StatusOK, ServicesToAttendingsResponse)
}

// listServicesToAttendingsbyAll
// @Summary list created services to attendings
// @Description list created services to attendings
// @Tags ServicesToAttendings
// @Accept	json
// @Produce  json
// @Param pageNumber query string true "page number"
// @Param pageSize query string true "page size"
// @Success 200 {object} []ServiceToAttendingResponse "ok"
// @Router /service_to_attending/list/all [get]

func (server *Server) listServicesToAttendingsbyAll(ctx *gin.Context) {

	services_to_attendings, err := server.store.ListServicesToAttendingsByServiceIDAndAttendingID(context.Background())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ServicesToAttendingsResponse := make([]ServiceToAttendingResponse, 0)
	for _, service_to_attending := range services_to_attendings {
		ServicesToAttendingsResponse = append(ServicesToAttendingsResponse, ServiceToAttendingResponse{
			ServiceID:   service_to_attending.ServiceID,
			AttendingID: service_to_attending.AttendingID,
			CreatedAt:   service_to_attending.CreatedAt,
		})
	}
	ctx.JSON(http.StatusOK, ServicesToAttendingsResponse)
}
