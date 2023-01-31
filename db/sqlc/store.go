package db

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	Db *pgxpool.Pool
	*Queries
}

func NewStore(db *pgxpool.Pool) *Store {
	return &Store{
		Db:      db,
		Queries: New(db),
	}
}
