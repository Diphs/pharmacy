# Pharmacy Transaction System

## Overview

This project is a distributed pharmacy transaction management system built with Go, consisting of three main services that communicate through RabbitMQ message queue. The system handles pharmacy transactions with GraphQL API endpoints, processes them asynchronously through a consumer service, and includes a mock third-party API service for testing.

## Architecture

The system follows a microservices architecture with the following components:

```
┌─────────────────┐    ┌─────────────┐    ┌─────────────────┐
│   GraphQL API   │───►│  RabbitMQ   │───►│    Consumer     │
│    Service      │    │   Queue     │    │    Service      │
└─────────────────┘    └─────────────┘    └─────────────────┘
         │                                          │
         ▼                                          ▼
┌─────────────────┐                    ┌─────────────────┐
│   MySQL DB      │                    │  Third Party    │
│                 │                    │   API (Mock)    │
└─────────────────┘                    └─────────────────┘
```

## Services

### 1. GraphQL Service (`/graphql`)

The GraphQL service provides the main API interface for the pharmacy system.

**Key Features:**

- GraphQL API with playground interface
- MySQL database integration for transaction storage
- RabbitMQ message publishing for asynchronous processing
- Transaction validation and error handling

**Technology Stack:**

- Go 1.21
- GraphQL with gqlgen (v0.17.31)
- MySQL with sqlx driver
- RabbitMQ AMQP client

**API Endpoints:**

- `POST /graphql` - GraphQL API endpoint
- `GET /` - GraphQL Playground interface

**GraphQL Schema:**

```graphql
type Transaction {
  id: ID!
  transactionId: String!
  medicineName: String!
  quantity: Int!
  price: Float!
  createdAt: String!
}

type Mutation {
  createTransaction(input: TransactionInput!): Transaction!
}

type Query {
  transactions: [Transaction!]!
}
```

### 2. Consumer Service (`/consumer`)

The consumer service processes transactions asynchronously from the RabbitMQ queue.

**Key Features:**

- RabbitMQ message consumption
- Third-party API integration
- Graceful shutdown handling
- Error handling and logging

**Technology Stack:**

- Go 1.21
- RabbitMQ AMQP client
- HTTP client for third-party integration

**Workflow:**

1. Consumes transaction messages from RabbitMQ queue
2. Deserializes JSON messages into Transaction structs
3. Forwards transactions to external third-party API
4. Handles errors and logs processing status

### 3. Third-Party API Service (`/thirdparty_api`)

A mock third-party API service that simulates an external pharmacy transaction processing system.

**Key Features:**

- Simple HTTP REST API
- Transaction logging and validation
- Mock response simulation for testing

**Technology Stack:**

- Go 1.21
- Standard HTTP library

**API Endpoints:**

- `POST /transactions` - Receives transaction data from consumer service

**Workflow:**

1. Receives transaction data via HTTP POST
2. Validates JSON payload
3. Logs transaction details
4. Returns success response

## Data Models

### Transaction Model

```go
type Transaction struct {
    TransactionID string  `json:"transaction_id"`
    MedicineName  string  `json:"medicine_name"`
    Quantity      int     `json:"quantity"`
    Price         float64 `json:"price"`
}
```

### Database Schema

```sql
CREATE TABLE transactions (
    id INT AUTO_INCREMENT PRIMARY KEY,
    transaction_id VARCHAR(255) NOT NULL,
    medicine_name VARCHAR(255) NOT NULL,
    quantity INT NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## Configuration

Both services use environment variables for configuration:

### GraphQL Service

- `DATABASE_URL`: MySQL connection string (default: `root:@tcp(localhost:3306)/pharmacy_db?parseTime=true`)
- `RABBITMQ_URL`: RabbitMQ connection string (default: `amqp://guest:guest@localhost:5672/`)
- `QUEUE_NAME`: RabbitMQ queue name (default: `transaction_queue`)
- `PORT`: HTTP server port (default: `8080`)

### Consumer Service

