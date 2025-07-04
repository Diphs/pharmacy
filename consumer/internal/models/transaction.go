package models

type Transaction struct {
    TransactionID string  `json:"transaction_id"`
    MedicineName  string  `json:"medicine_name"`
    Quantity      int     `json:"quantity"`
    Price         float64 `json:"price"`
}