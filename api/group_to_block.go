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

type createGroupToBlockRequest struct {
	AcademicYearName string `json:"academic_year_name"`
	GroupName        string `json:"group_name"`
	BlockName        string `json:"block_name"`
	PeriodName       string `json:"period_name"`
}

type groupToBlockResponse struct {
	AcademicYearName string    `json:"academic_year_name"`
	GroupName        string    `json:"group_name"`
	BlockName        string    `json:"block_name"`
	CreatedAt        time.Time `json:"created_at"`
}

// createGroupToBlock
// @Summary create new group to block
// @Description create new group to block
// @Tags GroupToBlock
// @Accept	json
// @Produce  json
// @Param body body createGroupToBlockRequest true "input required: academic year name, groupid, blockid"
// @Success 200 "ok"
// @Router /group_to_block [post]
func (server *Server) createGroupToBlock(ctx *gin.Context) {
	var req createGroupToBlockRequest
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

	period, err := server.store.GetPeriodByIndex(ctx, db.GetPeriodByIndexParams{
		AcademicYearID: academicYear.ID,
		Name:           req.PeriodName,
	})
	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	block, err := server.store.GetBlockByIndex(ctx, db.GetBlockByIndexParams{
		AcademicYearID: academicYear.ID,
		Period:         period.ID,
		Name:           req.BlockName,
	})
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

	_, err = server.store.CreateGroupToBlock(ctx, db.CreateGroupToBlockParams{
		AcademicYearID: academicYear.ID,
		BlockID:        block.ID,
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

	ctx.JSON(http.StatusOK, nil)
}

// listGroupToBlockByAcademicYear
// @Summary list groups and their blocks in that academic year
// @Description list groups and their blocks in that academic year
// @Tags GroupToBlock
// @Accept	json
// @Produce  json
// @Param academicYearName query string true "academic year name"
// @Success 200 {object} []groupToBlockResponse "ok"
// @Router /group_to_block/list/academic_year [get]
func (server *Server) listGroupToBlockByAcademicYear(ctx *gin.Context) {
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

	groupToBlocks, err := server.store.GetGroupToBlockByAcademicYearID(ctx, academicYear.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	groupToBlocksResponse := make([]groupToBlockResponse, 0)
	for _, groupToBlock := range groupToBlocks {
		block, _ := server.store.GetBlockByID(ctx, groupToBlock.BlockID)
		group, _ := server.store.GetGroupByID(ctx, groupToBlock.GroupID)
		groupToBlocksResponse = append(groupToBlocksResponse, groupToBlockResponse{
			AcademicYearName: academicYearName,
			BlockName:        block.Name,
			GroupName:        group.Name,
		})
	}
	ctx.JSON(http.StatusOK, groupToBlocksResponse)
}

// listGroupToBlockByGroupName
// @Summary list group_to_block of that block
// @Description list group_to_block of that block
// @Tags GroupToBlock
// @Accept	json
// @Produce  json
// @Param academicYearName query string true "academic year name"
// @Param groupName query string true "group name"
// @Success 200 {object} []groupToBlockResponse "ok"
// @Router /group_to_block/list/group [get]
func (server *Server) listGroupToBlockByGroupName(ctx *gin.Context) {
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

	groupToBlocks, err := server.store.GetGroupToBlockByGroupID(ctx, group.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	groupToBlocksResponse := make([]groupToBlockResponse, 0)
	for _, groupToBlock := range groupToBlocks {
		// TODO optimize SQL queries here
		block, _ := server.store.GetBlockByID(ctx, groupToBlock.BlockID)
		groupToBlocksResponse = append(groupToBlocksResponse, groupToBlockResponse{
			AcademicYearName: academicYear.Name,
			BlockName:        block.Name,
			GroupName:        group.Name,
		})
	}
	ctx.JSON(http.StatusOK, groupToBlocksResponse)
}

// listGroupToBlockByGroupName
// @Summary list groups in that block
// @Description list groups in that block
// @Tags GroupToBlock
// @Accept	json
// @Produce  json
// @Param academicYearName query string true "academic year name"
// @Param periodName query string true "period name"
// @Param blockName query string true "block name"
// @Success 200 {object} []groupToBlockResponse "ok"
// @Router /group_to_block/list/block [get]
func (server *Server) listGroupToBlockByBlockName(ctx *gin.Context) {
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

	periodName := ctx.Query("periodName")
	period, err := server.store.GetPeriodByIndex(ctx, db.GetPeriodByIndexParams{
		AcademicYearID: academicYear.ID,
		Name:           periodName,
	})
	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	blockName := ctx.Query("blockName")
	block, err := server.store.GetBlockByIndex(ctx, db.GetBlockByIndexParams{
		AcademicYearID: academicYear.ID,
		Period:         period.ID,
		Name:           blockName,
	})
	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	groupToBlocks, err := server.store.GetGroupToBlockByBlockID(ctx, block.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	groupToBlocksResponse := make([]groupToBlockResponse, 0)
	for _, groupToBlock := range groupToBlocks {
		// TODO optimize SQL queries here
		group, _ := server.store.GetGroupByID(ctx, groupToBlock.GroupID)
		groupToBlocksResponse = append(groupToBlocksResponse, groupToBlockResponse{
			AcademicYearName: academicYear.Name,
			BlockName:        blockName,
			GroupName:        group.Name,
		})
	}
	ctx.JSON(http.StatusOK, groupToBlocksResponse)
}
