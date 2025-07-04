package main

import (
	"log"
	"net/http"
	"pharmacy/graphql/graph"
	"pharmacy/graphql/internal/config"
	"pharmacy/graphql/internal/db"
	"pharmacy/graphql/internal/rabbitmq"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

type App struct {
    config        *config.Config
    db            *db.Database
    rabbitPublisher *rabbitmq.Publisher
}

func NewApp() (*App, error) {
    cfg, err := config.NewConfig()
    if err != nil {
        return nil, err
    }

    database, err := db.NewDatabase(cfg.DatabaseURL)
    if err != nil {
        return nil, err
    }

    publisher, err := rabbitmq.NewPublisher(cfg.RabbitMQURL, cfg.QueueName)
    if err != nil {
        database.Close()
        return nil, err
    }

    return &App{
        config:         cfg,
        db:             database,
        rabbitPublisher: publisher,
    }, nil
}

func (app *App) Run() error {
    defer app.db.Close()
    defer app.rabbitPublisher.Close()

    resolver := &graph.Resolver{
        DB:        app.db,
        Publisher: app.rabbitPublisher,
    }

    srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

    http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
    http.Handle("/graphql", srv)

    log.Printf("Server running at http://localhost:%s/graphql", app.config.Port)
    return http.ListenAndServe(":"+app.config.Port, nil)
}

func main() {
    app, err := NewApp()
    if err != nil {
        log.Fatalf("Failed to initialize app: %v", err)
    }

    if err := app.Run(); err != nil {
        log.Fatalf("Application error: %v", err)
    }
}