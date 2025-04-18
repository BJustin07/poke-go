package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Pokemon struct {
	Name string `json:"name"`
}

func GetPokemon(pokeName string) (Pokemon, error) {
	res, err := http.Get("https://pokeapi.co/api/v2/pokemon/" + pokeName)
	if err != nil {
		return Pokemon{}, fmt.Errorf("Failed to get %s", pokeName)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusNotFound {
		return Pokemon{}, fmt.Errorf("Pokemon not found")
	} else if res.StatusCode != http.StatusOK {
		return Pokemon{}, fmt.Errorf("Error: %s", res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, fmt.Errorf("Error reading response body")
	}

	fmt.Println("Response Body:", string(body))

	var pokemon Pokemon
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return Pokemon{}, fmt.Errorf("Error unmarshalling JSON")
	}
	return pokemon, nil

}
