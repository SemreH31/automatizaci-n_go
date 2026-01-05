package services

import (
	"context"
	"fmt"
	"os"
	"time"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

func CreateEvent(summary string, startDateTime string) (string, error) {
	ctx := context.Background()

	// 1. Leer el JSON de la Service Account
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		return "", fmt.Errorf("no se encontró el archivo de la Service Account: %v", err)
	}

	// 2. Crear la configuración JWT (esto genera el token en memoria automáticamente)
	// No necesitas token.json, la librería firma la petición usando la Private Key del JSON
	config, err := google.JWTConfigFromJSON(b, calendar.CalendarEventsScope)
	if err != nil {
		return "", fmt.Errorf("error al procesar la llave privada: %v", err)
	}

	// 3. Crear el cliente y el servicio
	client := config.Client(ctx)
	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return "", fmt.Errorf("error conectando con Google Calendar: %v", err)
	}

	// 4. Parsing de la fecha (AAAA-MM-DD HH:MM)
	const layout = "2006-01-02 15:04"
	startTime, err := time.ParseInLocation(layout, startDateTime, time.Local)
	if err != nil {
		return "", fmt.Errorf("formato de fecha inválido: %v", err)
	}

	event := &calendar.Event{
		Summary: summary,
		Start: &calendar.EventDateTime{
			DateTime: startTime.Format(time.RFC3339),
			TimeZone: "America/Caracas", // Cámbialo por tu zona horaria
		},
		End: &calendar.EventDateTime{
			DateTime: startTime.Add(1 * time.Hour).Format(time.RFC3339),
			TimeZone: "America/Caracas",
		},
	}

	// 5. Insertar el evento
	// Nota: "primary" se refiere al calendario principal de la Service Account.
	// Si compartiste TU calendario personal con el bot, usa el email de tu cuenta personal aquí.
	calendarId := "primary"
	event, err = srv.Events.Insert(calendarId, event).Do()
	if err != nil {
		return "", fmt.Errorf("error de red al insertar evento: %v", err)
	}

	return event.HtmlLink, nil
}
