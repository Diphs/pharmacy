package config

import (
    "os"
)

type Config struct {
    DatabaseURL  string
    RabbitMQURL  string
    QueueName    string
    Port         string
}

func NewConfig() (*Config, error) {
    return &Config{
        DatabaseURL: getEnv("DATABASE_URL", "root:@tcp(localhost:3306)/pharmacy_db?parseTime=true"),
        RabbitMQURL: getEnv("RABBITMQ_URL", "amqp://guest:guest@localhost:5672/"),
        QueueName:   getEnv("QUEUE_NAME", "transaction_queue"),
        Port:        getEnv("PORT", "8080"),
    }, nil
}

func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}