package config

import (
    "os"
)

type Config struct {
    RabbitMQURL    string
    QueueName      string
    ThirdPartyURL  string
}

func NewConfig() (*Config, error) {
    return &Config{
        RabbitMQURL:   getEnv("RABBITMQ_URL", "amqp://guest:guest@localhost:5672/"),
        QueueName:     getEnv("QUEUE_NAME", "transaction_queue"),
        ThirdPartyURL: getEnv("THIRD_PARTY_URL", "http://localhost:8082/transactions"),
    }, nil
}

func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}