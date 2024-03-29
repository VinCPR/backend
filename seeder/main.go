package main

import (
	"context"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
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

	conn, err := pgxpool.Connect(context.Background(), config.DBUrl)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}
	defer conn.Close()

	store := db.NewStore(conn)

	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server")
	}

	script.SeedAdminData(config.BasePath, server)
	script.SeedStudentData(config.BasePath, server)
	script.SeedAttendingData(config.BasePath, server)
	script.SeedHospitalData(config.BasePath, server)
	script.SeedSpecialtyData(config.BasePath, server)
	script.SeedServiceData(config.BasePath, server)
	script.SeedServiceToAttendingData(config.BasePath, server)
	script.SeedAcademicYearData(config.BasePath, server)
	script.SeedGroupData(config.BasePath, server)
	script.SeedStudentToGroupData(config.BasePath, server)
}
