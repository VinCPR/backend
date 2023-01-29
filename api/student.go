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

type createStudentRequest struct {
	StudentID string `json:"student_id" binding:"required"`
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Mobile    string `json:"mobile" binding:"required"`
	Biography string `json:"biography" binding:"required"`
	Image     string `json:"image" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=8,max=64"`
}

type studentResponse struct {
	Email     string    `json:"email" binding:"required,email"`
	StudentID string    `json:"student_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Mobile    string    `json:"mobile"`
	Biography string    `json:"biography"`
	Image     string    `json:"image"`
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

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	user, err := server.store.CreateUser(ctx, db.CreateUserParams{
		Email:          req.Email,
		HashedPassword: hashedPassword,
		RoleName:       "student",
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

	student, err := server.store.CreateStudent(ctx, db.CreateStudentParams{
		UserID:    user.ID,
		StudentID: req.StudentID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Mobile:    req.Mobile,
		Biography: req.Biography,
		Image:     req.Image,
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
		Email:     req.Email,
		StudentID: student.StudentID,
		FirstName: student.FirstName,
		LastName:  student.LastName,
		Mobile:    student.Mobile,
		Biography: student.Biography,
		Image:     student.Image,
		CreatedAt: student.CreatedAt,
	})
}

// getStudentInfoByEmail
// @Summary provide detail of a student given email address
// @Description provide detail of a student given email address
// @Tags Students
// @Accept	json
// @Produce  json
// @Param email query string true "student email address"
// @Success 200 {object} studentResponse "ok"
// @Router /student/info [get]
func (server *Server) getStudentInfoByEmail(ctx *gin.Context) {
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
	student, err := server.store.GetStudentByUserId(ctx, user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, studentResponse{
		Email:     email,
		StudentID: student.StudentID,
		FirstName: student.FirstName,
		LastName:  student.LastName,
		Mobile:    student.Mobile,
		Biography: student.Biography,
		Image:     student.Image,
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

	studentsResponse := make([]studentResponse, 0)
	for _, student := range students {
		userInfo, err := server.store.GetUserByID(ctx, student.UserID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		studentsResponse = append(studentsResponse, studentResponse{
			Email:     userInfo.Email,
			StudentID: student.StudentID,
			FirstName: student.FirstName,
			LastName:  student.LastName,
			Mobile:    student.Mobile,
			Biography: student.Biography,
			Image:     student.Image,
			CreatedAt: student.CreatedAt,
		})
	}
	ctx.JSON(http.StatusOK, studentsResponse)
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

	studentsResponse := make([]studentResponse, 0)
	for _, student := range students {
		userInfo, err := server.store.GetUserByID(ctx, student.UserID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		studentsResponse = append(studentsResponse, studentResponse{
			Email:     userInfo.Email,
			StudentID: student.StudentID,
			FirstName: student.FirstName,
			LastName:  student.LastName,
			Mobile:    student.Mobile,
			Biography: student.Biography,
			Image:     student.Image,
			CreatedAt: student.CreatedAt,
		})
	}
	ctx.JSON(http.StatusOK, studentsResponse)
}
