package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/jateen67/logger/client"
)

type Server struct {
	Router         chi.Router
	LogEntryClient client.LogEntryClient
}

type LoggerPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func NewServer(logEntryClient client.LogEntryClient) *Server {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
	}))

	s := &Server{
		Router:         r,
		LogEntryClient: logEntryClient,
	}
	s.routes()

	return s
}

func (s *Server) routes() {
	s.Router.Post("/logger", s.logger)
}

func (s *Server) logger(w http.ResponseWriter, r *http.Request) {

	var reqPayload LoggerPayload

	err := s.readJSON(w, r, &reqPayload)
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	doc := client.LogEntry{
		Name: reqPayload.Name,
		Data: reqPayload.Data,
	}

	err = s.LogEntryClient.InsertLogEntry(doc)
	if err != nil {
		s.errorJSON(w, errors.New("couldn't add new log to collection"), http.StatusBadRequest)
		return
	}

	resPayload := JSONResponse{
		Error:   false,
		Message: "Successfully logged activity!",
		Data:    doc,
	}

	err = s.writeJSON(w, resPayload, http.StatusOK)
	if err != nil {
		s.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	log.Println("logger service: successful log")
}
