package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type JSONResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

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
	s.Router.Post("/", s.Broker)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func (s *Server) writeJSON(w http.ResponseWriter, data JSONResponse) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	_, err = w.Write(out)

	if err != nil {
		return err
	}

	return nil
}

func (s *Server) Broker(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)

	payload := JSONResponse{
		Error:   false,
		Message: "Successfully hit the Broker!",
	}

	err := s.writeJSON(w, payload)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	fmt.Println("successful broker service call")
}
