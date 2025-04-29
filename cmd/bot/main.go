package main

import (
	"context"
	"github.com/cockroachdb/errors"
	"log"
	"main/internal/app"
	"main/internal/bot"
	"main/internal/pkg/config"
	"main/internal/pkg/logger"
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

	// 3. Инициализация логгера
	zLog, err := logger.New(conf.Env)
	if err != nil {
		return errors.Wrap(err, "logger.New")
	}
	zLog.Infow("Logger and config initialized successfully")

	b, err := bot.New(conf.Bot, conf.Env, &zLog)
	if err != nil {
		return errors.Wrap(err, "New")
	}

	a := app.New(b, zLog)
	if err = a.Run(ctx); err != nil {
		return err
	}

	<-ctx.Done()

	zLog.Infow("Application shutting down...")
	return nil
}
