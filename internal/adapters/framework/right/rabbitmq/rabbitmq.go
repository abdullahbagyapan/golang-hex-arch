package rabbitmq

import (
	"github.com/streadway/amqp"
	"log"
)

type Adapter struct {
	ch *amqp.Channel
	q  amqp.Queue
}

func NewMsgBrokerAdapter() *Adapter {
	conn, err := connectRabbitMQ()
	if err != nil {
		log.Fatalf("Error connecting to rabbitmq %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Error creating to channel %v", err)
	}

	// We create a Queue to send the message to.
	q, err := declareQueue(ch)
	if err != nil {
		log.Fatalf("Error declaring to queue %v", err)
	}

	return &Adapter{q: q, ch: ch}
}

func connectRabbitMQ() (*amqp.Connection, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	return conn, err
}

func registerConsumer(ch *amqp.Channel, q amqp.Queue) (<-chan amqp.Delivery, error) {
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	return msgs, err
}

func declareQueue(ch *amqp.Channel) (amqp.Queue, error) {
	q, err := ch.QueueDeclare(
		"hexarch", // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)

	return q, err
}

func (adapter Adapter) PublishMessage(msg string) error {
	err := adapter.ch.Publish(
		"",             // exchange
		adapter.q.Name, // routing key
		false,          // mandatory
		false,          // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})

	return err
}
