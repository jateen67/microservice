package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	username = "postgres"
	password = "password"
	database = "users_db"
)

type Database struct {
	Conn *sql.DB
}

func InitPostgresDb() (Database, error) {
	db := Database{}
	pgInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, username, password, database)

	conn, err := sql.Open("postgres", pgInfo)
	if err != nil {
		return db, err
	}

	db.Conn = conn
	err = db.Conn.Ping()
	if err != nil {
		return db, err
	}

	log.Println("connected to postgres successfully")
	return db, nil
}
