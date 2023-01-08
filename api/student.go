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

type createStudentRequest struct {
	UserID    int64  `json:"user_id" binding:"required"`
	StudentID string `json:"student_id" binding:"required"`
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Mobile    string `json:"mobile" binding:"required"`
}

type studentResponse struct {
	UserID    int64     `json:"user_id"`
	StudentID string    `json:"student_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Mobile    string    `json:"mobile"`
	CreatedAt time.Time `json:"created_at"`
}

// createStudent
// @Summary create new student
// @Description create new student
// @Tags Students
// @Accept	json
// @Produce  json
// @Param body body createStudentRequest true "input required: student user_id, student_id, firstname, lastname, mobile"
// @Success 200 {object} studentResponse "ok"
// @Router /student [post]
func (server *Server) createStudent(ctx *gin.Context) {
	var req createStudentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	student, err := server.store.CreateStudent(ctx, db.CreateStudentParams{
		UserID:    req.UserID,
		StudentID: req.StudentID,
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
	ctx.JSON(http.StatusOK, studentResponse{
		UserID:    student.UserID,
		StudentID: student.StudentID,
		FirstName: student.FirstName,
		LastName:  student.LastName,
		Mobile:    student.Mobile,
		CreatedAt: student.CreatedAt,
	})
}

// listStudentsByName
// @Summary list created student
// @Description list created student
// @Tags Students
// @Accept	json
// @Produce  json
// @Param pageNumber query string true "page number"
// @Param pageSize query string true "page size"
// @Success 200 {object} []studentResponse "ok"
// @Router /student/list/name [get]
func (server *Server) listStudentsByName(ctx *gin.Context) {
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

	students, err := server.store.ListStudentsByName(ctx, db.ListStudentsByNameParams{
		Limit:  p.Limit(),
		Offset: p.Offset(),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	StudentsResponse := make([]studentResponse, 0)
	for _, student := range students {
		StudentsResponse = append(StudentsResponse, studentResponse{
			UserID:    student.UserID,
			StudentID: student.StudentID,
			FirstName: student.FirstName,
			LastName:  student.LastName,
			Mobile:    student.Mobile,
			CreatedAt: student.CreatedAt,
		})
	}
	ctx.JSON(http.StatusOK, StudentsResponse)
}

// listStudentsByStudentID
// @Summary list created student
// @Description list created student
// @Tags Students
// @Accept	json
// @Produce  json
// @Param pageNumber query string true "page number"
// @Param pageSize query string true "page size"
// @Success 200 {object} []studentResponse "ok"
// @Router /student/list/studentID [get]
func (server *Server) listStudentsByStudentID(ctx *gin.Context) {
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

	students, err := server.store.ListStudentsByStudentID(ctx, db.ListStudentsByStudentIDParams{
		Limit:  p.Limit(),
		Offset: p.Offset(),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	StudentsResponse := make([]studentResponse, 0)
	for _, student := range students {
		StudentsResponse = append(StudentsResponse, studentResponse{
			UserID:    student.UserID,
			StudentID: student.StudentID,
			FirstName: student.FirstName,
			LastName:  student.LastName,
			Mobile:    student.Mobile,
			CreatedAt: student.CreatedAt,
		})
	}
	ctx.JSON(http.StatusOK, StudentsResponse)
}
