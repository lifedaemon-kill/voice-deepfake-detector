package command_handler

import (
	"github.com/cockroachdb/errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	anti_spoof_client "main/internal/anti-spoof-client"
	"main/internal/bot"
	"main/internal/pkg/logger"
)

type UpdateHandler struct {
	bot  bot.Bot
	ai   anti_spoof_client.Service
	zLog logger.Logger
}

func (h *UpdateHandler) HandleUpdate(update tgbotapi.Update, bot *bot.Bot, zlog logger.Logger) error {
	chatID := update.Message.Chat.ID
	message := update.Message.Text

	if message != "" {
		bot.SendMessage(chatID, "я работаю только с медиа-контентом, не отправляйте сообщения")
	}
	var audioPath string

	if update.Message.Voice != nil {
		//audioPath = bot.GetFileFromVoice()
	} else if update.Message.Audio != nil {
		// audioPath = bot.GetFileFromAudio
	} else if update.Message.Video != nil {
		// audioPath = bot.GetFileFromVideo()
	}

	predict, err := h.ai.GetPredict(audioPath)
	if err != nil {
		return errors.Wrap(err, "GetPredict")
	}

	bot.SendMessage(chatID, predict.ToString())
	zlog.Infow("update", "chatID", chatID)
	return nil
}
