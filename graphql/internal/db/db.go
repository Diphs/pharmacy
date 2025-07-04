package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
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

func (d *Database) GetAllTransactions() ([]Transaction, error) {
    query := `SELECT id, transaction_id, medicine_name, quantity, price, created_at FROM transactions ORDER BY created_at DESC`
    var transactions []Transaction
    err := d.db.Select(&transactions, query)
    return transactions, err
}

func (d *Database) Close() error {
    return d.db.Close()
}