package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandHelp() error {
	fmt.Println("welcome to the Pokedex!")
	fmt.Println("usage")

	// temporary solution until figuring out the initialization cycle problem
	println("help: Displays a help message")
	println("exit: Exit the Pokedex")

	// initialization cycle for commands Error
	// for k, v := range commands {
	// 	fmt.Println("%s: %s", k, v.description)
	// }

	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}

var commands = map[string]cliCommand{
	"help": {
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	},
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
}

func main() {

	for {
		var input string
		fmt.Print("Pokedex > ")
		fmt.Scanln(&input)

		val, ok := commands[input]

		if !ok {
			fmt.Println("command not exist")
		}

		err := val.callback()

		if err != nil {
			return
		}
	}

}
