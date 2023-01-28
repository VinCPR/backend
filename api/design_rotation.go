package api

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/rs/zerolog/log"

	db "github.com/VinCPR/backend/db/sqlc"
)

type designRotationRequest struct {
	AcademicYearName string `json:"academic_year_name" binding:"required"`
	GroupsPerBlock   int    `json:"groups_per_block" binding:"required"`
	NumberOfPeriod   int    `json:"number_of_period" binding:"required"`
	WeeksPerPeriod   int    `json:"weeks_per_period" binding:"required"`

	Periods []periodInfoRequest `json:"periods" binding:"required"`
	Blocks  []blockInfoRequest  `json:"blocks" binding:"required"`
}

type periodInfoRequest struct {
	PeriodName string    `json:"period_name" binding:"required"`
	StartDate  time.Time `json:"start_date" form:"start_date" binding:"required" time_format:"2006-01-02"`
	// EndDate = StartDate + WeeksPerPeriod
	// EndDate    time.Time `json:"end_date" form:"end_date" binding:"required" time_format:"2006-01-02"`
}

type blockInfoRequest struct {
	BlockName     string                   `json:"block_name"`
	GroupCalendar [][]specialtyInfoRequest `json:"group_calendar"`
}

type specialtyInfoRequest struct {
	SpecialtyName string                `json:"specialty_name"`
	Hospitals     []hospitalInfoRequest `json:"hospitals"`
}

type hospitalInfoRequest struct {
	HospitalName string               `json:"hospital_name"`
	Services     []serviceInfoRequest `json:"services"`
}

type serviceInfoRequest struct {
	ServiceName    string `json:"service_name"`
	DurationInWeek int    `json:"duration_in_week"`
}

