package db

import (
	"database/sql"
	"go-simple-bank/db/util"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../.")
	if err != nil {
		log.Fatalf("Failed to load config: %v\n", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalf("Failed to connect database: %v\n", err)
	}

	testQueries = New(testDB)
	os.Exit(m.Run())
}
