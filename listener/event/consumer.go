package event

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	conn *amqp.Connection
}

type Payload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewConsumer(conn *amqp.Connection) (Consumer, error) {

	consumer := Consumer{
		conn: conn,
	}

	// set up the consumer by opening up a channel and declaring an exchange
	channel, err := consumer.conn.Channel()
	if err != nil {
		return Consumer{}, err
	}

	err = channel.ExchangeDeclare(
		"logs_topic",
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return Consumer{}, err
	}

	return consumer, nil
}

// listens to the queue for specific topics
func (consumer *Consumer) Listen(topics []string) error {
	// go to our consumer channel and get things from it
	ch, err := consumer.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	// we have our channel, now we need to get a random queue
	q, err := ch.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	// we have our channel and our queue now

	// go through our list of topics
	for _, s := range topics {
		// bind our channel to each of these topics
		err = ch.QueueBind(
			q.Name,
			s,
			"logs_topic",
			false,
			nil,
		)
		if err != nil {
			return err
		}
	}

	// look for messages
	messages, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	// declaring a new channel
	forever := make(chan bool)
	// will run in background
	go func() {
		for d := range messages {

			var payload Payload
			_ = json.Unmarshal(d.Body, &payload)

			go handlePayload(payload)
		}
	}()

	fmt.Printf("waiting for message [exchange, queue] [logs_topic, %s]\n", q.Name)
	// keep the consumption going forever by making this blocking
	<-forever

	return nil
}

func handlePayload(payload Payload) {
	err := logEvent(payload)
	if err != nil {
		log.Println(err)
	}
}

func logEvent(entry Payload) error {

	jsonData, _ := json.Marshal(entry)

	request, err := http.NewRequest("POST", "http://authentication/authentication", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// make sure we get the correct status code from the logger service
	if res.StatusCode != http.StatusAccepted {
		return err
	}

	return nil
}
