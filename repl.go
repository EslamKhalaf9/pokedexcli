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
	callback    func() error
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
			name:        "map",
			description: "get 20 locations",
		},
	}
}

func startRepl() {
	scanner := bufio.NewScanner((os.Stdin))
	availableCommands := getCommands()
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		input := scanner.Text()
		cleaned := cleanInput(input)
		if len(cleaned) == 0 {
			continue
		}
		commandName := cleaned[0]
		command, ok := availableCommands[commandName]

		if !ok {
			fmt.Println("command not exist")
			os.Exit(0)
		}

		command.callback()
	}
}

func cleanInput(input string) []string {
	lowerCase := strings.ToLower(input)
	inputs := strings.Fields(lowerCase)
	return inputs
}
