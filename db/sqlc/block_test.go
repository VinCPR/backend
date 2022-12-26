package db

import (
	"context"
	"sort"
	"testing"

	"github.com/VinCPR/backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomBlock(t *testing.T, academicYearID int64) Block {
	arg := CreateBlockParams{
		AcademicYearID: academicYearID,
		Name:           util.RandomName(),
		Period:         createRandomPeriod(t, academicYearID).ID,
	}
	block, err := testQueries.CreateBlock(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, block)

	require.Equal(t, arg.AcademicYearID, block.AcademicYearID)
	require.Equal(t, arg.Name, block.Name)
	require.Equal(t, arg.Period, block.Period)

	require.NotZero(t, block.ID)
	return block
}

func TestListBlocksByAcademicYear(t *testing.T) {
	var (
		academicYear  = createRandomAcademicYear(t)
		createdBlocks = make([]Block, 0)
		n             = 5
	)
	for i := 0; i < n; i++ {
		createdBlocks = append(createdBlocks, createRandomBlock(t, academicYear.ID))
	}

	blocks, err := testQueries.ListBlocksByAcademicYear(context.Background(), academicYear.ID)
	require.NoError(t, err)
	require.NotEmpty(t, blocks)

	sort.Slice(createdBlocks, func(i, j int) bool {
		return createdBlocks[i].Period < createdBlocks[j].Period ||
			(createdBlocks[i].Period == createdBlocks[j].Period && createdBlocks[i].ID < createdBlocks[j].ID)
	})

	for i := 0; i < n; i++ {
		require.Equal(t, blocks[i].AcademicYearID, academicYear.ID)
		require.Equal(t, blocks[i].Name, blocks[i].Name)
		require.Equal(t, blocks[i].Period, blocks[i].Period)
	}
}
