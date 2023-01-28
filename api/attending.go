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
	"github.com/VinCPR/backend/util"
)

type createAttendingRequest struct {
	AttendingID string `json:"attending_id" binding:"required"`
	FirstName   string `json:"firstname" binding:"required"`
	LastName    string `json:"lastname" binding:"required"`
	Mobile      string `json:"mobile" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,min=8,max=64"`
}

type attendingResponse struct {
	Email       string    `json:"email"`
	AttendingID string    `json:"attending_id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Mobile      string    `json:"mobile"`
	CreatedAt   time.Time `json:"created_at"`
}

// createAttending
// @Summary create new Attending
// @Description create new Attending
// @Tags Attendings
// @Accept	json
// @Produce  json
// @Param body body createAttendingRequest true "input required: attending user_id, firstname, lastname, mobile"
// @Success 200 {object} attendingResponse "ok"
// @Router /attending [post]
func (server *Server) createAttending(ctx *gin.Context) {
	var req createAttendingRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	user, err := server.store.CreateUser(ctx, db.CreateUserParams{
		Email:          req.Email,
		HashedPassword: hashedPassword,
		RoleName:       "attending",
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

	attending, err := server.store.CreateAttending(ctx, db.CreateAttendingParams{
		UserID:      user.ID,
		AttendingID: req.AttendingID,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Mobile:      req.Mobile,
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
		Email:       req.Email,
		AttendingID: attending.AttendingID,
		FirstName:   attending.FirstName,
		LastName:    attending.LastName,
		Mobile:      attending.Mobile,
		CreatedAt:   attending.CreatedAt,
	})
}

// getAttendingInfoByEmail
// @Summary provide detail of an attending given email address
// @Description provide detail of an attending given email address
// @Tags Attendings
// @Accept	json
// @Produce  json
// @Param email query string true "attending email address"
// @Success 200 {object} attendingResponse "ok"
// @Router /attending/info [get]
func (server *Server) getAttendingInfoByEmail(ctx *gin.Context) {
	email := ctx.Query("email")
	user, err := server.store.GetUserByEmail(ctx, email)
	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	student, err := server.store.GetAttendingByUserId(ctx, user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, attendingResponse{
		Email:       email,
		AttendingID: student.AttendingID,
		FirstName:   student.FirstName,
		LastName:    student.LastName,
		Mobile:      student.Mobile,
		CreatedAt:   student.CreatedAt,
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
// @Success 200 {object} []attendingResponse "ok"
// @Router /attending/list/name [get]
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

	attendingsResponse := make([]attendingResponse, 0)
	for _, attending := range attendings {
		userInfo, err := server.store.GetUserByID(ctx, attending.UserID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		attendingsResponse = append(attendingsResponse, attendingResponse{
			Email:       userInfo.Email,
			AttendingID: attending.AttendingID,
			FirstName:   attending.FirstName,
			LastName:    attending.LastName,
			Mobile:      attending.Mobile,
			CreatedAt:   attending.CreatedAt,
		})
	}
	ctx.JSON(http.StatusOK, attendingsResponse)
}

// listAttendingsByAttendingID
// @Summary list created attending order by attending id
// @Description list created attending order by attending id
// @Tags Attendings
// @Accept	json
// @Produce  json
// @Param pageNumber query string true "page number"
// @Param pageSize query string true "page size"
// @Success 200 {object} []attendingResponse "ok"
// @Router /attending/list/attending_id [get]
func (server *Server) listAttendingsByAttendingID(ctx *gin.Context) {
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

	attendings, err := server.store.ListAttendingsByAttendingID(ctx, db.ListAttendingsByAttendingIDParams{
		Limit:  p.Limit(),
		Offset: p.Offset(),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	attendingsResponse := make([]attendingResponse, 0)
	for _, attending := range attendings {
		userInfo, err := server.store.GetUserByID(ctx, attending.UserID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		attendingsResponse = append(attendingsResponse, attendingResponse{
			Email:       userInfo.Email,
			AttendingID: attending.AttendingID,
			FirstName:   attending.FirstName,
			LastName:    attending.LastName,
			Mobile:      attending.Mobile,
			CreatedAt:   attending.CreatedAt,
		})
	}
	ctx.JSON(http.StatusOK, attendingsResponse)
}
