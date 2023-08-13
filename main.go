package main

import "pokedexcli/internal/pokeapi"

type config struct {
	pokeapiClient            pokeapi.Client
	nextLocationAreaPage     *string
	previousLocationAreaPage *string
}

func main() {
	cfg := config{
		pokeapiClient:            pokeapi.NewClient(),
		nextLocationAreaPage:     nil,
		previousLocationAreaPage: nil,
	}
	startRepl(&cfg)
}
