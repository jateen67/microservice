package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	chi.Router
}

func NewServer() *Server {
	r := chi.NewRouter()

	s := &Server{
		Router: r,
	}
	s.routes()

	return s
}

func (s *Server) routes() {
	s.Router.Post("/authentication", s.authentication)
}

func (s *Server) authentication(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)

	payload := JSONResponse{
		Error:   false,
		Message: "Successfully logged in!",
	}

	err := s.writeJSON(w, payload)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	log.Println("successful authentication service login")
}
