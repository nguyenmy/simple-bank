package main

import (
	"database/sql"
	server "go-simple-bank/api"
	db "go-simple-bank/db/sqlc"
	"go-simple-bank/db/util"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("Failed to load config: %v\n", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalf("Failed to connect database: %v\n", err)
	}

	store := db.NewStore(conn)
	server := server.NewServer(store)
	server.Start(config.ServerAddress)
}
