package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"

	"github.com/VinCPR/backend/util"
)

var testQueries *Queries
var testDB *pgx.Conn

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	testDB, err = pgx.Connect(context.Background(), config.DBUrl)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
