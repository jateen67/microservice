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
	srv := newServer()
	log.Println("starting broker server...")
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), srv)

	if errors.Is(err, http.ErrServerClosed) {
		log.Println("broker server closed")
	} else if err != nil {
		log.Println("error starting broker server:", err)
		os.Exit(1)
	}

}
