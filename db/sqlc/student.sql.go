// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: student.sql

package db

import (
	"context"
)

const createStudent = `-- name: CreateStudent :one
INSERT INTO "student" (
    user_id,
    student_id,
    first_name ,
    last_name ,
    mobile
) VALUES (
    $1 , $2 , $3, $4 , $5
) RETURNING id, user_id, student_id, first_name, last_name, mobile, created_at
`

type CreateStudentParams struct {
	UserID    int64  `json:"user_id"`
	StudentID string `json:"student_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Mobile    string `json:"mobile"`
}

func (q *Queries) CreateStudent(ctx context.Context, arg CreateStudentParams) (Student, error) {
	row := q.db.QueryRow(ctx, createStudent,
		arg.UserID,
		arg.StudentID,
		arg.FirstName,
		arg.LastName,
		arg.Mobile,
	)
	var i Student
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.StudentID,
		&i.FirstName,
		&i.LastName,
		&i.Mobile,
		&i.CreatedAt,
	)
	return i, err
}

const getStudentByStudentId = `-- name: GetStudentByStudentId :one
SELECT id, user_id, student_id, first_name, last_name, mobile, created_at FROM "student"
WHERE student_id = $1 LIMIT 1
`

func (q *Queries) GetStudentByStudentId(ctx context.Context, studentID string) (Student, error) {
	row := q.db.QueryRow(ctx, getStudentByStudentId, studentID)
	var i Student
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.StudentID,
		&i.FirstName,
		&i.LastName,
		&i.Mobile,
		&i.CreatedAt,
	)
	return i, err
}

const getStudentByUserId = `-- name: GetStudentByUserId :one
SELECT id, user_id, student_id, first_name, last_name, mobile, created_at FROM "student"
WHERE user_id = $1 LIMIT 1
`

func (q *Queries) GetStudentByUserId(ctx context.Context, userID int64) (Student, error) {
	row := q.db.QueryRow(ctx, getStudentByUserId, userID)
	var i Student
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.StudentID,
		&i.FirstName,
		&i.LastName,
		&i.Mobile,
		&i.CreatedAt,
	)
	return i, err
}

const listStudentsByName = `-- name: ListStudentsByName :many
SELECT id, user_id, student_id, first_name, last_name, mobile, created_at FROM "student"
ORDER BY first_name, last_name
LIMIT $1
OFFSET $2
`

type ListStudentsByNameParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListStudentsByName(ctx context.Context, arg ListStudentsByNameParams) ([]Student, error) {
	rows, err := q.db.Query(ctx, listStudentsByName, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Student{}
	for rows.Next() {
		var i Student
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.StudentID,
			&i.FirstName,
			&i.LastName,
			&i.Mobile,
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

const listStudentsByStudentID = `-- name: ListStudentsByStudentID :many
SELECT id, user_id, student_id, first_name, last_name, mobile, created_at FROM "student"
ORDER BY student_id
LIMIT $1
OFFSET $2
`

type ListStudentsByStudentIDParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListStudentsByStudentID(ctx context.Context, arg ListStudentsByStudentIDParams) ([]Student, error) {
	rows, err := q.db.Query(ctx, listStudentsByStudentID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Student{}
	for rows.Next() {
		var i Student
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.StudentID,
			&i.FirstName,
			&i.LastName,
			&i.Mobile,
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