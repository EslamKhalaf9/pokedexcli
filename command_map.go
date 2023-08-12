package main

import (
	"fmt"
	"log"
	"pokedexcli/internal/pokeapi"
)

func callbackMap() error {
	client := pokeapi.NewClient()
	res, err := client.ListLocationAreas()
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println("Location areas:")
	for _, area := range res.Results {
		fmt.Printf("\t - %s \n", area.Name)
	}
	return nil
}
