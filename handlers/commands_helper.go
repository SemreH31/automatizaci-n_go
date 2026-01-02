package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleHelp(bot *tgbotapi.BotAPI, chatID int64) {
	texto := "📖 *Guía de Comandos del Bot*\n\n" +
		"🔹 *Redes y Seguridad:*\n" +
		"• `/scan [IP]` - Escanea puertos comunes (22, 80, 443, 8080).\n" +
		"• `/netscan` - (Próximamente) Escaneo de red local.\n\n" +
		"🔹 *Finanzas (P2P):*\n" +
		"• `/precio_ves` - Muestra el precio actual de USDT en Binance.\n" +
		"• `/precio_ves [cantidad]` - Calcula la conversión de USDT a Bolívares.\n\n" +
		"🔹 *Otros:*\n" +
		"• `/hola` - El bot te saluda por tu nombre.\n" +
		"• `/menu` - Muestra botones interactivos.\n" +
		"• `/ayuda` - Muestra esta lista de comandos.\n\n" +
		"🚀 *Día 1/365* - Proyecto de Redes y Go."

	msg := tgbotapi.NewMessage(chatID, texto)
	msg.ParseMode = "Markdown" // Para que las negritas y puntos se vean bien
	bot.Send(msg)
}
