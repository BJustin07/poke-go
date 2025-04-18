package controller

import (
	"net/http"
	"strings"

	"github.com/BJustin07/poke-go/service"
)

func GetPokemonHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	pokeURL := r.URL.Path

	pokemonName := strings.TrimPrefix(pokeURL, "/GetPokemon/")
	if pokemonName == "" || pokemonName == "/" {
		http.Error(w, "Pokemon name is required", http.StatusBadRequest)
		return
	}

	pokemon, err := service.GetPokemon(pokemonName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Pokemon name: " + pokemon.Name))

}
