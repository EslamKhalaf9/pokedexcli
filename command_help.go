package main

import "fmt"

func callbackHelp() {
	fmt.Println("welcome to the Pokedex!")
	fmt.Println("usage")

	// temporary solution until figuring out the initialization cycle problem
	println("help: Displays a help message")
	println("exit: Exit the Pokedex")

	// initialization cycle for commands Error
	// for k, v := range commands {
	// 	fmt.Println("%s: %s", k, v.description)
	// }

}
