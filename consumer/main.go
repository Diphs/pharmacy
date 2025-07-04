package main

import (
    "context"
    "log"
    "os"
    "os/signal"
    "pharmacy/consumer/internal/config"
    "pharmacy/consumer/internal/rabbitmq"
    "pharmacy/consumer/internal/thirdparty"
    "syscall"
)

type App struct {
    config       *config.Config
    rabbitConsumer *rabbitmq.Consumer
    thirdPartyClient *thirdparty.Client
}

func NewApp() (*App, error) {
    cfg, err := config.NewConfig()
    if err != nil {
        return nil, err
    }

    rabbitConsumer, err := rabbitmq.NewConsumer(cfg.RabbitMQURL, cfg.QueueName)
    if err != nil {
        return nil, err
    }

    thirdPartyClient := thirdparty.NewClient(cfg.ThirdPartyURL)

    return &App{
        config:           cfg,
        rabbitConsumer:   rabbitConsumer,
        thirdPartyClient: thirdPartyClient,
    }, nil
}

func (app *App) Run() error {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    errChan := make(chan error, 1)
    go func() {
        if err := app.rabbitConsumer.Consume(ctx, app.thirdPartyClient); err != nil {
            errChan <- err
        }
    }()

    // Handle graceful shutdown
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
    
    select {
    case err := <-errChan:
        return err
    case <-sigChan:
        log.Println("Received shutdown signal")
        cancel()
        return app.rabbitConsumer.Close()
    }
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