package main

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"syscall"

	"db2sem/internal/app"
	"db2sem/internal/config"
)

const configPath = "config.json"

func run() error {
	ctx, stopCtx := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stopCtx()

	cfg, err := config.Load(configPath)
	if err != nil {
		return fmt.Errorf("Load(%q): %w", configPath, err)
	}

	if err := app.Run(ctx, cfg); err != nil {
		return fmt.Errorf("Run(%+v): %w", cfg, err)
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
