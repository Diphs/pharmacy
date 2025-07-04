@echo off
echo 🚀 Starting Pharmacy Transaction System...
echo ==========================================

echo.
echo 🔧 Starting services in separate windows...

echo Starting Third-party API...
start "Third-party API" cmd /k "cd thirdparty_api && go run main.go"

timeout /t 3 /nobreak >nul

echo Starting Consumer Service...
start "Consumer Service" cmd /k "cd consumer && go run main.go"

timeout /t 3 /nobreak >nul

echo Starting GraphQL Service...
start "GraphQL Service" cmd /c "cd graphql && go run main.go"

echo.
echo 🎉 All services started in separate windows!

:: Add a delay to give services time to start
echo.
echo ⏳ Waiting for services to initialize (5 seconds)...
timeout /t 5 /nobreak > nul

echo.
echo 📊 Service URLs:
echo    - GraphQL API: http://localhost:8080/graphql
echo    - GraphQL Playground: http://localhost:8080/
echo    - Third-party API: http://localhost:8082/transactions
echo    - RabbitMQ Management: http://localhost:15672

echo.
echo 🧪 To test the system:
echo    powershell -ExecutionPolicy Bypass -File test-graphql.ps1
echo.
echo Press any key to run the test...
pause > nul

:: Run the test script
powershell -ExecutionPolicy Bypass -File test-graphql.ps1

endlocal
