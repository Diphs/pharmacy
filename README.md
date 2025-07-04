# Pharmacy Transaction System - Setup Guide

## Prerequisites

- Go 1.21 or higher
- MySQL 8.0 or higher
- RabbitMQ Server
- Git

## Quick Start

### 1. Database Setup

Create a MySQL database and table:

```sql
CREATE DATABASE pharmacy_db;
USE pharmacy_db;

CREATE TABLE transactions (
    id INT AUTO_INCREMENT PRIMARY KEY,
    transaction_id VARCHAR(255) NOT NULL,
    medicine_name VARCHAR(255) NOT NULL,
    quantity INT NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### 2. RabbitMQ Setup

Install and start RabbitMQ:

```bash
# On Windows with Chocolatey
choco install rabbitmq

# Or download from https://www.rabbitmq.com/download.html
# Start RabbitMQ service
rabbitmq-server
```

### 3. Environment Variables

Set the following environment variables (or use defaults):

```bash
# GraphQL Service
DATABASE_URL=root:password@tcp(localhost:3306)/pharmacy_db?parseTime=true
RABBITMQ_URL=amqp://guest:guest@localhost:5672/
QUEUE_NAME=transaction_queue
PORT=8080

# Consumer Service
RABBITMQ_URL=amqp://guest:guest@localhost:5672/
QUEUE_NAME=transaction_queue
THIRD_PARTY_URL=http://localhost:8082/transactions
```

### 4. Build and Run

#### Build All Services

```bash
# Build GraphQL service
cd graphql
go build -o build/graphql.exe .

# Build Consumer service
cd ../consumer
go build -o build/consumer.exe .

# Build Third-party API (for testing)
cd ../thirdparty_api
go build -o build/thirdparty.exe .
```

#### Run Services

**Terminal 1 - Third-party API (Mock)**

```bash
cd thirdparty_api
go run main.go
```

**Terminal 2 - GraphQL API**

```bash
cd graphql
go run main.go
```

**Terminal 3 - Consumer Service**

```bash
cd consumer
go run main.go
```

### 5. Test the System

#### Create a Transaction

```bash
curl -X POST http://localhost:8080/graphql \
  -H "Content-Type: application/json" \
  -d '{
    "query": "mutation { createTransaction(input: { transactionId: \"123\", medicineName: \"Aspirin\", quantity: 2, price: 10.0 }) { id transactionId medicineName quantity price createdAt } }",
    "variables": {
      "input": {
        "transactionId": "123",
        "medicineName": "Aspirin",
        "quantity": 2,
        "price": 10.0
      }
    }
  }'
```

#### Get All Transactions

```bash
curl -X POST http://localhost:8080/graphql \
  -H "Content-Type: application/json" \
  -d '{
    "query": "query { transactions { id transactionId medicineName quantity price createdAt } }"
  }'
```

## Architecture Flow

1. **GraphQL API** receives transaction creation request
2. **Validates** transaction data
3. **Saves** transaction to MySQL database
4. **Publishes** message to RabbitMQ queue
5. **Consumer Service** receives message from queue
6. **Forwards** transaction to third-party API
7. **Third-party API** processes and logs transaction

## Service Ports

- **GraphQL API**: http://localhost:8080
- **Third-party API**: http://localhost:8082
- **RabbitMQ Management**: http://localhost:15672 (guest/guest)

## Development

### Adding New Fields

1. Update `schema.graphql` in the GraphQL service
2. Update database schema if needed
3. Update models and resolvers
4. Restart services

### Monitoring

- Check RabbitMQ management UI for queue status
- Monitor database for transaction records
- Check third-party API logs for processed transactions

## Troubleshooting

1. **Connection Issues**: Ensure MySQL and RabbitMQ are running
2. **Port Conflicts**: Change ports in configuration
3. **Database Errors**: Check database credentials and schema
4. **Queue Issues**: Verify RabbitMQ is accessible and queue exists
