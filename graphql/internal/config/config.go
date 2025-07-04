package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL  string
	RabbitMQURL  string
	QueueName    string
	Port         string
	ThirdPartyURL string
}

func NewConfig() (*Config, error) {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	port := os.Getenv("GRAPHQL_PORT")
	if port == "" {
		port = "8080"
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Format: user:password@tcp(host:port)/dbname?parseTime=true
	databaseURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbUser, dbPassword, dbHost, dbPort, dbName,
	)

	rabbitMQURL := os.Getenv("RABBITMQ_URL")
	if rabbitMQURL == "" {
		rabbitMQURL = "amqp://guest:guest@localhost:5672/"
	}

	queueName := os.Getenv("QUEUE_NAME")
	if queueName == "" {
		queueName = "transactions"
	}

	return &Config{
		Port:          port,
		DatabaseURL:   databaseURL,
		RabbitMQURL:   rabbitMQURL,
		QueueName:     queueName,
		ThirdPartyURL: os.Getenv("THIRDPARTY_URL"),
	}, nil
}