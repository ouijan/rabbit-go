package main

import (
	"fmt"
	"os/exec"
	"os"
	"strings"
)

type Command struct {
	Hop string
	To string
	Description string
}

func (command *Command) String() string {
	return fmt.Sprintf("'%s'", command.Hop)
}

func (command *Command) Matches(args []string) bool {
	hopArray := strings.Split(command.Hop, " ")
	for i, arg := range args {
		if (arg != hopArray[i]) {
			return false
		}	
	}
	return true
}

func (command *Command) Run() {
	cmd := exec.Command("echo", "Command", "Run", "'" + command.Hop + "'")
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  cmd.Run()
}