package bot

import (
	"bytes"
	"github.com/cockroachdb/errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"main/internal/pkg/config"
	"main/internal/pkg/logger"
)

type Bot struct {
	bot  *tgbotapi.BotAPI
	zlog *logger.Logger
}

func New(conf config.BotConfig, env string, log *logger.Logger) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(conf.Token)
	if err != nil {
		return nil, errors.Wrap(err, "New")
	}
	if env == "prod" {
		bot.Debug = false
	} else {
		bot.Debug = true
	}
	return &Bot{
		bot:  bot,
		zlog: log,
	}, nil
}

func (b *Bot) GetUpdateChan() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	return b.bot.GetUpdatesChan(u)
}
func (b *Bot) SendMessage(chatID int64, message string) error {
	msg := tgbotapi.NewMessage(chatID, message)
	b.bot.Send(msg)
	return nil
}
func (b *Bot) GetFileFromVoice(voice tgbotapi.Voice) *bytes.Buffer {
	//TODO

	//fileConfig := tgbotapi.FileConfig{FileID: voice.FileID}
	//file, err := b.bot.GetFile(fileConfig)
	//if err != nil {

	//}

	return nil
}
