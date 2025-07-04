#!/bin/bash

# Function to check if a service is running
check_service() {
    local url=$1
    local name=$2
    echo "Checking $name..."
    curl -s "$url" > /dev/null
    if [ $? -eq 0 ]; then
        echo "✅ $name is running"
        return 0
    else
        echo "❌ $name is not running"
        return 1
    fi
}

echo "🚀 Starting Pharmacy Transaction System..."
echo "=========================================="

# Check prerequisites
echo "📋 Checking prerequisites..."
check_service "http://localhost:3306" "MySQL" || echo "⚠️  Start MySQL first"
check_service "http://localhost:5672" "RabbitMQ" || echo "⚠️  Start RabbitMQ first"

echo ""
echo "🔧 Starting services..."

# Start Third-party API
echo "Starting Third-party API..."
cd thirdparty_api
go run main.go &
THIRDPARTY_PID=$!
cd ..

# Wait a bit
sleep 2

# Start Consumer Service
echo "Starting Consumer Service..."
cd consumer
go run main.go &
CONSUMER_PID=$!
cd ..

# Wait a bit
sleep 2

# Start GraphQL Service
echo "Starting GraphQL Service..."
cd graphql
go run main.go &
GRAPHQL_PID=$!
cd ..

echo ""
echo "🎉 All services started!"
echo "📊 Service URLs:"
echo "   - GraphQL API: http://localhost:8080/graphql"
echo "   - GraphQL Playground: http://localhost:8080/"
echo "   - Third-party API: http://localhost:8082/transactions"
echo "   - RabbitMQ Management: http://localhost:15672"
echo ""
echo "🧪 To test the system:"
echo "   powershell -ExecutionPolicy Bypass -File test-graphql.ps1"
echo ""
echo "🛑 To stop all services:"
echo "   kill $THIRDPARTY_PID $CONSUMER_PID $GRAPHQL_PID"

# Keep script running
wait
