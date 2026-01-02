package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func handleCallback(bot *tgbotapi.BotAPI, query *tgbotapi.CallbackQuery) {
	data := query.Data
	chatID := query.Message.Chat.ID

	// Respondemos según el botón presionado
	var respuesta string
	switch data {
	case "click_a":
		respuesta = "Seleccionaste la Opción A ✅"
	case "click_b":
		respuesta = "Seleccionaste la Opción B ✅"
	}
	// Enviamos la respuesta al chat
	msg := tgbotapi.NewMessage(chatID, respuesta)
	bot.Send(msg)

	// ¡Importante! Avisar a Telegram que ya procesamos el clic
	callback := tgbotapi.NewCallback(query.ID, "")
	bot.Request(callback)

}
