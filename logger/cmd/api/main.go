package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jateen67/logger/client"
)

const port = "80"

func main() {
	// start mongo
	mongoClient, err := client.ConnectToClient()
	if err != nil {
		log.Fatalf("could not connect to mongo: %s", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// close connection
	defer func() {
		err = mongoClient.Disconnect(ctx)
		if err != nil {
			panic(err)
		}
	}()

	logEntryClient := client.NewLogEntryClientImpl(mongoClient)
	// start logger server
	srv := NewServer(logEntryClient).Router
	log.Println("starting logger server")
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), srv)
	if errors.Is(err, http.ErrServerClosed) {
		log.Println("logger server closed")
	} else if err != nil {
		log.Println("error starting logger server: ", err)
		os.Exit(1)
	}
}
