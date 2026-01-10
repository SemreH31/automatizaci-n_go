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
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data DolarData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return &data, nil
}
