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
	s.Router.Post("/", s.broker)
}

func (s *Server) broker(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)

	payload := JSONResponse{
		Error:   false,
		Message: "Successfully hit the Broker!",
	}

	err := s.writeJSON(w, payload)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	log.Println("successful broker service call")
}
