package client

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type LogEntryClientImpl struct {
	Client *mongo.Client
}

func NewLogEntryClientImpl(client *mongo.Client) *LogEntryClientImpl {
	return &LogEntryClientImpl{Client: client}
}

func (l *LogEntryClientImpl) InsertLogEntry(logEntry LogEntry) error {
	coll := l.Client.Database("logs_db").Collection("logs")
	doc := LogEntry{Name: logEntry.Name, Data: logEntry.Data, CreatedAt: time.Now()}

	_, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Println("error inserting log entry:", err)
		return err
	}

	return nil
}
