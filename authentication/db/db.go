package db

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

type Models struct {
	User User
}

func ConnectToDB() (*sql.DB, error) {
	connString := os.Getenv("DATABASE_CONNECTION_STRING")
	count := 1

	for {
		db, err := sql.Open("postgres", connString)
		if err != nil {
			log.Println("could not connect to postgres. retrying... ")
			count++
		} else {
			err = db.Ping()
			if err != nil {
				log.Println("postgres connection test failed. retrying...")
				count++
				db.Close()
			} else {
				return db, nil
			}
		}

		if count > 10 {
			return nil, err
		}

		log.Println("retrying in 1 second...")
		time.Sleep(1 * time.Second)
	}
}

func CreateTable(db *sql.DB) error {
	query := `
        CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            email TEXT NOT NULL,
            password TEXT NOT NULL
        )
    `
	_, err := db.Exec(query)
	return err
}

func UserExists(db *sql.DB, email string) (bool, error) {
	query := "SELECT COUNT(*) FROM users WHERE email= $1"
	var count int
	err := db.QueryRow(query, email).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func InsertUser(db *sql.DB, email, password string) error {
	query := "INSERT INTO users (email, password) VALUES ($1, $2)"
	_, err := db.Exec(query, email, password)
	return err
}
