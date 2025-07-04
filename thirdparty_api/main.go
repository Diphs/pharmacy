package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

type Transaction struct {
    TransactionID string  `json:"transaction_id"`
    MedicineName  string  `json:"medicine_name"`
    Quantity      int     `json:"quantity"`
    Price         float64 `json:"price"`
}

func handleTransaction(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
        return
    }

    var tx Transaction
    err := json.NewDecoder(r.Body).Decode(&tx)
    if err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    log.Printf("Received transaction: %+v", tx)
    w.WriteHeader(http.StatusCreated)
    fmt.Fprintf(w, "Transaction received successfully")
}

func main() {
    http.HandleFunc("/transactions", handleTransaction)
    log.Println("Mock API running on http://localhost:8082")
    log.Fatal(http.ListenAndServe(":8082", nil))
}