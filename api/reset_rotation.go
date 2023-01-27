package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/rs/zerolog/log"

	db "github.com/VinCPR/backend/db/sqlc"
)

type resetRotationRequest struct {
	AcademicYearName string `json:"academic_year_name"`
}

// designRotation
// @Summary reset clinical rotation
// @Description reset all rotation events of an academic year
// @Tags ClinicalRotation
// @Accept	json
// @Produce  json
// @Param body body resetRotationRequest true "input required: academic year name"
// @Success 200 "ok"
// @Router /rotation/reset [post]
func (server *Server) resetRotation(ctx *gin.Context) {
	var req resetRotationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	academicYear, err := server.store.GetAcademicYearByName(ctx, req.AcademicYearName)
	if err != nil {
		if err == pgx.ErrNoRows {
			log.Info().Msgf("cannot find academic year %v", req.AcademicYearName)
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	if err = processResetRotation(ctx, server.store, req, academicYear); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

func processResetRotation(ctx context.Context, store *db.Store, req resetRotationRequest, academicYear db.AcademicYear) error {
	var err error
	// create db transaction (rollback if any queries among transaction fail)
	tx, err := store.Db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	qtx := store.WithTx(tx)

	if err = qtx.DeleteRotationEventsByAcademicYear(ctx, academicYear.ID); err != nil {
		return err
	}

	if err = qtx.DeleteGroupToBlocksByAcademicYear(ctx, academicYear.ID); err != nil {
		return err
	}

	if err = qtx.DeleteBlocksByAcademicYear(ctx, academicYear.ID); err != nil {
		return err
	}

	if err = qtx.DeletePeriodsByAcademicYear(ctx, academicYear.ID); err != nil {
		return err
	}

	return tx.Commit(ctx)
}
