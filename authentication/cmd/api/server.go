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

type server struct {
	Router chi.Router
	UserDB db.UserDB
}

type authenticationPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func newServer(userDB db.UserDB) *server {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
	}))

	s := &server{
		Router: r,
		UserDB: userDB,
	}
	s.routes()

	return s
}

func (s *server) routes() {
	s.Router.Post("/authentication", s.authentication)
}

func (s *server) authentication(w http.ResponseWriter, r *http.Request) {
	var reqPayload authenticationPayload

	err := s.readJSON(w, r, &reqPayload)
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	user, err := s.UserDB.GetUserByEmail(reqPayload.Email)
	if err != nil {
		s.errorJSON(w, errors.New("couldn't find user in database"), http.StatusNotFound)
		return
	}

	err = s.UserDB.PasswordCheck(user.Password, reqPayload.Password)
	if err != nil {
		s.errorJSON(w, errors.New("invalid credentials"), http.StatusUnauthorized)
		return
	}

	resPayload := jsonResponse{
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