- `RABBITMQ_URL`: RabbitMQ connection string (default: `amqp://guest:guest@localhost:5672/`)
- `QUEUE_NAME`: RabbitMQ queue name (default: `transaction_queue`)
- `THIRD_PARTY_URL`: External API endpoint (default: `http://localhost:8082/transactions`)

## Project Structure

```
pharmacy/
├── consumer/                    # Consumer microservice
│   ├── go.mod
│   ├── go.sum
│   ├── main.go                 # Application entry point
│   └── internal/
│       ├── config/
│       │   └── config.go       # Configuration management
│       ├── models/
│       │   └── transaction.go  # Transaction model
│       ├── rabbitmq/
│       │   └── consumer.go     # RabbitMQ consumer logic
│       └── thirdparty/
│           └── client.go       # Third-party API client
├── graphql/                    # GraphQL API microservice
│   ├── go.mod
│   ├── go.sum
│   ├── main.go                 # Application entry point
│   └── internal/
│       ├── config/
│       │   └── config.go       # Configuration management
│       ├── db/
│       │   └── db.go           # Database operations
│       ├── graphql/
│       │   ├── schema.graphql  # GraphQL schema definition
│       │   ├── resolver.go     # GraphQL resolvers
│       │   ├── generated/      # Generated GraphQL code
│       │   └── models/         # GraphQL models
│       ├── rabbitmq/
│       │   └── publisher.go    # RabbitMQ publisher
│       └── server/
│           └── server.go       # HTTP server setup
├── thirdparty_api/             # Mock third-party API service
│   └── main.go                 # Simple HTTP server for testing
└── docs/                       # Documentation
    └── project-summary.md      # This file
```

## Key Features

1. **Asynchronous Processing**: Transactions are processed asynchronously through RabbitMQ, improving system responsiveness.

2. **GraphQL API**: Modern GraphQL interface with built-in playground for easy testing and exploration.

3. **Microservices Architecture**: Loosely coupled services that can be deployed and scaled independently.

4. **Graceful Shutdown**: Both services implement proper shutdown handling for production environments.

5. **Configuration Management**: Environment-based configuration with sensible defaults.

6. **Error Handling**: Comprehensive error handling and logging throughout the system.

7. **Database Integration**: Persistent storage using MySQL with proper transaction management.

8. **Mock Third-Party Service**: Includes a mock external API service for testing and development purposes.

## Usage Flow

1. **Transaction Creation**: Client submits a transaction via GraphQL mutation
2. **Validation**: GraphQL service validates the transaction data
3. **Persistence**: Transaction is saved to MySQL database
4. **Message Publishing**: Transaction details are published to RabbitMQ queue
5. **Consumption**: Consumer service picks up the message from the queue
6. **External Processing**: Consumer forwards the transaction to third-party API (mock service)
7. **Completion**: Transaction processing is complete

## Development Notes

- The system uses Go modules for dependency management
- Both services follow clean architecture principles with internal packages
- GraphQL is implemented using gqlgen with proper code generation
- Schema is defined in `graph/schema.graphqls` and generates models automatically
- The typo in the RabbitMQ publisher JSON tag has been fixed
- All services compile and build successfully

## Recent Updates

- **Fixed GraphQL Implementation**: Now uses proper gqlgen-generated code instead of manual implementation
- **Corrected Type Conversions**: Fixed int32/int type mismatches between GraphQL and database models
- **Updated Dependencies**: All required gqlgen packages are properly installed
- **Improved Error Handling**: Better error messages and proper error wrapping
- **Clean Architecture**: Separated concerns with proper dependency injection

## Future Enhancements

1. Add comprehensive logging and monitoring
2. Implement retry mechanisms for failed third-party API calls
3. Add authentication and authorization
4. Implement transaction querying functionality
5. Add unit and integration tests
6. Add Docker containerization
7. Implement health checks and metrics
8. Add API rate limiting
9. Implement transaction status tracking
10. Add data validation and sanitization improvements
