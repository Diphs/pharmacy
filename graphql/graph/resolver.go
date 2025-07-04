package graph

import (
	"pharmacy/graphql/internal/db"
	"pharmacy/graphql/internal/rabbitmq"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	DB        *db.Database
	Publisher *rabbitmq.Publisher
}
