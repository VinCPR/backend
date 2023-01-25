package script

import (
	"context"
	"os"
	"testing"

	"github.com/jackc/pgx/v4"
	"github.com/rs/zerolog/log"

	"github.com/VinCPR/backend/api"
	db "github.com/VinCPR/backend/db/sqlc"
	"github.com/VinCPR/backend/util"
)

var server *api.Server
var config util.Config

func TestMain(m *testing.M) {
	var err error
	config, err = util.LoadConfig("..")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}
	conn, err := pgx.Connect(context.Background(), config.DBUrl)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}

	store := db.NewStore(conn)

	server, err = api.NewServer(config, store)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server")
	}
	os.Exit(m.Run())
}
