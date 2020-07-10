package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

func main() {
	godotenv.Load()

	queue := flag.String("queue", "messages_que", "specifies which queue to listen")
	flag.Parse()

	rabbitmqURL := os.Getenv("RABBITMQ_URL")

	conn, err := amqp.Dial(rabbitmqURL)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		*queue,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Consuming '%v'\n", *queue)
	for msg := range msgs {
		fmt.Printf("Message: %s\n", msg.Body)
	}
}
