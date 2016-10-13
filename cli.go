package main

import (
	"os"
	"fmt"
)

/**
 * Main Function
 */
func main() {
	rabbit := Rabbit{}
	rabbit.Load("rabbit.yaml")

	fmt.Println("Rabbit Command Line Hopper")
	args := os.Args[1:]
	matches := rabbit.Find(args)

	help(matches);
	// for _, cmd := range matches {
	// 	fmt.Println(cmd.Hop)
	// }
	return

	// If theres one run it
	if (len(matches) == 1) {
		matches[0].Run()
		return 
	} 

	// If multiple print help
	if (len(matches) > 1) {
		help(matches)
		return
	}

	// If theres none display help for all
	help(rabbit.Commands)
}

// help prints the commands given
func help(commands []Command) {
	for _, command := range commands {
		printCommand(command);
	}
}

// printCommand prints the given command
func printCommand (command Command) {
	fmt.Printf("Hop: %s \n", command.Hop)
	fmt.Printf("To: %s \n", command.To)
	if (command.Description != "") {
		fmt.Printf("Description: %s \n", command.Description)
	}
	fmt.Printf("\n")
}