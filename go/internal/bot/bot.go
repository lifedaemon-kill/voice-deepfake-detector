package bot

import (
	"github.com/cockroachdb/errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
	"io"
	"main/internal/pkg/config"
	"main/internal/pkg/logger"
	"net/http"
	"os"
	"path/filepath"
)

type Bot struct {
	bot  *tgbotapi.BotAPI
	zlog *logger.Logger
}

func New(conf config.BotConfig, env string, log *logger.Logger) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(conf.Token)
	if err != nil {
		return nil, errors.Wrap(err, "tgbotapi.NewBotAPI")
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

func (b *Bot) SendHelpMessage(chatID int64) error {
	msg := tgbotapi.NewMessage(chatID, help)
	b.bot.Send(msg)
	return nil
}

func (b *Bot) SendLicenceMessage(chatID int64) error {
	msg := tgbotapi.NewMessage(chatID, licence)
	b.bot.Send(msg)
	return nil
}

// DownloadFile возвращает путь до скачанного файла
func (b *Bot) DownloadFile(fileID, mimetype, filePath string) (string, error) {
	file, err := b.bot.GetFile(tgbotapi.FileConfig{FileID: fileID})
	if err != nil {
		return "", errors.Wrap(err, "tgbotapi.GetFile")
	}

	fileDirectURL, err := b.bot.GetFileDirectURL(file.FileID)
	if err != nil {
		return "", errors.Wrap(err, "tgbotapi.GetFileDirectURL")
	}

	response, err := http.Get(fileDirectURL)
	if err != nil {
		return "", errors.Wrap(err, "http.Get")
	}
	defer response.Body.Close()

	var extension string
	switch mimetype {
	case "audio/mpeg":
		extension = ".mp3"
	case "audio/ogg":
		extension = ".ogg"
	case "audio/wav":
		extension = ".wav"
	default:
		return "", errors.New("unsupported file type: " + extension)
	}

	audioPath := filePath + uuid.New().String() + extension
	out, err := os.Create(audioPath)
	defer out.Close()
	if err != nil {
		return "", errors.Wrap(err, "os.Create")
	}
	_, err = io.Copy(out, response.Body)

	absolutePath, err := filepath.Abs(audioPath)
	if err != nil {
		return "", errors.Wrap(err, "filepath.Abs")
	}

	return absolutePath, err
}
