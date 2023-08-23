package main

import (
	"bytes"
	"encoding/json"
	"errors"
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
	s.Router.Post("/", s.broker)
	s.Router.Post("/authentication", s.authentication)
}

func (s *Server) broker(w http.ResponseWriter, r *http.Request) {

	resPayload := JSONResponse{
		Error:   false,
		Message: "Successfully hit the Broker!",
	}

	err := s.writeJSON(w, resPayload, http.StatusOK)
	if err != nil {
		s.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	log.Println("broker service: successful broker service call")
}

func (s *Server) authentication(w http.ResponseWriter, r *http.Request) {
	var authPayload AuthenticationPayload

	err := s.readJSON(w, r, &authPayload)
	if authPayload.Email == "" || authPayload.Password == "" {
		s.errorJSON(w, errors.New("email and password must be non-empty"), http.StatusBadRequest)
		return
	}
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	reqPayload, err := json.Marshal(authPayload)
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	req, err := http.NewRequest("POST", "http://authentication/authentication", bytes.NewBuffer(reqPayload))
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	defer res.Body.Close()

	var resJSON JSONResponse

	err = json.NewDecoder(res.Body).Decode(&resJSON)
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = s.writeJSON(w, resJSON, http.StatusOK)
	if err != nil {
		s.errorJSON(w, err, http.StatusInternalServerError)
		return
	}
}
