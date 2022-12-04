package db

import (
	"github.com/jackc/pgx/v4"
)

type Store struct {
	db *pgx.Conn
	*Queries
}

func NewStore(db *pgx.Conn) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}
