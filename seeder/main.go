package main

import (
	"context"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/VinCPR/backend/api"
	db "github.com/VinCPR/backend/db/sqlc"
	"github.com/VinCPR/backend/seeder/script"
	"github.com/VinCPR/backend/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}

	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	// clear data
	script.ClearDataDBMigration(config.MigrationURL, config.DBUrl)

	conn, err := pgx.Connect(context.Background(), config.DBUrl)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}

	store := db.NewStore(conn)

	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server")
	}

	script.SeedStudentData(config.BasePath, server)
	script.SeedAttendingData(config.BasePath, server)
}
