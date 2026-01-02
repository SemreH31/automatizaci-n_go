package services

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Definimos la estructura aquí
type BitcoinPrice struct {
	Bitcoin struct {
		USD float64 `json:"usd"`
	} `json:"bitcoin"`
}

func GetBitcoinPrice() (string, error) { // Nota la G mayúscula para que sea pública
	url := "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=usd"
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result BitcoinPrice
	json.NewDecoder(resp.Body).Decode(&result)
	return fmt.Sprintf("Precio BTC: $%.2f USD", result.Bitcoin.USD), nil
}
