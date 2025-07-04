package rabbitmq

import (
	"context"
	"encoding/json"

	"github.com/rabbitmq/amqp091-go"
)

type Transaction struct {
    TransactionID string  `json:"transaction_id"`
    MedicineName  string  `json:"medicine_name"`
    Quantity      int     `json:"quantity"`
    Price         float64 `json:"price"`
}

type Publisher struct {
    conn    *amqp091.Connection
    channel *amqp091.Channel
    queue   amqp091.Queue
}

func NewPublisher(rabbitMQURL, queueName string) (*Publisher, error) {
    conn, err := amqp091.Dial(rabbitMQURL)
    if err != nil {
        return nil, err
    }

    ch, err := conn.Channel()
    if err != nil {
        conn.Close()
        return nil, err
    }

    q, err := ch.QueueDeclare(
        queueName,
        true,
        false,
        false,
        false,
        nil,
    )
    if err != nil {
        ch.Close()
        conn.Close()
        return nil, err
    }

    return &Publisher{conn, ch, q}, nil
}

func (p *Publisher) Publish(ctx context.Context, tx Transaction) error {
    body, err := json.Marshal(tx)
    if err != nil {
        return err
    }

    return p.channel.PublishWithContext(
        ctx,
        "",
        p.queue.Name,
        false,
        false,
        amqp091.Publishing{
            ContentType: "application/json",
            Body:        body,
        },
    )
}

func (p *Publisher) Close() error {
    if err := p.channel.Close(); err != nil {
        return err
    }
    return p.conn.Close()
}