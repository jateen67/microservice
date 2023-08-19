package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type JSONResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type Server struct {
	*mux.Router
}

func NewServer() *Server {
	s := &Server{
		Router: mux.NewRouter(),
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	s.HandleFunc("/", s.Broker).Methods("POST")
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
	payload := JSONResponse{
		Error:   false,
		Message: "Successfully hit the Broker!",
	}

	err := s.writeJSON(w, payload)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
