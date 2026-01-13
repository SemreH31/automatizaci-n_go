package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type DolarVzlaResponse struct {
	Current struct {
		USD  float64 `json:"usd"`
		EUR  float64 `json:"eur"`
		Date string  `json:"date"`
	} `json:"current"`
	Previous struct {
		USD float64 `json:"usd"`
		EUR float64 `json:"eur"`
	} `json:"previous"`
	ChangePercentage struct {
		USD float64 `json:"usd"`
	} `json:"changePercentage"`
}

// GetDolarMessage obtiene ambos valores y retorna un string formateado para Telegram
func GetDolarMessage() (string, error) {
	url := "https://api.dolarvzla.com/public/exchange-rate"
	client := &http.Client{Timeout: 10 * time.Second}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 Bot-Habitos-Go")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result DolarVzlaResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	// Determinamos el emoji según si subió o bajó
	emoji := "📈"
	if result.ChangePercentage.USD < 0 {
		emoji = "📉"
	}

	mensaje := fmt.Sprintf(
		"📊 *Monitor de Divisas*\n\n"+
			"💵 *Dólar:* %.2f Bs\n"+
			"💶 *Euro:* %.2f Bs\n\n"+
			"%s *Variación:* %.2f%%\n"+
			"📅 *Fecha:* %s",
		result.Current.USD,
		result.Current.EUR,
		emoji,
		result.ChangePercentage.USD,
		result.Current.Date,
	)

	return mensaje, nil
}
