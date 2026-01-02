package services

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Estructura para la respuesta de Open-Meteo
type WeatherData struct {
	CurrentWeather struct {
		Temperature float64 `json:"temperature"`
		Windspeed   float64 `json:"windspeed"`
		WeatherCode int     `json:"weathercode"`
	} `json:"current_weather"`
}

func GetWeather(lat, lon string) (string, error) {
	// URL para obtener clima actual en Celsius
	url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%s&longitude=%s&current_weather=true", lat, lon)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var data WeatherData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}

	// Traducimos el código de clima a algo legible
	estado := translateWeatherCode(data.CurrentWeather.WeatherCode)

	res := fmt.Sprintf("🌡️ Temperatura: %.1f°C\n💨 Viento: %.1f km/h\n☁️ Estado: %s",
		data.CurrentWeather.Temperature, data.CurrentWeather.Windspeed, estado)

	return res, nil
}

// Función auxiliar para entender los códigos de Open-Meteo
func translateWeatherCode(code int) string {
	switch code {
	case 0:
		return "Cielo despejado ☀️"
	case 1, 2, 3:
		return "Parcialmente nublado ⛅"
	case 45, 48:
		return "Niebla 🌫️"
	case 51, 53, 55:
		return "Llovizna 🌧️"
	case 61, 63, 65:
		return "Lluvia ☔"
	case 71, 73, 75:
		return "Nieve ❄️"
	case 95:
		return "Tormenta ⚡"
	default:
		return "Desconocido"
	}
}
