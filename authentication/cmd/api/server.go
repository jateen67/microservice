package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type Server struct {
	chi.Router
}

type AuthenticationPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
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

	var reqPayload AuthenticationPayload

	err := s.readJSON(w, r, &reqPayload)
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	resPayload := JSONResponse{
		Error:   false,
		Message: fmt.Sprintf("Successfully logged in as %s!", reqPayload.Email),
		Data:    "Replace with User data",
	}

	err = s.writeJSON(w, resPayload, http.StatusOK)

	if err != nil {
		s.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	log.Println("authentication service: succesful login")
}
