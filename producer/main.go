package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

func main() {
	godotenv.Load()

	exchange := flag.String("exchange", "messages_ex", "specifies to which exchange the messagens will be sended")
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

	fmt.Printf("To send a new message to '%v', write the message and press ENTER:\n", *exchange)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		message := scanner.Bytes()

		ch.Publish(
			*exchange,
			"",
			false,
			false,
			amqp.Publishing{
				Type: "text/plain",
				Body: message,
			},
		)
	}
}
