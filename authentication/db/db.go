package db

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func ConnectToDB() (*sql.DB, error) {
	connString := "host=postgres port=5432 user=postgres password=password dbname=users_db sslmode=disable"
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
				log.Println("connected to postgres successfully")
				return db, nil
			}
		}

		if count > 20 {
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
