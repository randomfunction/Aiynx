package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
	"user-api/config"
	"user-api/db/sqlc"
	"user-api/internal/handler"
	"user-api/internal/logger"
	"user-api/internal/repository"
	"user-api/internal/routes"
	"user-api/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	// 1. Load Config
	cfg := config.LoadConfig()

	// 2. Init Logger
	logger.InitLogger()
	defer logger.Sync()

	// 3. Init DB
	ctx := context.Background()
	dbPool, err := pgxpool.New(ctx, cfg.DatabaseURL)
	if err != nil {
		logger.Log.Fatal("Unable to connect to database: " + err.Error())
	}
	defer dbPool.Close()

	if err := dbPool.Ping(ctx); err != nil {
		logger.Log.Fatal("Unable to ping database: " + err.Error())
	}

	// 4. Init Layers
	queries := sqlc.New(dbPool)
	repo := repository.NewUserRepository(queries)
	svc := service.NewUserService(repo)
	h := handler.NewUserHandler(svc)

	// 5. Setup Server
	app := fiber.New(fiber.Config{
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	})

	routes.SetupRoutes(app, h)

	// 6. Start Server
	go func() {
		if err := app.Listen(":" + cfg.Port); err != nil {
			logger.Log.Fatal("Server failed to start: " + err.Error())
		}
	}()

	// 7. Graceful Shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	logger.Log.Info("Gracefully shutting down...")
	if err := app.Shutdown(); err != nil {
		logger.Log.Error("Server forced to shutdown: " + err.Error())
	}

	logger.Log.Info("Server exited")
}
