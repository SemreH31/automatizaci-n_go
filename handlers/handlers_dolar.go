package handlers

import (
	"log"
	"mi_bot_telegram/services" // Asegúrate que la ruta a tu servicio sea correcta

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleDolar(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	// 1. Verificación de seguridad (Evita el panic)
	if bot == nil || message == nil {
		log.Println("Error: bot o message es nil")
		return
	}

	// 2. Obtener el mensaje de la API
	// Asumo que tu función en services devuelve (string, error)
	texto, err := services.GetDolarMessage()

	if err != nil {
		log.Printf("Error al obtener dólar: %v", err)
		texto = "❌ No se pudo obtener el valor del dólar en este momento."
	}

	// 3. Enviar el mensaje
	msg := tgbotapi.NewMessage(message.Chat.ID, texto)
	msg.ParseMode = "Markdown"

	if _, err := bot.Send(msg); err != nil {
		log.Printf("Error al enviar mensaje a Telegram: %v", err)
	}
}
