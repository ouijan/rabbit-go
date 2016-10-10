package main

import (
	"fmt"
	"os/exec"
	"os"
)

type Command struct {
	Hop string
	To string
	Description string
}

func (command *Command) String() string {
	return fmt.Sprintf("'%s'", command.Hop)
}

func (command *Command) Run() {
	cmd := exec.Command("echo", "hello", "world")
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  cmd.Run()
}