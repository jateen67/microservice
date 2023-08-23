package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type Server struct {
	chi.Router
}

func NewServer() *Server {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
	}))

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

	payload := JSONResponse{
		Error:   false,
		Message: "Successfully logged in!",
		Data:    "Replace with User data",
	}

	err := s.writeJSON(w, payload)

	if err != nil {
		log.Println("error:", err)
		return
	}

	log.Println("successful authentication service login")
}
