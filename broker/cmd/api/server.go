package main

import (
	"bytes"
	"encoding/json"
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
	s.Router.Post("/", s.broker)
	s.Router.Post("/authentication", s.authentication)
}

func (s *Server) broker(w http.ResponseWriter, r *http.Request) {

	responsePayload := JSONResponse{
		Error:   false,
		Message: "Successfully hit the Broker!",
	}

	err := s.writeJSON(w, responsePayload, http.StatusOK)
	if err != nil {
		s.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	log.Println("successful broker service call from broker")
}

func (s *Server) authentication(w http.ResponseWriter, r *http.Request) {

	jsonData, _ := json.MarshalIndent(r.Body, "", "\t")

	req, err := http.NewRequest("POST", "http://authentication/authentication", bytes.NewBuffer(jsonData))
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

	var jsonFromService JSONResponse

	err = json.NewDecoder(res.Body).Decode(&jsonFromService)
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = s.writeJSON(w, jsonFromService, http.StatusOK)
	if err != nil {
		s.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	log.Println("successful authentication service call from broker")
}
