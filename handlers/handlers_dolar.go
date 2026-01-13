package handlers

import (
	"mi_bot_telegram/services"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleDolar(bot *tgbotapi.BotAPI, msg *tgbotapi.Message, update *tgbotapi.Update) {
	texto, err := services.GetDolarMessage()
	if err != nil {
		texto = "Error al obtener los datos del dólar."
	}
	res := tgbotapi.NewMessage(update.Message.Chat.ID, texto+err.Error())
	res.ParseMode = "Markdown"
	bot.Send(res)
}
