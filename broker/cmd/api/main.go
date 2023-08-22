package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
)

const port = "80"

func main() {
	// start broker server
	srv := NewServer()
	log.Println("starting broker server...")
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), srv)

	if errors.Is(err, http.ErrServerClosed) {
		log.Println("server closed")
	} else if err != nil {
		log.Println("error starting server:", err)
		os.Exit(1)
	}

	log.Println("broker server started")

}
