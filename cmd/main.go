package main

import (
	"context"
	"github.com/cockroachdb/errors"
	"log"
	anti_spoof_client "main/internal/anti-spoof-client"
	"main/internal/bot"
	"main/internal/pkg/config"
	"main/internal/pkg/logger"
	update_handler "main/internal/update-handler"
	"os"
	"time"
)

func main() {
	ctx := context.Background()

	if err := bootstrap(ctx); err != nil {
		log.Fatalf("[main] bootstrap failed: %v", err)
	}

	<-ctx.Done()
	time.Sleep(1 * time.Second)
	os.Exit(0)
}

func bootstrap(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	conf, err := config.Load(config.Path)
	if err != nil {
		return errors.Wrap(err, "config.Load")
	}

	zLog, err := logger.New(conf.Env)
	if err != nil {
		return errors.Wrap(err, "logger.New")
	}
	zLog.Infow("Logger and config initialized successfully")

	botExecutor, err := bot.New(conf.Bot, conf.Env, &zLog)
	if err != nil {
		return errors.Wrap(err, "New")
	}

	client := anti_spoof_client.NewClient(conf.App)

	updateHandler := update_handler.New(botExecutor, client, zLog)

	if err = updateHandler.Run(ctx); err != nil {
		return err
	}

	<-ctx.Done()

	zLog.Infow("Application shutting down...")
	return nil
}
