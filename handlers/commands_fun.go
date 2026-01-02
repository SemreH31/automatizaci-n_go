package handlers

import (
	"mi_bot_telegram/services"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func handleFunCommands(bot *tgbotapi.BotAPI, msg *tgbotapi.Message, cmd string, args []string) {
	switch cmd {
	case "/precio":
		price, _ := services.GetBitcoinPrice()
		bot.Send(tgbotapi.NewMessage(msg.Chat.ID, price))

	case "/poke":
		if len(args) < 2 {
			bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "Dime el nombre de un Pokémon."))
			return
		}
		info, fotoUrl, err := services.GetPokemon(args[1])
		if err != nil {
			bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "No encontrado 🍃"))
			return
		}
		fotoMsg := tgbotapi.NewPhoto(msg.Chat.ID, tgbotapi.FileURL(fotoUrl))
		fotoMsg.Caption = info
		bot.Send(fotoMsg)

	case "/clima":
		lat := "10.0364"
		lon := "-68.0136"
		reporte, err := services.GetWeather(lat, lon)
		if err != nil {
			bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "Error al consultar el clima ❌"))
			return
		}
		msgClima := tgbotapi.NewMessage(msg.Chat.ID, "📍 Clima actual:\n"+reporte)
		bot.Send(msgClima)
	case "/usdt_ves":
		var cantidad float64 = 1 // Valor por defecto

		// Si el usuario escribió algo después del comando (ej: /precio_ves 10)
		if len(args) > 1 {
			// Intentamos convertir el segundo argumento a número
			val, err := strconv.ParseFloat(args[1], 64)
			if err == nil {
				cantidad = val
			}
		}

		// Llamamos al servicio con la cantidad (sea 1 o la que puso el usuario)
		reporte, err := services.GetUSDTPrice(cantidad)
		if err != nil {
			bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "Error al obtener precios ❌"))
			return
		}

		msgReply := tgbotapi.NewMessage(msg.Chat.ID, reporte)
		msgReply.ParseMode = "Markdown"
		bot.Send(msgReply)

	}

}
