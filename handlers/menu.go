package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func handleMenu(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	filas := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			// Texto del botón, Datos que recibe el bot
			tgbotapi.NewInlineKeyboardButtonData("Opción A", "click_a"),
			tgbotapi.NewInlineKeyboardButtonData("Opción B", "click_b"),
		),
	)
	reply := tgbotapi.NewMessage(update.Message.Chat.ID, "¿Qué prefieres?")

	reply.ReplyMarkup = filas
	bot.Send(reply)
}
