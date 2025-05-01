package main

import (
	"context"
	"github.com/cockroachdb/errors"
	"log"
	"main/internal/anti-spoof-client"
	"main/internal/bot"
	"main/internal/pkg/config"
	"main/internal/pkg/logger"
	"main/internal/update-handler"
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
	//1. Глобальный контекст
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	//2. Загрузка конфига
	conf, err := config.Load(config.ConfigPath)
	if err != nil {
		return errors.Wrap(err, "config.Load")
	}

	//3. Создание логгера
	zLog, err := logger.New(conf.Env)
	if err != nil {
		return errors.Wrap(err, "logger.New")
	}
	zLog.Infow("Logger and config initialized successfully")

	//4. Инициализация бота
	botExecutor, err := bot.New(conf.Bot, conf.Env, &zLog)
	if err != nil {
		return errors.Wrap(err, "New")
	}

	//5. Созадние http клиента для использования внешнего апи с моделями
	client := anti_spoof_client.NewClient(conf.App)

	//6. Создаем директорию для временных аудио файлов
	if err = os.MkdirAll(conf.App.AudioTempDir, os.ModePerm); err != nil {
		return errors.Wrap(err, "mkdir")
	}

	//7. Инициализация обработчика обновлений телеграм чатов
	updateHandler := update_handler.New(botExecutor, client, conf.App.AudioTempDir, zLog)

	//8. Запуск обработки обновлений
	if err = updateHandler.Run(ctx); err != nil {
		return errors.Wrap(err, "updateHandler.Run")
	}

	//9. Ожидание завершения контекста
	<-ctx.Done()

	zLog.Infow("Application shutting down...")
	return nil
}
