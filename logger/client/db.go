package client

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Models struct {
	LogEntry LogEntry
}

func ConnectToClient() (*mongo.Client, error) {
	connString := os.Getenv("MONGO_CONNECTION_STRING")
	opts := options.Client().ApplyURI(connString)
	count := 1

	for {
		client, err := mongo.Connect(context.TODO(), opts)
		if err != nil {
			log.Println("could not connect to mongo. retrying... ")
			count++
		} else {
			return client, err
		}

		if count > 10 {
			return nil, err
		}

		log.Println("retrying in 1 second...")
		time.Sleep(1 * time.Second)
	}
}

func CreateCollection(client *mongo.Client) {

}
