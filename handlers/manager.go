package handlers

import (
	"mi_bot_telegram/services"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleUpdate(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	if update.Message != nil {
		handleMessage(bot, update.Message, update)
	} else if update.CallbackQuery != nil {
		handleCallback(bot, update.CallbackQuery)
	}
}

func handleMessage(bot *tgbotapi.BotAPI, msg *tgbotapi.Message, update *tgbotapi.Update) {
	args := strings.Fields(msg.Text)
	if len(args) == 0 {
		return
	}

	comando := args[0]

	switch comando {
	case "/start", "/ayuda", "/hola":
		handleStart(bot, msg, update) // Definida en otro archivo
	case "/menu":
		handleMenu(bot, update)
	case "/precio", "/poke", "/clima":
		handleFunCommands(bot, msg, comando, args) // Definida en commands_fun.go
	case "/scan":
		HandleNetworkScan(bot, msg, args) // Tu reto del Día 1
	case "/netscan":
		HandleNetScan(bot, msg)
	case "/arp":
		reporte := services.GetCleanARP()
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, reporte)
		msg.ParseMode = "Markdown"
		bot.Send(msg)
	case "/usdt_ves":
		handleFunCommands(bot, msg, comando, args)
	default:
		bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "Comando no reconocido"))
	}
}
