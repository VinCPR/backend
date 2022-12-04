package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"

	db "github.com/VinCPR/backend/db/sqlc"
	"github.com/VinCPR/backend/util"

	"github.com/VinCPR/backend/api"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot read config file", err)
	}

	conn, err := pgx.Connect(context.Background(), config.DBUrl)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("Cannot create server", err)
	}

	if err := server.Start(config.HTTPServerAddress); err != nil {
		log.Fatal("cannot start server", err)
	}
}
