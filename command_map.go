package main

import (
	"fmt"
)

func callbackMap(cfg *config) error {
	res, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaPage)
	if err != nil {
		return err
	}
	fmt.Println("Location areas:")
	for _, area := range res.Results {
		fmt.Printf("\t - %s \n", area.Name)
	}

	cfg.nextLocationAreaPage = res.Next
	cfg.previousLocationAreaPage = res.Previous
	return nil
}

func callbackMapb(cfg *config) error {
	if cfg.previousLocationAreaPage == nil {
		return fmt.Errorf("you are already in the first page")
	}
	res, err := cfg.pokeapiClient.ListLocationAreas(cfg.previousLocationAreaPage)
	if err != nil {
		return err
	}
	fmt.Println("Location areas:")
	for _, area := range res.Results {
		fmt.Printf("\t - %s \n", area.Name)
	}

	cfg.nextLocationAreaPage = res.Next
	cfg.previousLocationAreaPage = res.Previous
	return nil
}
