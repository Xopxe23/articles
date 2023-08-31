package rabbitmq

import (
	"context"
	"encoding/json"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func faiOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

type LogItem struct {
	UserId    int
	ArticleId int
	Action    string
	Time      time.Time
}

type Client struct {
	conn *amqp.Connection
}

func NewClient() Client {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	faiOnError(err, "failed to connect to RabbitMQ")
	return Client{conn: conn}
}

func (c *Client) CloseConnection() error {
	return c.conn.Close()
}

func (c *Client) SendLog(logs LogItem) error {
	ch, err := c.conn.Channel()
	faiOnError(err, "failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare("logs", false, false, false, false, nil)
	faiOnError(err, "failed to declare a queue")

	ctx, cancel := context.WithTimeout(context.Background() ,10*time.Second)
	defer cancel()

	body, err := json.Marshal(logs)
	faiOnError(err, "failed to marshal log item")

	err = ch.PublishWithContext(ctx, "", q.Name, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body: body,
	})
	faiOnError(err, "failed to publish message")
	log.Print("[x] Sent log message")
	return nil
}
