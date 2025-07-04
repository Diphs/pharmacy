# Test GraphQL Mutation - Create Transaction
$mutation = @"
{
  "query": "mutation CreateTransaction(`$input: TransactionInput!) { createTransaction(input: `$input) { id transactionId medicineName quantity price createdAt } }",
  "variables": {
    "input": {
      "transactionId": "TXN-$(Get-Date -Format 'yyyyMMddHHmmss')",
      "medicineName": "Aspirin",
      "quantity": 2,
      "price": 15.50
    }
  }
}
"@

Write-Host "Testing GraphQL Mutation - Create Transaction..." -ForegroundColor Green
Write-Host "Request:" -ForegroundColor Yellow
Write-Host $mutation -ForegroundColor Cyan

$response = Invoke-RestMethod -Uri "http://localhost:8080/graphql" -Method Post -Body $mutation -ContentType "application/json"

Write-Host "`nResponse:" -ForegroundColor Yellow
$response | ConvertTo-Json -Depth 10 | Write-Host -ForegroundColor Cyan

# Test GraphQL Query - Get All Transactions
$query = @"
{
  "query": "query { transactions { id transactionId medicineName quantity price createdAt } }"
}
"@

Write-Host "`n`nTesting GraphQL Query - Get All Transactions..." -ForegroundColor Green
Write-Host "Request:" -ForegroundColor Yellow
Write-Host $query -ForegroundColor Cyan

$response = Invoke-RestMethod -Uri "http://localhost:8080/graphql" -Method Post -Body $query -ContentType "application/json"

Write-Host "`nResponse:" -ForegroundColor Yellow
$response | ConvertTo-Json -Depth 10 | Write-Host -ForegroundColor Cyan
