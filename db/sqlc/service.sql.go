// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: service.sql

package db

import (
	"context"
)

const createService = `-- name: CreateService :one
INSERT INTO "service" (
   specialty_id,
   hospital_id,
   name,
   description
) VALUES (
   $1 , $2 , $3, $4
) RETURNING id, specialty_id, hospital_id, name, description, created_at
`

type CreateServiceParams struct {
	SpecialtyID int64  `json:"specialty_id"`
	HospitalID  int64  `json:"hospital_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (q *Queries) CreateService(ctx context.Context, arg CreateServiceParams) (Service, error) {
	row := q.db.QueryRow(ctx, createService,
		arg.SpecialtyID,
		arg.HospitalID,
		arg.Name,
		arg.Description,
	)
	var i Service
	err := row.Scan(
		&i.ID,
		&i.SpecialtyID,
		&i.HospitalID,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
	)
	return i, err
}

const getServiceByID = `-- name: GetServiceByID :one
SELECT id, specialty_id, hospital_id, name, description, created_at FROM "service"
WHERE "id" = $1 LIMIT 1
`

func (q *Queries) GetServiceByID(ctx context.Context, id int64) (Service, error) {
	row := q.db.QueryRow(ctx, getServiceByID, id)
	var i Service
	err := row.Scan(
		&i.ID,
		&i.SpecialtyID,
		&i.HospitalID,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
	)
	return i, err
}

const getServiceByIndex = `-- name: GetServiceByIndex :one
 SELECT id, specialty_id, hospital_id, name, description, created_at FROM "service"
 WHERE specialty_id = $1 AND hospital_id = $2 AND name = $3 LIMIT 1
`

type GetServiceByIndexParams struct {
	SpecialtyID int64  `json:"specialty_id"`
	HospitalID  int64  `json:"hospital_id"`
	Name        string `json:"name"`
}

func (q *Queries) GetServiceByIndex(ctx context.Context, arg GetServiceByIndexParams) (Service, error) {
	row := q.db.QueryRow(ctx, getServiceByIndex, arg.SpecialtyID, arg.HospitalID, arg.Name)
	var i Service
	err := row.Scan(
		&i.ID,
		&i.SpecialtyID,
		&i.HospitalID,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
	)
	return i, err
}

const getServiceByName = `-- name: GetServiceByName :one
SELECT id, specialty_id, hospital_id, name, description, created_at FROM "service"
WHERE name = $1 LIMIT 1
`

func (q *Queries) GetServiceByName(ctx context.Context, name string) (Service, error) {
	row := q.db.QueryRow(ctx, getServiceByName, name)
	var i Service
	err := row.Scan(
		&i.ID,
		&i.SpecialtyID,
		&i.HospitalID,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
	)
	return i, err
}

const listServicesByHospitalID = `-- name: ListServicesByHospitalID :many
SELECT id, specialty_id, hospital_id, name, description, created_at FROM "service"
ORDER BY "hospital_id"
LIMIT $1
OFFSET $2
`

type ListServicesByHospitalIDParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListServicesByHospitalID(ctx context.Context, arg ListServicesByHospitalIDParams) ([]Service, error) {
	rows, err := q.db.Query(ctx, listServicesByHospitalID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Service{}
	for rows.Next() {
		var i Service
		if err := rows.Scan(
			&i.ID,
			&i.SpecialtyID,
			&i.HospitalID,
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

const listServicesBySpecialtyID = `-- name: ListServicesBySpecialtyID :many
SELECT id, specialty_id, hospital_id, name, description, created_at FROM "service"
ORDER BY "specialty_id"
LIMIT $1
OFFSET $2
`

type ListServicesBySpecialtyIDParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListServicesBySpecialtyID(ctx context.Context, arg ListServicesBySpecialtyIDParams) ([]Service, error) {
	rows, err := q.db.Query(ctx, listServicesBySpecialtyID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Service{}
	for rows.Next() {
		var i Service
		if err := rows.Scan(
			&i.ID,
			&i.SpecialtyID,
			&i.HospitalID,
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

const listServicesBySpecialtyIDAndHospitalID = `-- name: ListServicesBySpecialtyIDAndHospitalID :many
SELECT id, specialty_id, hospital_id, name, description, created_at FROM "service"
ORDER BY "specialty_id","hospital_id"
LIMIT $1
OFFSET $2
`

type ListServicesBySpecialtyIDAndHospitalIDParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListServicesBySpecialtyIDAndHospitalID(ctx context.Context, arg ListServicesBySpecialtyIDAndHospitalIDParams) ([]Service, error) {
	rows, err := q.db.Query(ctx, listServicesBySpecialtyIDAndHospitalID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Service{}
	for rows.Next() {
		var i Service
		if err := rows.Scan(
			&i.ID,
			&i.SpecialtyID,
			&i.HospitalID,
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
