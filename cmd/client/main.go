package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

	"github.com/LobovVit/GophKeeper/internal/client/config"
	"github.com/LobovVit/GophKeeper/internal/client/ui"
	"github.com/LobovVit/GophKeeper/pkg/logger"
)

var (
	buildVersion = "N/A"
	buildDate    = "N/A"
	buildCommit  = "N/A"
)

func main() {
	fmt.Printf("Build version: %s\n", buildVersion)
	fmt.Printf("Build date: %s\n", buildDate)
	fmt.Printf("Build commit: %s\n", buildCommit)
	if err := run(context.Background()); err != nil {
		panic(err)
	}
}

func run(ctx context.Context) error {
	cfg, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("get config: %w", err)
	}
	if err = logger.Initialize(cfg.LogLevel); err != nil {
		return fmt.Errorf("log initialize: %w", err)
	}
	ctx, cancel := signal.NotifyContext(ctx, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer cancel()
	app, err := ui.New(cfg, fmt.Sprintf(" client version %v-%v-%v ", buildVersion, buildDate, buildCommit))
	if err != nil {
		return fmt.Errorf("ui: %w", err)
	}
	logger.Log.Info("client initialized")
	return app.Run(ctx)
}
