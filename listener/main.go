package main

import (
	"log"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {

	conn, err := connectToRabbitMQ()
	if err != nil {
		log.Panicln(err)
	}
	defer conn.Close()

	log.Println("successfully connected to rabbitmq")
}

func connectToRabbitMQ() (*amqp.Connection, error) {

	count := 1

	for {
		conn, err := amqp.Dial(os.Getenv("RABBITMQ_CONNECTION_STRING"))
		if err != nil {
			log.Println("rabbitmq not yet ready...")
			count++
		} else {
			return conn, nil
		}

		if count > 10 {
			log.Println(err)
			return nil, err
		}

		log.Println("retrying in 2 seconds...")
		time.Sleep(2 * time.Second)
	}
}
