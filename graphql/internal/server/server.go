package server

import (
    "github.com/99designs/gqlgen/graphql/handler"
    "github.com/99designs/gqlgen/graphql/playground"
    "net/http"
    "pharmacy/graphql/internal/db"
    "pharmacy/graphql/internal/graphql"
    "pharmacy/graphql/internal/rabbitmq"
)

type Server struct {
    handler http.Handler
}

func NewServer(db *db.Database, publisher *rabbitmq.Publisher) *Server {
    resolver := graphql.NewResolver(db, publisher)
    srv := handler.NewDefaultServer(graphql.NewExecutableSchema(graphql.Config{Resolvers: resolver}))

    mux := http.NewServeMux()
    mux.Handle("/graphql", srv)
    mux.Handle("/", playground.Handler("GraphQL playground", "/graphql"))

    return &Server{mux}
}

func (s *Server) Handler() http.Handler {
    return s.handler
}