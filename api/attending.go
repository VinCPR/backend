package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	db "github.com/VinCPR/backend/db/sqlc"
)

type AttendingResponse struct {
	UserID    int64     `json:"user_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Mobile    string    `json:"mobile"`
	CreatedAt time.Time `json:"created_at"`
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

	AttendingsResponse := make([]AttendingResponse, 0)
	for _, attending := range attendings {
		AttendingsResponse = append(AttendingsResponse, AttendingResponse{
			UserID:    attending.UserID,
			FirstName: attending.FirstName,
			LastName:  attending.LastName,
			Mobile:    attending.Mobile,
			CreatedAt: attending.CreatedAt,
		})
	}
	ctx.JSON(http.StatusOK, AttendingsResponse)
}
