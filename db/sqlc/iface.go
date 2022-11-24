package db

import (
	"context"

	"github.com/google/uuid"
)

type IStoreInstruction interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	GetUser(ctx context.Context, username string) (User, error)
	ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error)

	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	GetSession(ctx context.Context, id uuid.UUID) (Session, error)
}
