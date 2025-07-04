package db

import (
    "github.com/jmoiron/sqlx"
    _ "github.com/go-sql-driver/mysql"
)

type Transaction struct {
    ID            int     `db:"id"`
    TransactionID string  `db:"transaction_id"`
    MedicineName  string  `db:"medicine_name"`
    Quantity      int     `db:"quantity"`
    Price         float64 `db:"price"`
    CreatedAt     string  `db:"created_at"`
}

type Database struct {
    db *sqlx.DB
}

func NewDatabase(dsn string) (*Database, error) {
    db, err := sqlx.Connect("mysql", dsn)
    if err != nil {
        return nil, err
    }
    return &Database{db}, nil
}

func (d *Database) SaveTransaction(tx Transaction) error {
    query := `
        INSERT INTO transactions (transaction_id, medicine_name, quantity, price)
        VALUES (:transaction_id, :medicine_name, :quantity, :price)
    `
    _, err := d.db.NamedExec(query, tx)
    return err
}

func (d *Database) Close() error {
    return d.db.Close()
}