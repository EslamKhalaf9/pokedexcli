package main

import "fmt"

func callbackHelp() error {
	fmt.Println("welcome to the Pokedex!")
	fmt.Println("usage")

	commands := getCommands()
	for _, command := range commands {
		fmt.Printf(" \t- %s: %s \n", command.name, command.description)
	}

	return nil
}
