package services

import (
	"encoding/json"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	UserID     string    `json:"userID"`
	Lat        float64   `json:"lat"`
	Lon        float64   `json:"lon"`
	SignUpDate time.Time `json:"signUpDate"`
}

func FetchAndStoreWeatherData(userID primitive.ObjectID, lat, lon float64, signUpDate time.Time) error {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"weather",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	msg := Message{
		UserID:     userID.Hex(),
		Lat:        lat,
		Lon:        lon,
		SignUpDate: signUpDate,
	}

	body, err := json.Marshal(msg)
	failOnError(err, "Failed to marshal message to JSON")

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	failOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s", body)
	return nil
}

func failOnError(err error, msg string) error {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		return err
	}
	return nil
}