// designRotation
// @Summary design clinical rotation
// @Description provide templates of each period to auto generate calendar for all students
// @Tags ClinicalRotation
// @Accept	json
// @Produce  json
// @Param body body designRotationRequest true "input required: academic year name, templates for each period"
// @Success 200 "ok"
// @Router /rotation/design [post]
func (server *Server) designRotation(ctx *gin.Context) {
	var req designRotationRequest
	var err error
	if err = ctx.ShouldBindJSON(&req); err != nil {
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

	if err = validateDesignRotationRequest(req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err = processDesignRotation(ctx, server.store, req, academicYear); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

func validateDesignRotationRequest(req designRotationRequest) error {
	if len(req.Periods) != req.NumberOfPeriod {
		return fmt.Errorf("mismatch length of periods")
	}
	if len(req.Periods) != len(req.Blocks) {
		return fmt.Errorf("number of periods should be equal to number of blocks")
	}
	for _, block := range req.Blocks {
		if len(block.GroupCalendar) != req.GroupsPerBlock {
			return fmt.Errorf("match number of groups per block")
		}
		// TODO: validate that sum of weeks in a block = WeeksPerPeriod
	}
	for i := 1; i < req.NumberOfPeriod; i++ {
		endDateOfPreviousPeriod := req.Periods[i-1].StartDate.AddDate(0, 0, 7*req.WeeksPerPeriod)
		if req.Periods[i].StartDate.Before(endDateOfPreviousPeriod) {
			return fmt.Errorf("period overlapping")
		}
	}
	return nil
}

func processDesignRotation(ctx context.Context, store *db.Store, req designRotationRequest, academicYear db.AcademicYear) error {
	var err error
	// create db transaction (rollback if any queries among transaction fail)
	tx, err := store.Db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	qtx := store.WithTx(tx)

	// create periods
	periods, err := processCreatePeriods(ctx, qtx, req.Periods, req.WeeksPerPeriod, academicYear.ID)
	if err != nil {
		return err
	}

	blocksForEachPeriod, err := processCreateBlocks(ctx, qtx, periods, req.Blocks, academicYear.ID)
	if err != nil {
		return err
	}

	// list groups in an academic year
	groups, err := qtx.ListGroupsByName(ctx, academicYear.ID)
	if err != nil {
		return err
	}

	// DivideGroup 1,2,3,4 -> Block 1. 5,6,7,8 -> Block 2. 9,10,11,12 -> Block 3. 13,14,15-> Block 4
	if err = processFillGroupToBlocks(ctx, qtx, blocksForEachPeriod, groups, req.Periods, req.Blocks, req.GroupsPerBlock, academicYear.ID); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func processCreatePeriods(ctx context.Context, qtx *db.Queries, periodsInfo []periodInfoRequest, weeksPerPeriod int, academicYearID int64) ([]db.Period, error) {
	var periods []db.Period
	for _, periodInfo := range periodsInfo {
		period, err := qtx.CreatePeriod(ctx, db.CreatePeriodParams{
			AcademicYearID: academicYearID,
			Name:           periodInfo.PeriodName,
			StartDate:      periodInfo.StartDate,
			EndDate:        periodInfo.StartDate.AddDate(0, 0, 7*weeksPerPeriod),
		})
		if err != nil {
			return nil, err
		}
		periods = append(periods, period)
	}
	return periods, nil
}

func processCreateBlocks(ctx context.Context, qtx *db.Queries, periods []db.Period, blocksInfo []blockInfoRequest, academicYearID int64) ([][]db.Block, error) {
	if len(periods) != len(blocksInfo) {
		return nil, fmt.Errorf("number of periods should be equal to number of blocks")
	}
	var blocksForEachPeriod [][]db.Block
	for _, period := range periods {
		var blocks []db.Block
		for _, blockInfo := range blocksInfo {
			block, err := qtx.CreateBlock(ctx, db.CreateBlockParams{
				AcademicYearID: academicYearID,
				Name:           blockInfo.BlockName,
				Period:         period.ID,
			})
			if err != nil {
				return nil, err
			}
			blocks = append(blocks, block)
		}
		blocksForEachPeriod = append(blocksForEachPeriod, blocks)
	}
	return blocksForEachPeriod, nil
}

func processFillGroupToBlocks(ctx context.Context, qtx *db.Queries, blocksForEachPeriod [][]db.Block, groups []db.Group,
	periodsInfo []periodInfoRequest, blocksInfo []blockInfoRequest, groupsPerBlock int, academicYearID int64) error {
	// numberOfPeriods = numberOfBlockPerPeriod
	// -> blocksForEachPeriod has size (numberOfPeriods, numberOfPeriods)
	// numberOfPeriods := len(blocksForEachPeriod)
	services, err := getServiceID(ctx, qtx, blocksInfo)
	if err != nil {
		return err
	}
	// TODO remove nested loop here
	for i, group := range groups {
		// blocks: all blocks in period j
		for j, blocks := range blocksForEachPeriod {
			var blockInfoID = (i/groupsPerBlock + j) % len(blocks)
			_, err = qtx.CreateGroupToBlock(ctx, db.CreateGroupToBlockParams{
				AcademicYearID: academicYearID,
				GroupID:        group.ID,
				BlockID:        blocks[blockInfoID].ID,
			})
			if err != nil {
				return err
			}
			err = processCreateClinicalRotationEvent(ctx, qtx, periodsInfo[j].StartDate, blocksInfo[blockInfoID].GroupCalendar[i%groupsPerBlock], services, group.ID, academicYearID)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func processCreateClinicalRotationEvent(ctx context.Context, qtx *db.Queries, startDate time.Time, groupCalendar []specialtyInfoRequest, services map[string]int64, groupID, academicYearID int64) error {
	for _, specialtyInfo := range groupCalendar {
		for _, hospitalInfo := range specialtyInfo.Hospitals {
			for _, serviceInfo := range hospitalInfo.Services {
				_, err := qtx.CreateRotationEvent(ctx, db.CreateRotationEventParams{
					AcademicYearID: academicYearID,
					GroupID:        groupID,
					ServiceID:      services[join(specialtyInfo.SpecialtyName, hospitalInfo.HospitalName, serviceInfo.ServiceName)],
					StartDate:      startDate,
					EndDate:        startDate.AddDate(0, 0, serviceInfo.DurationInWeek*7-1),
				})
				if err != nil {
					return err
				}
				startDate = startDate.AddDate(0, 0, serviceInfo.DurationInWeek*7)
			}
		}
	}
	return nil
}

func getServiceID(ctx context.Context, qtx *db.Queries, blocksInfo []blockInfoRequest) (map[string]int64, error) {
	var (
		specialty db.Specialty
		hospital  db.Hospital
		service   db.Service

		servicesByName = make(map[string]int64)
		err            error
	)
	for _, blockInfo := range blocksInfo {
		for _, groupInfo := range blockInfo.GroupCalendar {
			for _, specialtyInfo := range groupInfo {
				specialty, err = qtx.GetSpecialtyByName(ctx, specialtyInfo.SpecialtyName)
				if err != nil {
					return nil, err
				}
				for _, hospitalInfo := range specialtyInfo.Hospitals {
					hospital, err = qtx.GetHospitalByName(ctx, hospitalInfo.HospitalName)
					if err != nil {
						return nil, err
					}
					for _, serviceInfo := range hospitalInfo.Services {
						service, err = qtx.GetServiceByIndex(ctx, db.GetServiceByIndexParams{
							SpecialtyID: specialty.ID,
							HospitalID:  hospital.ID,
							Name:        serviceInfo.ServiceName,
						})
						if err != nil {
							return nil, err
						}
						servicesByName[join(specialty.Name, hospital.Name, service.Name)] = service.ID
					}
				}
			}
		}
	}
	return servicesByName, nil
}

func join(args ...string) string {
	return strings.Join(args, "___")
}
