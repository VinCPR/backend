package api

import (
	"errors"
	"net/http"
	"time"

	db "github.com/VinCPR/backend/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v4"
)

type createStudentToGroupRequest struct {
	AcademicYearName string `json:"academic_year_name"`
	StudentName      string `json:"student_name`
	GroupName        string `json:"group_name"`
}
type studentToGroupResponse struct {
	AcademicYearName string    `json:"academic_year_name"`
	StudentName      string    `json:"student_name`
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

	//academicYearName := ctx.Query("academicYearName")

	academicYear, err := server.store.GetAcademicYearByName(ctx, req.AcademicYearName)

	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	//blockName := ctx.Query("blockName")

	student, err := server.store.GetBlockByName(ctx, req.StudentName)
	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	//groupName := ctx.Query("groupName")

	group, err := server.store.GetGroupByName(ctx, req.GroupName)
	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
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
	}

	studentToGroups, err := server.store.GetStudentToGroupByAcademicYearID(ctx, academicYear.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	studentToGroupsResponse := make([]studentToGroupResponse, 0)
	for _, studentToGroup := range studentToGroups {
		student, _ := server.store.GetBlockByID(ctx, studentToGroup.StudentID)
		group, _ := server.store.GetGroupByID(ctx, studentToGroup.GroupID)
		studentToGroupsResponse = append(studentToGroupsResponse, studentToGroupResponse{
			AcademicYearName: academicYearName,
			StudentName:      student.Name,
			GroupName:        group.Name,
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
// @Param groupName query string true "group name"
// @Success 200 {object} []studentToGroupResponse "ok"
// @Router /student_to_group/list/group [get]
func (server *Server) listStudentToGroupByGroupID(ctx *gin.Context) {
	groupName := ctx.Query("groupName")

	group, err := server.store.GetGroupByName(ctx, groupName)
	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	studentToGroups, err := server.store.GetStudentToGroupByGroupID(ctx, group.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	studentToGroupsResponse := make([]studentToGroupResponse, 0)
	for _, studentToGroup := range studentToGroups {
		student, _ := server.store.GetBlockByID(ctx, studentToGroup.StudentID)
		academicYear, _ := server.store.GetAcademicYearByID(ctx, studentToGroup.AcademicYearID)
		studentToGroupsResponse = append(studentToGroupsResponse, studentToGroupResponse{
			AcademicYearName: academicYear.Name,
			StudentName:      student.Name,
			GroupName:        group.Name,
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
	}

	studentToGroups, err := server.store.GetStudentToGroupByStudentID(ctx, student.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	studentToGroupsResponse := make([]studentToGroupResponse, 0)
	for _, studentToGroup := range studentToGroups {
		academicYear, _ := server.store.GetAcademicYearByID(ctx, studentToGroup.AcademicYearID)
		group, _ := server.store.GetGroupByID(ctx, studentToGroup.GroupID)
		studentToGroupsResponse = append(studentToGroupsResponse, studentToGroupResponse{
			AcademicYearName: academicYear.Name,
			StudentName:      student.FirstName + " " + student.LastName,
			GroupName:        group.Name,
		})
	}
	ctx.JSON(http.StatusOK, studentToGroupsResponse)
}
