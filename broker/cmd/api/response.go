package main

import (
	"encoding/json"
	"net/http"
)

type JSONResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func (s *Server) writeJSON(w http.ResponseWriter, data JSONResponse, status int) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)

	if err != nil {
		return err
	}

	return nil
}

func (s *Server) errorJSON(w http.ResponseWriter, err error, status int) error {
	resPayload := JSONResponse{
		Error:   true,
		Message: err.Error(),
	}

	return s.writeJSON(w, resPayload, status)
}
