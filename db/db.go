package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

type ConnConfig struct {
	User     string
	Password string
	Host     string
	Port     int16
	DBName   string
}

func createConnString() string {
	config := ConnConfig{
		User:     "postgres",
		Password: "postgres",
		Host:     "172.17.69.80",
		Port:     5432,
		DBName:   "gorm",
	}

	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DBName,
	)
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
