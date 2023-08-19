package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

const port = "80"

func main() {
	srv := NewServer()
	fmt.Println("running server...")
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), srv)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("server closed")
	} else if err != nil {
		fmt.Println("error starting server:", err)
		os.Exit(1)
	}
}
