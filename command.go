package main

import "fmt"

type Command struct {
	Hop string
	To string
	Description string
}

func (command Command) String() string {
	return fmt.Sprintf("'%s'", command.Hop)
}

func (command Command) Help() {
	fmt.Printf("Hop: %s \n", command.Hop)
	fmt.Printf("To: %s \n", command.To)
	fmt.Printf("Description: %s \n", command.Description)
	fmt.Printf("\n")
}