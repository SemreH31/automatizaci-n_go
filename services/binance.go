package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv" // Para convertir string a números
)

// Estructura para la petición a Binance
type P2PRequest struct {
	Asset          string   `json:"asset"`
	TradeType      string   `json:"tradeType"`
	Fiat           string   `json:"fiat"`
	TransLoggingNo string   `json:"transLoggingNo"`
	Page           int      `json:"page"`
	Rows           int      `json:"rows"`
	PayTypes       []string `json:"payTypes"`
}

func GetUSDTPrice(amount float64) (string, error) {
	url := "https://p2p.binance.com/bapi/c2c/v2/friendly/c2c/adv/search"

	requestBody, _ := json.Marshal(P2PRequest{
		Asset:     "USDT",
		TradeType: "BUY",
		Fiat:      "VES",
		Page:      1,
		Rows:      10,
		PayTypes:  []string{},
	})

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var result map[string]interface{}
	json.Unmarshal(body, &result)

	data, ok := result["data"].([]interface{})
	if !ok || len(data) == 0 {
		return "No disponible", fmt.Errorf("no se encontraron datos")
	}

	var minPrice, maxPrice, total float64
	count := 0

	for i, item := range data {
		adv := item.(map[string]interface{})["adv"].(map[string]interface{})
		priceStr := adv["price"].(string)
		price, _ := strconv.ParseFloat(priceStr, 64)

		if i == 0 {
			minPrice = price
		}
		maxPrice = price
		total += price
		count++
	}

	average := total / float64(count)

	// --- CÁLCULO DE LA CONVERSIÓN ---
	conversion := amount * minPrice

	// Formateamos el mensaje
	header := "📊 *Reporte P2P Binance (VES)*\n"
	if amount > 1 {
		header = fmt.Sprintf("📊 *Conversión: %.2f USDT a VES*\n💰 *Total: %.2f Bs*\n", amount, conversion)
	}

	response := fmt.Sprintf(
		"%s\n"+
			"🔹 *Mínimo:* %.2f Bs\n"+
			"🔸 *Máximo:* %.2f Bs\n"+
			"📈 *Promedio:* %.2f Bs\n\n"+
			"_(Basado en los mejores %d anuncios)_",
		header, minPrice, maxPrice, average, count,
	)

	return response, nil
}
