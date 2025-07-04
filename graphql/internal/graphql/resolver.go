package graphql

import (
	"context"
	"errors"
	"pharmacy/graphql/internal/db"
	"pharmacy/graphql/internal/graphql/generated"
	"pharmacy/graphql/internal/graphql/models"
	"pharmacy/graphql/internal/rabbitmq"
	"strconv"
)

// THIS IS GENERATED CODE - DO NOT MODIFY
// Instead, modify schema.graphql and run `go run github.com/99designs/gqlgen generate`

type Resolver struct {
    db       *db.Database
    publisher *rabbitmq.Publisher
}

func NewResolver(db *db.Database, publisher *rabbitmq.Publisher) *Resolver {
    return &Resolver{db, publisher}
}

func (r *Resolver) Mutation() generated.MutationResolver {
    return &mutationResolver{r}
}

func (r *Resolver) Query() generated.QueryResolver {
    return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTransaction(ctx context.Context, input models.TransactionInput) (*models.Transaction, error) {
    // Validation
    if input.TransactionID == "" || input.MedicineName == "" {
        return nil, errors.New("transactionId and medicineName cannot be empty")
    }
    if input.Quantity <= 0 {
        return nil, errors.New("quantity must be positive")
    }
    if input.Price <= 0 {
        return nil, errors.New("price must be positive")
    }

    tx := db.Transaction{
        TransactionID: input.TransactionID,
        MedicineName:  input.MedicineName,
        Quantity:      input.Quantity,
        Price:         input.Price,
    }

    // Save to database
    if err := r.db.SaveTransaction(tx); err != nil {
        return nil, err
    }

    // Publish to RabbitMQ
    if err := r.publisher.Publish(ctx, rabbitmq.Transaction{
        TransactionID: input.TransactionID,
        MedicineName:  input.MedicineName,
        Quantity:      input.Quantity,
        Price:         input.Price,
    }); err != nil {
        return nil, err
    }

    return &models.Transaction{
        TransactionID: tx.TransactionID,
        MedicineName:  tx.MedicineName,
        Quantity:      tx.Quantity,
        Price:         tx.Price,
    }, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Transactions(ctx context.Context) ([]*models.Transaction, error) {
    dbTransactions, err := r.db.GetAllTransactions()
    if err != nil {
        return nil, err
    }

    var transactions []*models.Transaction
    for _, tx := range dbTransactions {
        transactions = append(transactions, &models.Transaction{
            ID:            strconv.Itoa(tx.ID),
            TransactionID: tx.TransactionID,
            MedicineName:  tx.MedicineName,
            Quantity:      tx.Quantity,
            Price:         tx.Price,
            CreatedAt:     tx.CreatedAt,
        })
    }

    return transactions, nil
}