package rabbitmq

import (
    "context"
    "encoding/json"
    "log"
    "pharmacy/consumer/internal/models"
    "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
    conn    *amqp091.Connection
    channel *amqp091.Channel
    queue   amqp091.Queue
}

func NewConsumer(rabbitMQURL, queueName string) (*Consumer, error) {
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

    return &Consumer{conn, ch, q}, nil
}

func (c *Consumer) Consume(ctx context.Context, client interface {
    SendTransaction(models.Transaction) error
}) error {
    msgs, err := c.channel.Consume(
        c.queue.Name,
        "",
        true,
        false,
        false,
        false,
        nil,
    )
    if err != nil {
        return err
    }

    for {
        select {
        case <-ctx.Done():
            return nil
        case msg := <-msgs:
            var tx models.Transaction
            if err := json.Unmarshal(msg.Body, &tx); err != nil {
                log.Printf("Failed to unmarshal message: %v", err)
                continue
            }

            if err := client.SendTransaction(tx); err != nil {
                log.Printf("Failed to send transaction to third party: %v", err)
            }
        }
    }
}

func (c *Consumer) Close() error {
    if err := c.channel.Close(); err != nil {
        return err
    }
    return c.conn.Close()
}