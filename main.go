package main

import (
	"fmt"
	"log"
	"mi_bot_telegram/handlers"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	// 1. Intentar cargar el archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error: No se pudo encontrar el archivo .env")
	}

	// 2. Leer el token
	token := os.Getenv("TELEGRAM_APITOKEN")
	if token == "" {
		log.Fatal("Error: La variable TELEGRAM_TOKEN está vacía en el .env")
	}
	bot, _ := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	fmt.Println("hola funcionando")

	for update := range updates {
		if update.Message != nil {
			log.Printf("Mensaje de %s: %s", update.Message.From.UserName, update.Message.Text)
		}
		if update.CallbackQuery != nil {
			log.Printf("Botón presionado: %s", update.CallbackQuery.Data)
		}
		handlers.HandleUpdate(bot, &update)
	}
}
