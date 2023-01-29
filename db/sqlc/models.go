// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"time"

	"github.com/google/uuid"
)

type AcademicCalendarEvent struct {
	ID             int64     `json:"id"`
	AcademicYearID int64     `json:"academic_year_id"`
	Name           string    `json:"name"`
	Type           string    `json:"type"`
	StartDate      time.Time `json:"start_date"`
	EndDate        time.Time `json:"end_date"`
	CreatedAt      time.Time `json:"created_at"`
}

type AcademicYear struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	CreatedAt time.Time `json:"created_at"`
}

type Attending struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"user_id"`
	AttendingID string    `json:"attending_id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Mobile      string    `json:"mobile"`
	Biography   string    `json:"biography"`
	Image       string    `json:"image"`
	CreatedAt   time.Time `json:"created_at"`
}

type Block struct {
	ID             int64     `json:"id"`
	AcademicYearID int64     `json:"academic_year_id"`
	Name           string    `json:"name"`
	Period         int64     `json:"period"`
	CreatedAt      time.Time `json:"created_at"`
}

type ClinicalRotationEvent struct {
	ID             int64     `json:"id"`
	AcademicYearID int64     `json:"academic_year_id"`
	GroupID        int64     `json:"group_id"`
	ServiceID      int64     `json:"service_id"`
	StartDate      time.Time `json:"start_date"`
	EndDate        time.Time `json:"end_date"`
	CreatedAt      time.Time `json:"created_at"`
}

type Group struct {
	ID             int64     `json:"id"`
	AcademicYearID int64     `json:"academic_year_id"`
	Name           string    `json:"name"`
	CreatedAt      time.Time `json:"created_at"`
}

type GroupToBlock struct {
	ID             int64     `json:"id"`
	AcademicYearID int64     `json:"academic_year_id"`
	GroupID        int64     `json:"group_id"`
	BlockID        int64     `json:"block_id"`
	CreatedAt      time.Time `json:"created_at"`
}

type Hospital struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Address     string    `json:"address"`
	CreatedAt   time.Time `json:"created_at"`
}

type Period struct {
	ID             int64     `json:"id"`
	AcademicYearID int64     `json:"academic_year_id"`
	Name           string    `json:"name"`
	StartDate      time.Time `json:"start_date"`
	EndDate        time.Time `json:"end_date"`
	CreatedAt      time.Time `json:"created_at"`
}

type Service struct {
	ID          int64     `json:"id"`
	SpecialtyID int64     `json:"specialty_id"`
	HospitalID  int64     `json:"hospital_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type ServiceToAttending struct {
	ID          int64     `json:"id"`
	ServiceID   int64     `json:"service_id"`
	AttendingID int64     `json:"attending_id"`
	CreatedAt   time.Time `json:"created_at"`
}

type Session struct {
	ID           uuid.UUID `json:"id"`
	UserEmail    string    `json:"user_email"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type Specialty struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type Student struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	StudentID string    `json:"student_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Mobile    string    `json:"mobile"`
	Biography string    `json:"biography"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
}

type StudentToGroup struct {
	ID             int64     `json:"id"`
	AcademicYearID int64     `json:"academic_year_id"`
	StudentID      int64     `json:"student_id"`
	GroupID        int64     `json:"group_id"`
	CreatedAt      time.Time `json:"created_at"`
}

type User struct {
	ID             int64     `json:"id"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"hashed_password"`
	RoleName       string    `json:"role_name"`
	CreatedAt      time.Time `json:"created_at"`
}
