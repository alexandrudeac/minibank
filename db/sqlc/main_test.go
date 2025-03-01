package db

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
	"os"
	"testing"
)

const (
	dbDriver = "pgx"
	dbSource = "postgres://root:secret@localhost:5432/minibank?sslmode=disable"
)

var testStore Store
var db *sql.DB

func TestMain(m *testing.M) {
	openedConn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	testStore = NewStore(openedConn)
	db = openedConn
	os.Exit(m.Run())
}
