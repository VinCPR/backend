package main

import (
	"context"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	db "github.com/VinCPR/backend/db/sqlc"
	"github.com/VinCPR/backend/docs"
	"github.com/VinCPR/backend/util"

	"github.com/VinCPR/backend/api"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}

	docs.SwaggerInfo.BasePath = config.BasePath
	docs.SwaggerInfo.Host = config.Host

	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	runDBMigration(config.MigrationURL, config.DBUrl)

	conn, err := pgxpool.Connect(context.Background(), config.DBUrl)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}
	defer conn.Close()

	exitOnDisconnectDB(conn)

	store := db.NewStore(conn)

	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server")
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start server")
	}
}

func runDBMigration(migrationURL string, dbUrl string) {
	migration, err := migrate.New(migrationURL, dbUrl)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create new migrate instance")
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal().Err(err).Msg("failed to run migrate up")
	}

	log.Info().Msg("db migrated successfully")
}

func exitOnDisconnectDB(conn *pgxpool.Pool) {
	c := cron.New()
	c.AddFunc("@every 5s", func() {
		if err := conn.Ping(context.Background()); err != nil {
			log.Fatal().Err(err).Msg("connection to db failed, restarting api")
			os.Exit(1)
		}
	})
	c.Start()
}
