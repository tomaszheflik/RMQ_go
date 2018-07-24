package main

import (
	"github.com/streadway/amqp"
)

func readFromQueue(rabbitHost string, queueName string, mchan chan Message) {
	var message Message 
	conn, err := amqp.Dial(rabbitHost)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName, // name
		true,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")
	
	ack := make(chan bool)
	message.ACK = ack
	for {
		for d := range msgs {
			message.BODY = string(d.Body)
			mchan <- message
			if <-ack {
				d.Ack(true)
			}
		}
	}
	
}