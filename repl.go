package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func()
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    callbackExit,
		},
		"map": {
			name: "map",
		},
	}
}

func startRepl() {
	for {
		scanner := bufio.NewScanner((os.Stdin))
		fmt.Print("Pokedex > ")
		scanner.Scan()

		input := scanner.Text()
		cleaned := cleanInput(input)
		command := cleaned[0]
		val, ok := getCommands()[command]

		if !ok {
			fmt.Println("command not exist")
		}

		val.callback()
	}
}

func cleanInput(input string) []string {
	lowerCase := strings.ToLower(input)
	inputs := strings.Split(lowerCase, " ")
	return inputs
}
