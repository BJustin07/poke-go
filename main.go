package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/BJustin07/poke-go/controller"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloFromPostman)
	mux.HandleFunc("/GetPokemon/", controller.GetPokemonHandler)

	fmt.Println("Server started on port 8000")
	err := http.ListenAndServe(":8000", mux)

	if err != nil {
		log.Fatal(err)
	}

}

func helloFromPostman(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Postman!"))
}
