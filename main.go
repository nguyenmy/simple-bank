package main

import (
	"database/sql"
	server "go-simple-bank/api"
	db "go-simple-bank/db/sqlc"
	"log"

	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5400/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8092"
)

func main() {
	var err error
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatalf("Failed to connect database: %v\n", err)
	}

	store := db.NewStore(conn)
	server := server.NewServer(store)
	server.Start(serverAddress)
}
