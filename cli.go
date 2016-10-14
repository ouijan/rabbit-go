package main

import (
	"os"
	"fmt"
)

/**
 * Main Function
 */
func main() {
	fmt.Println("Rabbit Command Line Hopper")
	configFile := "rabbit.yaml";

	rabbit := Rabbit{}
	err := rabbit.Load(configFile)
	if (err != nil) {
		fmt.Println("Missing " + configFile + " config file")
		os.Exit(1)
	}

	
	args := os.Args[1:]

	matches := rabbit.Find(args)
	exectuables := matches.FilterExecutables(args)

	// If one executable command run it
	if (len(exectuables) == 1) {
		err := matches[0].Run(args)
		if (err != nil) {
			os.Exit(1)
			return 
		}
		return
	} 

	// If multiple matches print help
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