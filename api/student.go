package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	db "github.com/VinCPR/backend/db/sqlc"
)

type StudentResponse struct {
	UserID    int64     `json:"user_id"`
	StudentID string    `json:"student_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Mobile    string    `json:"mobile"`
	CreatedAt time.Time `json:"created_at"`
}

// listStudentsByName
// @Summary list created student
// @Description list created student
// @Tags Students
// @Accept	json
// @Produce  json
// @Param pageNumber query string true "page number"
// @Param pageSize query string true "page size"
// @Success 200 {object} []StudentResponse "ok"
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

	StudentsResponse := make([]StudentResponse, 0)
	for _, student := range students {
		StudentsResponse = append(StudentsResponse, StudentResponse{
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
// @Success 200 {object} []StudentResponse "ok"
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

	StudentsResponse := make([]StudentResponse, 0)
	for _, student := range students {
		StudentsResponse = append(StudentsResponse, StudentResponse{
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
