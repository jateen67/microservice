package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jateen67/authentication/db"
)

const port = "80"

func main() {
	// start postgres
	database, err := db.ConnectToDB()
	if err != nil {
		log.Fatalf("could not connect to postgres: %v", err)
	}
	defer database.Close()

	log.Println("connected to postgres successfully")

	err = db.CreateTable(database)
	if err != nil {
		log.Fatalf("could not create users table: %v", err)
	}

	log.Println("users table created successfully")

	// check if user exists already
	userExists, err := db.UserExists(database, "admin@example.com")
	if err != nil {
		log.Fatalf("error checking if user exists: %v", err)
	}

	if !userExists {
		// insert a new user
		err = db.InsertUser(database, "admin@example.com", "password123", "John", "Doe", time.Now())
		if err != nil {
			log.Fatalf("error inserting user: %v", err)
		}
		log.Println("user inserted successfully")
	} else {
		log.Println("user already inserted")
	}

	userDB := db.NewUserDBImpl(database)
	// start auth server
	srv := newServer(userDB).Router
	log.Println("starting authentication server...")
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), srv)

	if errors.Is(err, http.ErrServerClosed) {
		log.Println("auth server closed")
	} else if err != nil {
		log.Println("error starting auth server:", err)
		os.Exit(1)
	}
}
