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

type createAttendingRequest struct {
	UserID    int64  `json:"user_id" binding:"required"`
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Mobile    string `json:"mobile" binding:"required"`
}

type attendingResponse struct {
	UserID    int64     `json:"user_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Mobile    string    `json:"mobile"`
	CreatedAt time.Time `json:"created_at"`
}

// createAttending
// @Summary create new Attending
// @Description create new Attending
// @Tags Attending
// @Accept	json
// @Produce  json
// @Param body body createAttendingRequest true "input required: attending user_id, firstname, lastname, mobile"
// @Success 200 {object} AttendingResponse "ok"
// @Router /attending [post]
func (server *Server) createAttending(ctx *gin.Context) {
	var req createAttendingRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	attending, err := server.store.CreateAttending(ctx, db.CreateAttendingParams{
		UserID:    req.UserID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Mobile:    req.Mobile,
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
	ctx.JSON(http.StatusOK, attendingResponse{
		UserID:    attending.UserID,
		FirstName: attending.FirstName,
		LastName:  attending.LastName,
		Mobile:    attending.Mobile,
		CreatedAt: attending.CreatedAt,
	})
}

// listAttendingsByName
// @Summary list created attending
// @Description list created attending
// @Tags Attendings
// @Accept	json
// @Produce  json
// @Param pageNumber query string true "page number"
// @Param pageSize query string true "page size"
// @Success 200 {object} []AttendingResponse "ok"
// @Router /attending/list[get]
func (server *Server) listAttendingsByName(ctx *gin.Context) {
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

	attendings, err := server.store.ListAttendingsByName(ctx, db.ListAttendingsByNameParams{
		Limit:  p.Limit(),
		Offset: p.Offset(),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	AttendingsResponse := make([]attendingResponse, 0)
	for _, attending := range attendings {
		AttendingsResponse = append(AttendingsResponse, attendingResponse{
			UserID:    attending.UserID,
			FirstName: attending.FirstName,
			LastName:  attending.LastName,
			Mobile:    attending.Mobile,
			CreatedAt: attending.CreatedAt,
		})
	}
	ctx.JSON(http.StatusOK, AttendingsResponse)
}
