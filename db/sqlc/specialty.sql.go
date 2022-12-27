// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: specialty.sql

package db

import (
	"context"
)

const createSpecialty = `-- name: CreateSpecialty :one
INSERT INTO "specialty" (
    name,
    description
) VALUES (
    $1 , $2 
) RETURNING id, name, description, created_at
`

type CreateSpecialtyParams struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (q *Queries) CreateSpecialty(ctx context.Context, arg CreateSpecialtyParams) (Specialty, error) {
	row := q.db.QueryRow(ctx, createSpecialty, arg.Name, arg.Description)
	var i Specialty
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
	)
	return i, err
}

const getSpecialtyByName = `-- name: GetSpecialtyByName :one
SELECT id, name, description, created_at FROM "specialty"
WHERE name = $1 LIMIT 1
`

func (q *Queries) GetSpecialtyByName(ctx context.Context, name string) (Specialty, error) {
	row := q.db.QueryRow(ctx, getSpecialtyByName, name)
	var i Specialty
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
	)
	return i, err
}

const listSpecialtiesByName = `-- name: ListSpecialtiesByName :many
SELECT id, name, description, created_at FROM "specialty"
ORDER BY name 
LIMIT $1
OFFSET $2
`

type ListSpecialtiesByNameParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListSpecialtiesByName(ctx context.Context, arg ListSpecialtiesByNameParams) ([]Specialty, error) {
	rows, err := q.db.Query(ctx, listSpecialtiesByName, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Specialty{}
	for rows.Next() {
		var i Specialty
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
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