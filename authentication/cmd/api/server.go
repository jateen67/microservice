package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/jateen67/authentication/db"
)

type Server struct {
	Router chi.Router
	UserDB db.UserDB
}

type AuthenticationPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewServer(userDB db.UserDB) *Server {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
	}))

	s := &Server{
		Router: r,
		UserDB: userDB,
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

	user, err := s.UserDB.GetUserByEmail(reqPayload.Email)
	if err != nil {
		s.errorJSON(w, errors.New("couldn't find user in database"), http.StatusBadRequest)
		return
	}

	resPayload := JSONResponse{
		Error:   false,
		Message: fmt.Sprintf("Successfully signed in as %s!", user.Email),
		Data:    user,
	}

	err = s.writeJSON(w, resPayload, http.StatusOK)
	if err != nil {
		s.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	log.Println("authentication service: successful signin")
}
