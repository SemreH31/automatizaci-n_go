package handlers

import (
	"mi_bot_telegram/services"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleDolar(bot *tgbotapi.BotAPI) {
	texto, err := services.GetDolarMessage()
	if err != nil {
		texto = "Error al obtener los datos del dólar."
	}
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, texto)
	msg.ParseMode = "Markdown"
	bot.Send(msg)
}
