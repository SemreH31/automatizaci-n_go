package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func handleStart(bot *tgbotapi.BotAPI, msg *tgbotapi.Message, update *tgbotapi.Update) {

	switch msg.Text {
	case "/ayuda", "/start": // El comando /start suele mostrar la ayuda al inicio

		msg.Text = "¡Hola! Soy tu bot en Go."
		reply := tgbotapi.NewMessage(msg.Chat.ID, msg.Text)
		bot.Send(reply)
		HandleHelp(bot, msg.Chat.ID)

	case "/hola":
		msg.Text = ("¡Hola, " + update.Message.From.FirstName + "! 👋")
		reply := tgbotapi.NewMessage(msg.Chat.ID, msg.Text)
		bot.Send(reply)

	}
}
