package handlers

import (
	"fmt"
	"mi_bot_telegram/services"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleNetworkScan(bot *tgbotapi.BotAPI, msg *tgbotapi.Message, args []string) {
	if len(args) < 2 {
		bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "Por favor, dime una IP. Ejemplo: /scan 1.1.1.1"))
		return
	}

	target := args[1]
	bot.Send(tgbotapi.NewMessage(msg.Chat.ID, fmt.Sprintf("🔍 Escaneando puertos comunes en %s...", target)))

	// Lista de puertos a revisar
	puertos := []int{22, 80, 443, 8080}
	resultados := ""

	for _, p := range puertos {
		abierto := services.ScanPort(target, p)
		if abierto {
			resultados += fmt.Sprintf("✅ Puerto %d: ABIERTO\n", p)
		} else {
			resultados += fmt.Sprintf("❌ Puerto %d: Cerrado\n", p)
		}
	}

	bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "Resultados para "+target+":\n\n"+resultados))
}

func HandleNetScan(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "🔍 Escaneando red local (192.168.1.0/24)... esto tardará unos segundos."))

	var wg sync.WaitGroup
	results := make(chan string, 254)
	baseIP := "192.168.1" // Cambia esto según el rango de tu casa

	for i := 1; i <= 254; i++ {
		ip := fmt.Sprintf("%s.%d", baseIP, i)
		wg.Add(1)
		go services.CheckHost(ip, &wg, results) // Lanzamos 254 escaneos en paralelo
	}

	// Cerramos el canal cuando todos terminen
	go func() {
		wg.Wait()
		close(results)
	}()

	encontrados := ""
	for ip := range results {
		encontrados += fmt.Sprintf("🖥 Host detectado: `%s`\n", ip)
	}

	if encontrados == "" {
		encontrados = "No se detectaron dispositivos abiertos."
	}

	bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "✅ *Escaneo Finalizado*\n\n"+encontrados))
}
