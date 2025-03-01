package main

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"gitlab.com/alexandrudeac/minibank/api"
	db "gitlab.com/alexandrudeac/minibank/db/sqlc"
	"log"
)

const (
	dbDriver      = "pgx"
	dbSource      = "postgres://root:secret@localhost:5432/minibank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)

	server := api.NewServer(store)
	err = server.Run(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
