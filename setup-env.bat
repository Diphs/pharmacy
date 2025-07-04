@echo off
echo Setting up environment variables for Pharmacy System...

REM GraphQL Service
set DATABASE_URL=root:password@tcp(localhost:3306)/pharmacy_db?parseTime=true
set RABBITMQ_URL=amqp://guest:guest@localhost:5672/
set QUEUE_NAME=transaction_queue
set PORT=8080

REM Consumer Service
set THIRD_PARTY_URL=http://localhost:8082/transactions

echo Environment variables set!
echo.
echo DATABASE_URL=%DATABASE_URL%
echo RABBITMQ_URL=%RABBITMQ_URL%
echo QUEUE_NAME=%QUEUE_NAME%
echo PORT=%PORT%
echo THIRD_PARTY_URL=%THIRD_PARTY_URL%
echo.
echo Now you can start the services in separate terminals:
echo 1. cd thirdparty_api && go run main.go
echo 2. cd consumer && go run main.go  
echo 3. cd graphql && go run main.go
pause
