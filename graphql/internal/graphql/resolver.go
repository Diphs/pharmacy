package graphql

import (
    "context"
    "errors"
    "pharmacy/graphql/internal/db"
    "pharmacy/graphql/internal/rabbitmq"
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

func (r *Resolver) Mutation() MutationResolver {
    return &mutationResolver{r}
}

func (r *Resolver) Query() QueryResolver {
    return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTransaction(ctx context.Context, input TransactionInput) (*Transaction, error) {
    // Validation
    if input.TransactionId == "" || input.MedicineName == "" {
        return nil, errors.New("transactionId and medicineName cannot be empty")
    }
    if input.Quantity <= 0 {
        return nil, errors.New("quantity must be positive")
    }
    if input.Price <= 0 {
        return nil, errors.New("price must be positive")
    }

    tx := db.Transaction{
        TransactionID: input.TransactionId,
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
        TransactionID: input.TransactionId,
        MedicineName:  input.MedicineName,
        Quantity:      input.Quantity,
        Price:         input.Price,
    }); err != nil {
        return nil, err
    }

    return &Transaction{
        TransactionId: tx.TransactionID,
        MedicineName:  tx.MedicineName,
        Quantity:      tx.Quantity,
        Price:         tx.Price,
    }, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Transactions(ctx context.Context) ([]*Transaction, error) {
    // Implement if needed
    return nil, nil
}