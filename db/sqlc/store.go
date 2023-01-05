package db

import (
	"github.com/jackc/pgx/v4"
)

type Store struct {
	Db *pgx.Conn
	*Queries
}

func NewStore(db *pgx.Conn) *Store {
	return &Store{
		Db:      db,
		Queries: New(db),
	}
}
