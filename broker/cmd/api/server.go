package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	logger "github.com/jateen67/broker/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type server struct {
	chi.Router
}

type authenticationPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loggerPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func newServer() *server {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
	}))

	s := &server{
		Router: r,
	}
	s.routes()

	return s
}

func (s *server) routes() {
	s.Router.Post("/", s.broker)
	s.Router.Post("/authentication", s.authentication)
	s.Router.Post("/grpc-logger", s.gRPCLogger)
	s.Router.Post("/rabbitmq-authentication", s.rabbitMQAuthentication)
}

func (s *server) broker(w http.ResponseWriter, r *http.Request) {

	resPayload := jsonResponse{
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

func (s *server) authentication(w http.ResponseWriter, r *http.Request) {
	var authPayload authenticationPayload

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

	var resJSON jsonResponse

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

func (s *server) gRPCLogger(w http.ResponseWriter, r *http.Request) {
	var logPayload loggerPayload

	err := s.readJSON(w, r, &logPayload)
	if logPayload.Name == "" || logPayload.Data == "" {
		s.errorJSON(w, errors.New("name and data must be non-empty"), http.StatusBadRequest)
		return
	}
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	conn, err := grpc.Dial("logger:50001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	defer conn.Close()

	client := logger.NewLoggerServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	res, err := client.LogActivity(ctx, &logger.LogRequest{
		Name: logPayload.Name,
		Data: logPayload.Data,
	})
	if err != nil {
		s.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	resJSON := jsonResponse{
		Error:   res.Error,
		Message: res.Message,
		Data:    res.LogEntry,
	}

	err = s.writeJSON(w, resJSON, http.StatusOK)
	if err != nil {
		s.errorJSON(w, err, http.StatusInternalServerError)
		return
	}
}

func (s *server) rabbitMQAuthentication(w http.ResponseWriter, r *http.Request) {}
