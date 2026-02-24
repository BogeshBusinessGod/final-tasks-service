package main

import (
	"context"
	log "final/internal/utils/observability"
	"os"
	"os/signal"
	"syscall"
	"time"

	"final/internal/app/final/task"
	"final/internal/config"
	"final/internal/repository/postgres"
	"final/internal/service"
)

func main() {
	logger := log.NewLogger(log.LevelDebug)
	logger.Info("Starting TasksService...")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, syscall.SIGINT, syscall.SIGTERM)

	cfg, err := config.Load()
	if err != nil {
		logger.Error("failed to load config", err)
		os.Exit(1)

	}

	db, err := postgres.NewPostgres(ctx, logger, cfg.Postgres)
	if err != nil {
		logger.Error("failed to connect postgres", err)
		os.Exit(1)
	}
	defer db.Close()

	svc := service.NewService(db)
	server := task.NewServer(cfg, logger, svc)

	go func() {
		if err := server.Listen(); err != nil {
			logger.Error("failed to listen service", err)
			os.Exit(1)
		}
	}()

	<-stopCh
	logger.Warn("shutdown signal received")

	shutdownCtx, cancelShutdown := context.WithTimeout(ctx, 5*time.Second)
	defer cancelShutdown()

	if err := server.Stop(shutdownCtx); err != nil {
		logger.Error("failed to stop server", err)
	}

	logger.Info("Server stopped cleanly")
}
