// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: academic_calendar_event.sql

package db

import (
	"context"
	"time"
)

const createEvent = `-- name: CreateEvent :one
INSERT INTO "academic_calendar_event" (
    "academic_year_id", "name", "type", "start_date", "end_date"
) VALUES(
    $1, $2, $3, $4, $5
) RETURNING id, academic_year_id, name, type, start_date, end_date, created_at
`

type CreateEventParams struct {
	AcademicYearID int64     `json:"academic_year_id"`
	Name           string    `json:"name"`
	Type           string    `json:"type"`
	StartDate      time.Time `json:"start_date"`
	EndDate        time.Time `json:"end_date"`
}

func (q *Queries) CreateEvent(ctx context.Context, arg CreateEventParams) (AcademicCalendarEvent, error) {
	row := q.db.QueryRow(ctx, createEvent,
		arg.AcademicYearID,
		arg.Name,
		arg.Type,
		arg.StartDate,
		arg.EndDate,
	)
	var i AcademicCalendarEvent
	err := row.Scan(
		&i.ID,
		&i.AcademicYearID,
		&i.Name,
		&i.Type,
		&i.StartDate,
		&i.EndDate,
		&i.CreatedAt,
	)
	return i, err
}

type CreateEventsParams struct {
	AcademicYearID int64     `json:"academic_year_id"`
	Name           string    `json:"name"`
	Type           string    `json:"type"`
	StartDate      time.Time `json:"start_date"`
	EndDate        time.Time `json:"end_date"`
}

const listEventsByAcademicYearID = `-- name: ListEventsByAcademicYearID :many
SELECT id, academic_year_id, name, type, start_date, end_date, created_at FROM "academic_calendar_event"
WHERE "academic_year_id" = $1
ORDER BY "start_date"
LIMIT $2
OFFSET $3
`

type ListEventsByAcademicYearIDParams struct {
	AcademicYearID int64 `json:"academic_year_id"`
	Limit          int32 `json:"limit"`
	Offset         int32 `json:"offset"`
}

func (q *Queries) ListEventsByAcademicYearID(ctx context.Context, arg ListEventsByAcademicYearIDParams) ([]AcademicCalendarEvent, error) {
	rows, err := q.db.Query(ctx, listEventsByAcademicYearID, arg.AcademicYearID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AcademicCalendarEvent{}
	for rows.Next() {
		var i AcademicCalendarEvent
		if err := rows.Scan(
			&i.ID,
			&i.AcademicYearID,
			&i.Name,
			&i.Type,
			&i.StartDate,
			&i.EndDate,
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
