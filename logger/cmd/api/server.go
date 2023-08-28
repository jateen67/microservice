package main

import (
	"context"
	"log"
	"time"

	"github.com/jateen67/logger/client"
	logger "github.com/jateen67/logger/protos"
)

type Server struct {
	logger.UnimplementedLoggerServiceServer
	loggerClient client.LogEntryClient
}

func (s *Server) WriteLog(ctx context.Context, req *logger.LogRequest) (*logger.LogResponse, error) {
	doc := client.LogEntry{
		Name:      req.Name,
		Data:      req.Data,
		CreatedAt: time.Now(),
	}

	err := s.loggerClient.InsertLogEntry(doc)
	if err != nil {
		log.Fatalf("failed to insert log into database: %s", err)
		return nil, err
	}

	res := &logger.LogResponse{
		Error:   false,
		Message: "Succesfully logged activity!",
		LogEntry: &logger.LogEntry{
			Name:      doc.Name,
			Data:      doc.Data,
			CreatedAt: doc.CreatedAt.Format(time.RFC3339),
		},
	}

	return res, nil
}
