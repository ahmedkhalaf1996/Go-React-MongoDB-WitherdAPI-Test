package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/ahmedkhalaf1996/Go-React-MongoDB-WitherdAPI-Test-WeatherApi/services"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	UserID     string    `json:"userID"`
	Lat        float64   `json:"lat"`
	Lon        float64   `json:"lon"`
	SignUpDate time.Time `json:"signUpDate"`
}

func main() {
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

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan struct{})

	go func() {
		for d := range msgs {
			var msg Message
			err := json.Unmarshal(d.Body, &msg)
			if err != nil {
				log.Printf("Failed to decode message: %v", err)
			} else {
				log.Printf("Received a message:")
				log.Printf("  UserID: %s", msg.UserID)
				log.Printf("  Lat: %f", msg.Lat)
				log.Printf("  Lon: %f", msg.Lon)
				log.Printf("  SignUpDate: %s", msg.SignUpDate)
				// Call the Service
				userID, err := primitive.ObjectIDFromHex(msg.UserID)
				if err != nil {
					fmt.Println("Invalid UserID:", err)
				}
				go func() {
					if err := services.FetchAndStoreWeatherData(userID, msg.Lat, msg.Lon, msg.SignUpDate); err != nil {
						log.Print("error : Error while creating wather data.")

						return
					}
				}()
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
