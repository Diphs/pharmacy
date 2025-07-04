package main

import (
    "log"
    "net/http"
    "pharmacy/graphql/internal/config"
    "pharmacy/graphql/internal/db"
    "pharmacy/graphql/internal/rabbitmq"
    "pharmacy/graphql/internal/server"
)

type App struct {
    config        *config.Config
    db            *db.Database
    rabbitPublisher *rabbitmq.Publisher
    server        *server.Server
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

    srv := server.NewServer(database, publisher)

    return &App{
        config:         cfg,
        db:             database,
        rabbitPublisher: publisher,
        server:         srv,
    }, nil
}

func (app *App) Run() error {
    defer app.db.Close()
    defer app.rabbitPublisher.Close()

    log.Printf("Server running at http://localhost:%s/graphql", app.config.Port)
    return http.ListenAndServe(":"+app.config.Port, app.server.Handler())
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