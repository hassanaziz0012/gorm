package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func createConnString() string {
	user := "postgres"
	password := "postgres"
	host := "172.17.69.80"
	port := 5432
	dbname := "gorm"

	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s", user, password, host, port, dbname)
}

func Connect() *pgxpool.Pool {
	connString := createConnString()
	dbpool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		log.Fatal(err)
	}
	DB = dbpool
	return DB
}
