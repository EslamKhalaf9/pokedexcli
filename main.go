package main

import (
	"fmt"
	"log"
	"pokedexcli/internal/pokeapi"
)

func main() {
	pokeapiClient := pokeapi.NewClient()

	res, err := pokeapiClient.ListLocationAreas()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
	// startRepl()
}
