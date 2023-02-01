package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomGroupToBlock(t *testing.T, academic_year AcademicYear, group Group, block Block) GroupToBlock {

	arg := CreateGroupToBlockParams{
		AcademicYearID: academic_year.ID,
		GroupID:        group.ID,
		BlockID:        block.ID,
	}

	group_to_block, err := testQueries.CreateGroupToBlock(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, group_to_block)

	require.Equal(t, arg.GroupID, group_to_block.GroupID)
	require.Equal(t, arg.AcademicYearID, group_to_block.AcademicYearID)
	require.Equal(t, arg.BlockID, group_to_block.BlockID)

	require.NotZero(t, group_to_block.ID)
	require.NotZero(t, group_to_block.CreatedAt)

	return group_to_block
}

func TestCreateGroupToBlock(t *testing.T) {
	academic_year := createRandomAcademicYear(t)
	group := createRandomGroup(t, academic_year)
	block := createRandomBlock(t, academic_year.ID)
	createRandomGroupToBlock(t, academic_year, group, block)
}

func TestGetGroupToBlockByAcademicYearID(t *testing.T) {
	academic_year := createRandomAcademicYear(t)
	group := createRandomGroup(t, academic_year)
	block := createRandomBlock(t, academic_year.ID)
	group_to_block1 := createRandomGroupToBlock(t, academic_year, group, block)
	group_to_block2, err := testQueries.GetGroupToBlockByAcademicYearID(context.Background(), group_to_block1.AcademicYearID)
	require.NoError(t, err)
	require.NotEmpty(t, group_to_block2)

	for i := 0; i < len(group_to_block2); i++ {
		require.Equal(t, group_to_block1.ID, group_to_block2[i].ID)
		require.Equal(t, group_to_block1.AcademicYearID, group_to_block2[i].AcademicYearID)
		require.Equal(t, group_to_block1.BlockID, group_to_block2[i].BlockID)
		require.Equal(t, group_to_block1.GroupID, group_to_block2[i].GroupID)
		require.WithinDuration(t, group_to_block1.CreatedAt, group_to_block2[i].CreatedAt, time.Second)
	}
}

func TestGetGroupToBlockByGroupID(t *testing.T) {
	academic_year := createRandomAcademicYear(t)
	group := createRandomGroup(t, academic_year)
	block := createRandomBlock(t, academic_year.ID)
	group_to_block1 := createRandomGroupToBlock(t, academic_year, group, block)
	group_to_block2, err := testQueries.GetGroupToBlockByGroupID(context.Background(), group_to_block1.GroupID)
	require.NoError(t, err)
	require.NotEmpty(t, group_to_block2)

	for i := 0; i < len(group_to_block2); i++ {
		require.Equal(t, group_to_block1.ID, group_to_block2[i].ID)
		require.Equal(t, group_to_block1.AcademicYearID, group_to_block2[i].AcademicYearID)
		require.Equal(t, group_to_block1.BlockID, group_to_block2[i].BlockID)
		require.Equal(t, group_to_block1.GroupID, group_to_block2[i].GroupID)
		require.WithinDuration(t, group_to_block1.CreatedAt, group_to_block2[i].CreatedAt, time.Second)
	}
}

func TestGetGroupToBlockByBlockID(t *testing.T) {
	academic_year := createRandomAcademicYear(t)
	group := createRandomGroup(t, academic_year)
	block := createRandomBlock(t, academic_year.ID)
	group_to_block1 := createRandomGroupToBlock(t, academic_year, group, block)
	group_to_block2, err := testQueries.GetGroupToBlockByBlockID(context.Background(), group_to_block1.BlockID)
	require.NoError(t, err)
	require.NotEmpty(t, group_to_block2)

	for i := 0; i < len(group_to_block2); i++ {
		require.Equal(t, group_to_block1.ID, group_to_block2[i].ID)
		require.Equal(t, group_to_block1.AcademicYearID, group_to_block2[i].AcademicYearID)
		require.Equal(t, group_to_block1.BlockID, group_to_block2[i].BlockID)
		require.Equal(t, group_to_block1.GroupID, group_to_block2[i].GroupID)
		require.WithinDuration(t, group_to_block1.CreatedAt, group_to_block2[i].CreatedAt, time.Second)
	}
}

func TestDeleteGroupToBlocksByAcademicYear(t *testing.T) {
	academic_year := createRandomAcademicYear(t)
	group := createRandomGroup(t, academic_year)
	block := createRandomBlock(t, academic_year.ID)
	group_to_block1 := createRandomGroupToBlock(t, academic_year, group, block)
	group_to_block2 := testQueries.DeleteGroupToBlocksByAcademicYear(context.Background(), group_to_block1.AcademicYearID)
	require.Empty(t, group_to_block2)
}
