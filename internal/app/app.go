package app

import (
	"context"
	"github.com/cockroachdb/errors"
	"main/internal/bot"
	"main/internal/handler"
	"main/internal/pkg/logger"
)

type App struct {
	bot     *bot.Bot
	zLog    logger.Logger
	handler handler.UpdateHandler
}

func New(bot *bot.Bot, zLog logger.Logger) *App {
	return &App{
		bot:  bot,
		zLog: zLog,
	}
}

func (app *App) Run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			for update := range app.bot.GetUpdateChan() {
				if err := app.handler.HandleUpdate(update, app.bot, app.zLog); err != nil {
					return errors.Wrap(err, "handler.HandleUpdate")
				}
			}
		}
	}
}
