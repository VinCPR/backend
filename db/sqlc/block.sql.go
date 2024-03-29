// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: block.sql

package db

import (
	"context"
)

const createBlock = `-- name: CreateBlock :one
INSERT INTO "block" (
  academic_year_id,
  name,
  period
) VALUES (
  $1, $2, $3
) RETURNING id, academic_year_id, name, period, created_at
`

type CreateBlockParams struct {
	AcademicYearID int64  `json:"academic_year_id"`
	Name           string `json:"name"`
	Period         int64  `json:"period"`
}

func (q *Queries) CreateBlock(ctx context.Context, arg CreateBlockParams) (Block, error) {
	row := q.db.QueryRow(ctx, createBlock, arg.AcademicYearID, arg.Name, arg.Period)
	var i Block
	err := row.Scan(
		&i.ID,
		&i.AcademicYearID,
		&i.Name,
		&i.Period,
		&i.CreatedAt,
	)
	return i, err
}

const deleteBlocksByAcademicYear = `-- name: DeleteBlocksByAcademicYear :exec
DELETE FROM "block" WHERE "academic_year_id" = $1
`

func (q *Queries) DeleteBlocksByAcademicYear(ctx context.Context, academicYearID int64) error {
	_, err := q.db.Exec(ctx, deleteBlocksByAcademicYear, academicYearID)
	return err
}

const getBlockByID = `-- name: GetBlockByID :one
SELECT id, academic_year_id, name, period, created_at FROM "block"
WHERE "id" = $1 LIMIT 1
`

func (q *Queries) GetBlockByID(ctx context.Context, id int64) (Block, error) {
	row := q.db.QueryRow(ctx, getBlockByID, id)
	var i Block
	err := row.Scan(
		&i.ID,
		&i.AcademicYearID,
		&i.Name,
		&i.Period,
		&i.CreatedAt,
	)
	return i, err
}

const getBlockByIndex = `-- name: GetBlockByIndex :one
SELECT id, academic_year_id, name, period, created_at FROM "block"
WHERE "academic_year_id" = $1 AND "period" = $2 AND "name" = $3
LIMIT 1
`

type GetBlockByIndexParams struct {
	AcademicYearID int64  `json:"academic_year_id"`
	Period         int64  `json:"period"`
	Name           string `json:"name"`
}

func (q *Queries) GetBlockByIndex(ctx context.Context, arg GetBlockByIndexParams) (Block, error) {
	row := q.db.QueryRow(ctx, getBlockByIndex, arg.AcademicYearID, arg.Period, arg.Name)
	var i Block
	err := row.Scan(
		&i.ID,
		&i.AcademicYearID,
		&i.Name,
		&i.Period,
		&i.CreatedAt,
	)
	return i, err
}

const listBlocksByAcademicYear = `-- name: ListBlocksByAcademicYear :many
SELECT id, academic_year_id, name, period, created_at FROM "block"
WHERE "academic_year_id" = $1
ORDER BY "period"
`

func (q *Queries) ListBlocksByAcademicYear(ctx context.Context, academicYearID int64) ([]Block, error) {
	rows, err := q.db.Query(ctx, listBlocksByAcademicYear, academicYearID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Block{}
	for rows.Next() {
		var i Block
		if err := rows.Scan(
			&i.ID,
			&i.AcademicYearID,
			&i.Name,
			&i.Period,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
