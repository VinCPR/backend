-- name: CreateStudentToGroup :one
INSERT INTO "student_to_group" (
    academic_year_id,
    student_id,
    group_id
) VALUES (
    $1 , $2, $3
) RETURNING *;

-- name: GetStudentToGroupByAcademicYearID :many
SELECT * FROM "student_to_group"
WHERE academic_year_id = $1;

-- name: GetStudentToGroupByStudentID :many
SELECT * FROM "student_to_group"
WHERE student_id = $1;

-- name: GetStudentToGroupByGroupID :many
SELECT * FROM "student_to_group"
WHERE group_id = $1;

-- name: GetStudentToGroupByAcademicYearIDAndStudentID :many
SELECT * FROM "student_to_group"
WHERE academic_year_id = $1 AND student_id = $2;