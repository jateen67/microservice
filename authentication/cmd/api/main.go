package main

import (
	"log"

	"github.com/jateen67/authentication/db"
)

func main() {
	database, err := db.ConnectToDB()
	if err != nil {
		log.Fatalf("could not connect to postgres: %s", err)
	}
	defer database.Close()

	err = db.CreateTable(database)
	if err != nil {
		log.Fatalf("could not create users table: %v", err)
	}

	log.Println("users table created successfully")
}
