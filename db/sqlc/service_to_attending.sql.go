// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: service_to_attending.sql

package db

import (
	"context"
)

const createServiceToAttending = `-- name: CreateServiceToAttending :one
INSERT INTO "service_to_attending" (
    service_id,
    attending_id
) VALUES (
    $1 , $2 
) RETURNING id, service_id, attending_id, created_at
`

type CreateServiceToAttendingParams struct {
	ServiceID   int64 `json:"service_id"`
	AttendingID int64 `json:"attending_id"`
}

func (q *Queries) CreateServiceToAttending(ctx context.Context, arg CreateServiceToAttendingParams) (ServiceToAttending, error) {
	row := q.db.QueryRow(ctx, createServiceToAttending, arg.ServiceID, arg.AttendingID)
	var i ServiceToAttending
	err := row.Scan(
		&i.ID,
		&i.ServiceID,
		&i.AttendingID,
		&i.CreatedAt,
	)
	return i, err
}

const getServiceToAttendingByAttendingID = `-- name: GetServiceToAttendingByAttendingID :one
SELECT id, service_id, attending_id, created_at FROM "service_to_attending"
WHERE attending_id = $1 LIMIT 1
`

func (q *Queries) GetServiceToAttendingByAttendingID(ctx context.Context, attendingID int64) (ServiceToAttending, error) {
	row := q.db.QueryRow(ctx, getServiceToAttendingByAttendingID, attendingID)
	var i ServiceToAttending
	err := row.Scan(
		&i.ID,
		&i.ServiceID,
		&i.AttendingID,
		&i.CreatedAt,
	)
	return i, err
}

const getServiceToAttendingByServiceID = `-- name: GetServiceToAttendingByServiceID :one
SELECT id, service_id, attending_id, created_at FROM "service_to_attending"
WHERE service_id = $1 LIMIT 1
`

// TODO: change list by service id -> many
// attending_id -> many
func (q *Queries) GetServiceToAttendingByServiceID(ctx context.Context, serviceID int64) (ServiceToAttending, error) {
	row := q.db.QueryRow(ctx, getServiceToAttendingByServiceID, serviceID)
	var i ServiceToAttending
	err := row.Scan(
		&i.ID,
		&i.ServiceID,
		&i.AttendingID,
		&i.CreatedAt,
	)
	return i, err
}

const listServicesToAttendingsByAttendingID = `-- name: ListServicesToAttendingsByAttendingID :many
SELECT id, service_id, attending_id, created_at FROM "service_to_attending"
ORDER BY "attending_id"
`

func (q *Queries) ListServicesToAttendingsByAttendingID(ctx context.Context) ([]ServiceToAttending, error) {
	rows, err := q.db.Query(ctx, listServicesToAttendingsByAttendingID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ServiceToAttending{}
	for rows.Next() {
		var i ServiceToAttending
		if err := rows.Scan(
			&i.ID,
			&i.ServiceID,
			&i.AttendingID,
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

const listServicesToAttendingsByServiceID = `-- name: ListServicesToAttendingsByServiceID :many
SELECT id, service_id, attending_id, created_at FROM "service_to_attending"
ORDER BY "service_id"
`

func (q *Queries) ListServicesToAttendingsByServiceID(ctx context.Context) ([]ServiceToAttending, error) {
	rows, err := q.db.Query(ctx, listServicesToAttendingsByServiceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ServiceToAttending{}
	for rows.Next() {
		var i ServiceToAttending
		if err := rows.Scan(
			&i.ID,
			&i.ServiceID,
			&i.AttendingID,
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

const listServicesToAttendingsByServiceIDAndAttendingID = `-- name: ListServicesToAttendingsByServiceIDAndAttendingID :many
SELECT id, service_id, attending_id, created_at FROM "service_to_attending"
ORDER BY "service_id","attending_id"
`

func (q *Queries) ListServicesToAttendingsByServiceIDAndAttendingID(ctx context.Context) ([]ServiceToAttending, error) {
	rows, err := q.db.Query(ctx, listServicesToAttendingsByServiceIDAndAttendingID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ServiceToAttending{}
	for rows.Next() {
		var i ServiceToAttending
		if err := rows.Scan(
			&i.ID,
			&i.ServiceID,
			&i.AttendingID,
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
