package main
import (
	"github.com/streadway/amqp"
	"fmt"
)

func sendToQueue(body string, rabbitHost string, qname string) error {
	conn, err := amqp.Dial(rabbitHost)
	failOnError(err, "Unable to connect to RabbitMQ")
	defer conn.Close()	
	
	ch, err := conn.Channel()
	failOnError(err, "Unable to open Channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		qname,
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare queue")

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:			[]byte(body),
	})

	if err != nil {
		fmt.Print("Unable to publish message")
		return err
	} 
	return err
}