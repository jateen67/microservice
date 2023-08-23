package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jateen67/authentication/db"
)

const port = "80"

func main() {
	// start postgres
	database, err := db.ConnectToDB()
	if err != nil {
		log.Fatalf("could not connect to postgres: %s", err)
	}
	defer database.Close()

	log.Println("connected to postgres successfully")

	err = db.CreateTable(database)
	if err != nil {
		log.Fatalf("could not create users table: %v", err)
	}

	log.Println("users table created successfully")

	// start auth server
	srv := NewServer()
	log.Println("starting authentication server...")
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), srv)

	if errors.Is(err, http.ErrServerClosed) {
		log.Println("server closed")
	} else if err != nil {
		log.Println("error starting server: ", err)
		os.Exit(1)
	}

}
