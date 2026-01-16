package main

import "github.com/luis13005/pos-go/pkg/rabbitmq"

func main() {
	ch, err := rabbitmq.OpenChannel()

	if err != nil {
		panic(err)
	}
	defer ch.Close()

	msg := "Olá, mundo!"
	exchange := "amq.direct"
	rabbitmq.Publish(ch, msg, exchange)
}
