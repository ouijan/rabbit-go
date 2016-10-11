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
	rabbit.Load()

	fmt.Println("Rabbit Command Line Hopper")
	
	args := os.Args[1:]
	matches := rabbit.Find(args)

	// If theres one run it
	if (len(matches) == 1) {
		matches[0].Run()
		return 
	}

	// If multiple print help
	help(matches)
	// If theres none display help for all
}

func help(commands []Command) {
	for _, command := range commands {
		printCommand(command);
	}
}

func printCommand (command Command) {
	fmt.Printf("Hop: %s \n", command.Hop)
	fmt.Printf("To: %s \n", command.To)
	fmt.Printf("Description: %s \n", command.Description)
	fmt.Printf("\n")
}