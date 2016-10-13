package main

import (
	"fmt"
	"os/exec"
	"os"
	"strings"
)

// Command is the base Command Class
type Command struct {
	Hop string
	To string
	Description string
}

// String method used to convert struct to string
func (command *Command) String() string {
	return fmt.Sprintf("'%s'", command.Hop)
}

// hopArray returns the array of hop arguments
func (command *Command) hopArray() []string {
	return strings.Split(command.Hop, " ")
}

// MatchStrength calculates how well this command matches the given args
func (command *Command) MatchStrength(args []string) int {
	strength := 0
	for i, arg := range command.hopArray() {
		if (i >= len(args)) {
			return strength
		}
		if (arg == args[i]) {
			strength ++
		}
	}
	return strength
}

// IsExecutable determines the command is executable with the given args
func (command *Command) IsExecutable(args []string) bool {
	hopArray := command.hopArray();

	// If there isnt enough arguemnts return false
	if (len(hopArray) > len(args)) {
		return false
	}

	// If any arguments dont match return false
	for i, arg := range hopArray {
		if (arg != args[i]) {
			return false
		}
	}

	return true
}

// Run is used to execute the 'To' command
func (command *Command) Run() {
	cmd := exec.Command("echo", "Command", "Run", "'" + command.Hop + "'")
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  cmd.Run()
}