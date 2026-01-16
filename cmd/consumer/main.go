package main

import (
	"fmt"

	"github.com/luis13005/pos-go/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}

	defer ch.Close()

	msgs := make(chan amqp.Delivery)
	queue := "minhafila"
	go rabbitmq.Consume(ch, msgs, queue)

	for msg := range msgs {
		fmt.Println(string(msg.Body))
		msg.Ack(false)
	}
}
