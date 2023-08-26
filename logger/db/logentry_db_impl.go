package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type LogEntryDBImpl struct {
	DB *mongo.Client
}

func NewLogEntryDBImpl(db *mongo.Client) *LogEntryDBImpl {
	return &LogEntryDBImpl{DB: db}
}

func (l *LogEntryDBImpl) InsertLogEntry(logEntry LogEntry) error {
	coll := l.DB.Database("logs_db").Collection("logs")
	doc := LogEntry{Name: logEntry.Name, Data: logEntry.Data, CreatedAt: time.Now()}

	_, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Println("error inserting log entry:", err)
		return err
	}

	return nil
}
