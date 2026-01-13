package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type DolarData struct {
	Promedio float64 `json:"promedio"`
	Fecha    string  `json:"fechaActualizacion"`
}

// GetDolarMessage obtiene ambos valores y retorna un string formateado para Telegram
func GetDolarMessage() (string, error) {
	client := &http.Client{Timeout: 10 * time.Second}

	// Consultamos ambos endpoints
	oficial, err := fetchDolar(client, "https://ve.dolarapi.com/v1/dolares/oficial")
	if err != nil {
		return "", fmt.Errorf("error BCV: %v", err)
	}

	paralelo, err := fetchDolar(client, "https://ve.dolarapi.com/v1/dolares/paralelo")
	if err != nil {
		return "", fmt.Errorf("error Paralelo: %v", err)
	}

	// Formateamos el mensaje con Markdown para que se vea bien en Telegram
	mensaje := fmt.Sprintf(
		"📊 *Control Cambiario*\n\n"+
			"🏛 *BCV:* %.2f Bs\n"+
			"💸 *Paralelo:* %.2f Bs\n\n"+
			"🕒 _Actualizado: %s_",
		oficial.Promedio,
		paralelo.Promedio,
		paralelo.Fecha[:10], // Cortamos la fecha para que sea legible
	)

	return mensaje, nil
}

func fetchDolar(client *http.Client, url string) (*DolarData, error) {
	req, _ := http.NewRequest("GET", url, nil)

	// Esto hace que tu programa parezca un navegador Chrome real
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: %d", resp.StatusCode)
	}

	var data DolarData
	json.NewDecoder(resp.Body).Decode(&data)
	return &data, nil
}
