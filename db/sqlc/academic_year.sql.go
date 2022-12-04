// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: academic_year.sql

package db

import (
	"context"
	"time"
)

const createAcademicYear = `-- name: CreateAcademicYear :one
INSERT INTO "academic_year" (
    "name", "start_date", "end_date"
) VALUES(
    $1, $2, $3
) RETURNING id, name, start_date, end_date
`

type CreateAcademicYearParams struct {
	Name      string    `json:"name"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

func (q *Queries) CreateAcademicYear(ctx context.Context, arg CreateAcademicYearParams) (AcademicYear, error) {
	row := q.db.QueryRowContext(ctx, createAcademicYear, arg.Name, arg.StartDate, arg.EndDate)
	var i AcademicYear
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.StartDate,
		&i.EndDate,
	)
	return i, err
}

const getAcademicYearByName = `-- name: GetAcademicYearByName :one
SELECT id, name, start_date, end_date FROM "academic_year"
WHERE "name" = $1 LIMIT 1
`

func (q *Queries) GetAcademicYearByName(ctx context.Context, name string) (AcademicYear, error) {
	row := q.db.QueryRowContext(ctx, getAcademicYearByName, name)
	var i AcademicYear
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.StartDate,
		&i.EndDate,
	)
	return i, err
}

const listUsersByName = `-- name: ListUsersByName :many
SELECT id, email, hashed_password, role_name, created_at FROM "user"
ORDER BY "name"
LIMIT $1
OFFSET $2
`

type ListUsersByNameParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListUsersByName(ctx context.Context, arg ListUsersByNameParams) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsersByName, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.HashedPassword,
			&i.RoleName,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
