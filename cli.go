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

	fmt.Println("Rabbit Command Line Hopper \n")
	help(rabbit)
	
	arguments := os.Args[1:]
	fmt.Println(arguments)
}



func help(rabbit Rabbit) {
	for _, command := range rabbit.Commands {
		printCommand(command);
	}
}

func printCommand (command Command) {
	fmt.Printf("Hop: %s \n", command.Hop)
	fmt.Printf("To: %s \n", command.To)
	fmt.Printf("Description: %s \n", command.Description)
	fmt.Printf("\n")
}