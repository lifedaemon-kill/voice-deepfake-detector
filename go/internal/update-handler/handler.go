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
					return errors.Wrap(err, "handler.HandleUpdate")
				}
			}
		}
	}
}

func (h *UpdateHandler) HandleUpdate(update tgbotapi.Update) error {
	chatID := update.Message.Chat.ID
	message := update.Message.Text

	if message == "/help" || message == "/start" {
		h.bot.SendHelpMessage(chatID)
		return nil
	} else if message == "/licence" {
		h.bot.SendLicenceMessage(chatID)
		return nil
	}

	if message != "" {
		h.bot.SendMessage(chatID, "я работаю только с аудио-контентом")
		return nil
	}

	var fileID string
	var mimetype string
	var filename string
	//Если пришло голосовое сообщение или аудио
	if update.Message.Voice != nil {
		fileID = update.Message.Voice.FileID
		mimetype = update.Message.Voice.MimeType
		filename = "Голосовое сообщение"
	} else if update.Message.Audio != nil {
		fileID = update.Message.Audio.FileID
		mimetype = update.Message.Audio.MimeType
		filename = update.Message.Audio.FileName
	} else {
		h.bot.SendMessage(chatID, "Данный тип файла не поддерживается, смотрите /help")
		return nil
	}

	audioPath, err := h.bot.DownloadFile(fileID, mimetype, h.tempAudioDir)
	defer func() {
		h.zLog.Infow("Удаление файла " + audioPath)
		os.Remove(audioPath)

		h.zLog.Infow("Файл удалён " + audioPath)
	}()
	if err != nil {
		h.bot.SendMessage(chatID, "Произошла ошибка при обработке файла, попробуйте еще или напишите на почту об ошибке")
		return errors.Wrap(err, "DownloadFile")
	}

	predict, err := h.client.SendRequest(audioPath)
	if err != nil {
		return errors.Wrap(err, "GetPredict")
	}

	h.bot.SendMessage(chatID, predict.ToString(filename))
	h.zLog.Infow("update", "chatID", chatID)
	return nil
}
