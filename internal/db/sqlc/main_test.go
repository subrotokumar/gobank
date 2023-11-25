package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dbDriver = "postgres"
	dbUrl    = "postgres://root:mysecretpassword@localhost:5432/simple_bank?sslmode=disable"
)

var testQuery *Queries

func TestMain(m *testing.M) {
	ctx := context.Background()
	conn, err := pgxpool.New(ctx, dbUrl)
	if err != nil {
		log.Fatal("connect to database", err)
	}
	testQuery = New(conn)
	os.Exit(m.Run())
}
