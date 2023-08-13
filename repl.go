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
	callback    func(*config) error
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
			description: "get next locations page",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "get previous locations page",
			callback:    callbackMapb,
		},
	}
}

func startRepl(cfg *config) {
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
			continue
		}

		err := command.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(input string) []string {
	lowerCase := strings.ToLower(input)
	inputs := strings.Fields(lowerCase)
	return inputs
}
