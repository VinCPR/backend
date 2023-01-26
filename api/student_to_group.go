package api

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v4"

	db "github.com/VinCPR/backend/db/sqlc"
)

type createStudentToGroupRequest struct {
	AcademicYearName string `json:"academic_year_name" binding:"required"`
	StudentID        string `json:"student_id" binding:"required"`
	GroupName        string `json:"group_name" binding:"required"`
}

type studentToGroupResponse struct {
	AcademicYearName string    `json:"academic_year_name"`
	StudentID        string    `json:"student_id"`
	GroupName        string    `json:"group_name"`
	CreatedAt        time.Time `json:"created_at"`
}

// createStudentToGroup
// @Summary create new student to group
// @Description create new student to group
// @Tags StudentToGroup
// @Accept	json
// @Produce  json
// @Param body body createStudentToGroupRequest true "input required: academic year name, studentid, groupid"
// @Success 200 {object} studentToGroupResponse "ok"
// @Router /student_to_group [post]
func (server *Server) createStudentToGroup(ctx *gin.Context) {
	var req createStudentToGroupRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	academicYear, err := server.store.GetAcademicYearByName(ctx, req.AcademicYearName)

	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	student, err := server.store.GetStudentByStudentId(ctx, req.StudentID)
	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	group, err := server.store.GetGroupByIndex(ctx, db.GetGroupByIndexParams{
		AcademicYearID: academicYear.ID,
		Name:           req.GroupName,
	})
	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	studentToGroup, err := server.store.CreateStudentToGroup(ctx, db.CreateStudentToGroupParams{
		AcademicYearID: academicYear.ID,
		StudentID:      student.ID,
		GroupID:        group.ID,
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
	ctx.JSON(http.StatusOK, studentToGroupResponse{
		CreatedAt: studentToGroup.CreatedAt,
	})
}

// listStudentToGroupByAcademicYear
// @Summary list students and their groups in that academic year
// @Description list students and their groups in that academic year
// @Tags StudentToGroup
// @Accept	json
// @Produce  json
// @Param academicYearName query string true "academic year name"
// @Success 200 {object} []studentToGroupResponse "ok"
// @Router /student_to_group/list/academic_year [get]
func (server *Server) listStudentToGroupByAcademicYear(ctx *gin.Context) {
	academicYearName := ctx.Query("academicYearName")

	academicYear, err := server.store.GetAcademicYearByName(ctx, academicYearName)
	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	studentToGroups, err := server.store.GetStudentToGroupByAcademicYearID(ctx, academicYear.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	studentToGroupsResponse := make([]studentToGroupResponse, 0)
	for _, studentToGroup := range studentToGroups {
		var student db.Student
		var group db.Group

		student, err = server.store.GetStudentByID(ctx, studentToGroup.StudentID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		group, err = server.store.GetGroupByID(ctx, studentToGroup.GroupID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		studentToGroupsResponse = append(studentToGroupsResponse, studentToGroupResponse{
			AcademicYearName: academicYearName,
			StudentID:        student.StudentID,
			GroupName:        group.Name,
			CreatedAt:        studentToGroup.CreatedAt,
		})
	}
	ctx.JSON(http.StatusOK, studentToGroupsResponse)
}

// listStudentToGroupByGroupID
// @Summary list students of that group
// @Description list students of that group
// @Tags StudentToGroup
// @Accept	json
// @Produce  json
// @Param academicYearName query string true "academic year name"
// @Param groupName query string true "group name"
// @Success 200 {object} []studentToGroupResponse "ok"
// @Router /student_to_group/list/group [get]
func (server *Server) listStudentToGroupByGroupID(ctx *gin.Context) {
	academicYearName := ctx.Query("academicYearName")
	academicYear, err := server.store.GetAcademicYearByName(ctx, academicYearName)
	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	groupName := ctx.Query("groupName")
	group, err := server.store.GetGroupByIndex(ctx, db.GetGroupByIndexParams{
		AcademicYearID: academicYear.ID,
		Name:           groupName,
	})
	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	studentToGroups, err := server.store.GetStudentToGroupByGroupID(ctx, group.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	studentToGroupsResponse := make([]studentToGroupResponse, 0)
	for _, studentToGroup := range studentToGroups {
		var student db.Student
		student, err = server.store.GetStudentByID(ctx, studentToGroup.StudentID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		studentToGroupsResponse = append(studentToGroupsResponse, studentToGroupResponse{
			AcademicYearName: academicYear.Name,
			StudentID:        student.StudentID,
			GroupName:        group.Name,
			CreatedAt:        studentToGroup.CreatedAt,
		})
	}
	ctx.JSON(http.StatusOK, studentToGroupsResponse)
}

// listStudentToGroupByStudentID
// @Summary list groups of that student
// @Description list groups of that student
// @Tags StudentToGroup
// @Accept	json
// @Produce  json
// @Param studentID query string true "studentID"
// @Success 200 {object} []studentToGroupResponse "ok"
// @Router /student_to_group/list/student [get]
func (server *Server) listStudentToGroupByStudentID(ctx *gin.Context) {
	studentID := ctx.Query("studentID")

	student, err := server.store.GetStudentByStudentId(ctx, studentID)
	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	studentToGroups, err := server.store.GetStudentToGroupByStudentID(ctx, student.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	studentToGroupsResponse := make([]studentToGroupResponse, 0)
	for _, studentToGroup := range studentToGroups {
		var academicYear db.AcademicYear
		var group db.Group

		academicYear, err = server.store.GetAcademicYearByID(ctx, studentToGroup.AcademicYearID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		group, err = server.store.GetGroupByID(ctx, studentToGroup.GroupID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		studentToGroupsResponse = append(studentToGroupsResponse, studentToGroupResponse{
			AcademicYearName: academicYear.Name,
			StudentID:        studentID,
			GroupName:        group.Name,
			CreatedAt:        studentToGroup.CreatedAt,
		})
	}
	ctx.JSON(http.StatusOK, studentToGroupsResponse)
}
