package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Pokemon struct {
	Name    string `json:"name"`
	Height  int    `json:"height"`
	Weight  int    `json:"weight"`
	Sprites struct {
		FrontDefault string `json:"front_default"`
	} `json:"sprites"`
}

func GetPokemon(name string) (string, string, error) {
	// Pasamos el nombre a minúsculas porque la API así lo requiere
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", strings.ToLower(name))

	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		return "", "", fmt.Errorf("no encontrado")
	}
	defer resp.Body.Close()

	var p Pokemon
	if err := json.NewDecoder(resp.Body).Decode(&p); err != nil {
		return "", "", err
	}

	info := fmt.Sprintf("¡Apareció %s! \nAltura: %.1fm \nPeso: %.1fkg",
		cases.Title(language.Spanish).String(p.Name), float64(p.Height)/10, float64(p.Weight)/10)

	// Retornamos la info y la URL de su foto
	return info, p.Sprites.FrontDefault, nil
}
