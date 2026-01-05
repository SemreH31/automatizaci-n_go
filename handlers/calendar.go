package handlers

import (
	"fmt"
	"mi_bot_telegram/services"
	"regexp"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// HandleCalendarEvent processes the /evento command.
// Expected format: /evento "Description" "YYYY-MM-DD HH:MM"
func HandleCalendarEvent(bot *tgbotapi.BotAPI, msg *tgbotapi.Message, args []string) {
	// Reconstruct the full message text to handle spaces inside quotes
	fullText := msg.Text

	// Regex to extract quotes content
	// Matches: /evento "something" "something else"
	re := regexp.MustCompile(`/evento\s+"(.*?)"\s+"(.*?)"`)
	matches := re.FindStringSubmatch(fullText)

	if len(matches) != 3 {
		reply := "Formato incorrecto. Uso:\n/evento \"Descripción\" \"AAAA-MM-DD HH:MM\""
		bot.Send(tgbotapi.NewMessage(msg.Chat.ID, reply))
		return
	}

	summary := matches[1]
	startDateTime := matches[2]

	link, err := services.CreateEvent(summary, startDateTime)
	if err != nil {
		reply := fmt.Sprintf("Error al crear evento: %v", err)
		bot.Send(tgbotapi.NewMessage(msg.Chat.ID, reply))
		return
	}

	reply := fmt.Sprintf("✅ Evento creado exitosamente:\n%s", link)
	msgResponse := tgbotapi.NewMessage(msg.Chat.ID, reply)
	msgResponse.DisableWebPagePreview = true
	bot.Send(msgResponse)
}
