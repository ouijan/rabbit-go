package main

import (
	"fmt"
	"os/exec"
	"os"
	"strings"
)

/**
 * Command is the base Command Class
 */
type Command struct {
	Hop string
	To string
	Description string
	Commands []Command
}

/**
 * String method used to convert struct to string
 */
func (command *Command) String() string {
	return fmt.Sprintf("'%s'", command.Hop)
}

/**
 * hopArray returns the array of hop arguments
 */
func (command *Command) hopArray() []string {
	return strings.Split(command.Hop, " ")
}

/**
 * toArray returns the array of hop arguments
 */
func (command *Command) toArray() []string {
	return strings.Split(command.To, " ")
}

/**
 * Run is used to execute the 'To' command
 */
func (command *Command) Run(args []string) error {
	// Build Executed args
	execArgs := command.toArray()
	execArgs = append(execArgs, args...)

	// Build Command
	cmd := exec.Command(execArgs[0], execArgs[1:]...)
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr

  // Run Command
  return cmd.Run()
}