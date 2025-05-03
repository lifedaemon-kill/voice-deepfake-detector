package update_handler

import (
	"context"
	"github.com/cockroachdb/errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	antispoofclient "main/internal/anti-spoof-client"
	"main/internal/bot"
	"main/internal/pkg/logger"
	"os"
)

type UpdateHandler struct {
	bot          *bot.Bot
	client       *antispoofclient.Client
	tempAudioDir string
	zLog         logger.Logger
}

func New(bot *bot.Bot, client *antispoofclient.Client, tempAudioDir string, zLog logger.Logger) *UpdateHandler {
	return &UpdateHandler{
		bot:          bot,
		client:       client,
		tempAudioDir: tempAudioDir,
		zLog:         zLog,
	}
}

func (h *UpdateHandler) Run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			for update := range h.bot.GetUpdateChan() {
				if err := h.HandleUpdate(update); err != nil {
					h.zLog.Errorw("handler.HandleUpdate", "err", err)
				}
			}
		}
	}
}

func (h *UpdateHandler) HandleUpdate(update tgbotapi.Update) error {
	chatID := update.Message.Chat.ID
	message := update.Message.Text
	h.zLog.Infow("handle update", "chatID", update.Message.Chat.ID, "user", update.Message.From.UserName, "mimetype", update.Message)

	if message == "/help" || message == "/start" {
		err := h.bot.SendHelpMessage(chatID)
		h.zLog.Infow("send help", "err", err)
		return nil
	} else if message == "/license" {
		err := h.bot.SendLicenseMessage(chatID)
		h.zLog.Infow("send license", "err", err)
		return nil
	}

	if message != "" {
		err := h.bot.SendMessage(chatID, "Я работаю только с аудио-контентом")
		h.zLog.Infow("send message", "msg", update.Message, "err", err)
		return nil
	}

	var fileID string
	var mimetype string
	var filename string

	if update.Message.Voice != nil {
		fileID = update.Message.Voice.FileID
		mimetype = update.Message.Voice.MimeType
		filename = "Голосовое сообщение"
	} else if update.Message.Audio != nil {
		fileID = update.Message.Audio.FileID
		mimetype = update.Message.Audio.MimeType
		filename = update.Message.Audio.FileName
	} else {
		err := h.bot.SendMessage(chatID, "Данный тип файла не поддерживается, смотрите /help")
		h.zLog.Infow("Unsupported filetype", "err", err)
		return nil
	}
	h.zLog.Infow("Скачивание файла: ", "fileID", fileID, "mimetype", mimetype, "filename", filename)

	h.bot.SendMessage(chatID, "Обработка файла началась")
	audioPath, err := h.bot.DownloadAudioFile(fileID, mimetype, h.tempAudioDir)
	defer func() {
		h.zLog.Infow("Удаление файла " + audioPath)
		os.Remove(audioPath)
		h.zLog.Infow("Файл удалён " + audioPath)
	}()
	if err != nil {
		h.bot.SendMessage(chatID, "Произошла ошибка при обработке файла")
		return errors.Wrap(err, "DownloadAudioFile")
	}

	h.bot.SendMessage(chatID, "Расчеты моделей...")
	predict, err := h.client.SendRequest(audioPath)
	if err != nil {
		h.bot.SendMessage(chatID, "Ошибка во время работы моделей")
		return errors.New("[client.SendRequest] " + err.Error())
	}

	h.bot.SendMessage(chatID, predict.ToString(filename))
	h.zLog.Infow("update", "chatID", chatID)
	return nil
}
