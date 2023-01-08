-- name: CreateStudent :one
INSERT INTO "student" (
    user_id,
    student_id,
    first_name ,
    last_name ,
    mobile
) VALUES (
    $1 , $2 , $3, $4 , $5
) RETURNING *;

-- name: GetStudentByID :one
SELECT * FROM "student"
WHERE id = $1 LIMIT 1;

-- name: GetStudentByUserId :one
SELECT * FROM "student"
WHERE user_id = $1 LIMIT 1;

-- name: GetStudentByStudentId :one
SELECT * FROM "student"
WHERE student_id = $1 LIMIT 1;

-- name: ListStudentsByName :many
SELECT * FROM "student"
ORDER BY first_name, last_name
LIMIT $1
OFFSET $2;

-- name: ListStudentsByStudentID :many
SELECT * FROM "student"
ORDER BY student_id
LIMIT $1
OFFSET $2;