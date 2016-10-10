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