# Pharmacy Transaction System - Setup Guide

## Project Overview

This project is a distributed pharmacy transaction system designed to handle medication sales, persist transaction data, and communicate with a third-party API for external processing. It uses a microservices architecture with the following components:

- **GraphQL Service**: The main entry point for creating and retrieving transactions. It exposes a GraphQL API, saves data to a MySQL database, and publishes messages to a RabbitMQ queue.
- **Consumer Service**: Listens for messages on the RabbitMQ queue and forwards them to a mock third-party API.
- **Third-Party API**: A mock server that simulates an external service receiving transaction data.

## Tech Stack

- **Backend**: Go
- **API**: GraphQL (gqlgen)
- **Database**: MySQL
- **Message Broker**: RabbitMQ
- **Containerization**: Docker

---

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- **Docker** and **Docker Compose**: For running MySQL and RabbitMQ services.
- **Go**: Version 1.21 or higher.
- **Git**: For cloning the repository.

### 1. Clone the Repository

```bash
git clone <repository-url>
cd pharmacy
```

### 2. Set Up Environment Variables

Each service requires its own `.env` file for configuration. Example files are provided in the `graphql` and `consumer` directories. Copy them to create your local configuration:

**For the GraphQL Service:**

```powershell
copy .\graphql\.env.example .\graphql\.env
```

**For the Consumer Service:**

```powershell
copy .\consumer\.env.example .\consumer\.env
```

### 3. Start Services with Docker

We use Docker to run the necessary backend services (MySQL and RabbitMQ). A `docker-compose.yml` file is provided to simplify this process.

From the project root, run:

```bash
docker-compose up -d
```

This command will start MySQL and RabbitMQ in the background. To view logs, you can run `docker-compose logs -f`.

### 4. Run the Application

Use the provided batch script to start all Go services (GraphQL, Consumer, and Third-Party API) in separate terminal windows.

```powershell
.\start-services.bat
```

### 5. Test the System

After starting the services, the script will prompt you to press any key to run an end-to-end test. This test uses `test-graphql.ps1` to:

1.  Create a new transaction via the GraphQL API.
2.  Query the API to retrieve all transactions, including the new one.

If the tests run successfully, you will see the newly created transaction in the output.

---

## Manual Setup (Without Docker)

If you prefer to run MySQL and RabbitMQ natively, follow these steps:

### 1. Install and Start MySQL

- Install MySQL Server 8.0+.
- Ensure the `mysql` command-line client is available in your system's PATH.
- Create the database and table using the `schema.sql` file:

  ```shell
  mysql -u root -p < schema.sql
  ```

### 2. Install and Start RabbitMQ

- Install RabbitMQ using a package manager like Chocolatey (`choco install rabbitmq`) or download it from the official website.
- Start the RabbitMQ server.

### 3. Run the Application

Follow steps 2, 4, and 5 from the Docker-based setup.

---

## Troubleshooting

- **`Invoke-RestMethod : Unable to connect to the remote server`**: This error usually means the GraphQL service is not running or hasn't started yet. The `start-services.bat` script includes a delay to prevent this, but if it persists, check the GraphQL service terminal window for startup errors.
- **`mysql: The term 'mysql' is not recognized`**: This occurs during manual setup if the MySQL client is not in your system's PATH. The recommended solution is to use the Docker setup. If you must use a local installation, you will need to add the MySQL `bin` directory to your system's PATH environment variable.
- **Database connection errors**: Check that the MySQL container is running (`docker ps`) and that the credentials in `graphql/.env` match the `docker-compose.yml` file.
