package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

func main() {
	srv := NewServer()
	fmt.Println("running server...")
	err := http.ListenAndServe(":80", srv)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("server closed")
	} else if err != nil {
		fmt.Println("error starting server:", err)
		os.Exit(1)
	}
}
