package main

import (
	"fmt"
	// "os/exec"
	// "os"
)

/**
 * Main Function
 */
func rabbit() {
	fmt.Printf("Rabbit Command Line Hopper\n")

	// Read Yaml
	config := Config{}
  config.Load()


  // Build Commands
  fmt.Printf("%s \n", config.Commands)
  // for index, command := range config.Commands {
  	// fmt.Printf("%i: %s", index, command)
	// }
	
	// Handle Reserved Commands
	
	// Build Commands
	
	// Determine Command
	
	// Run Command
  // cmd := exec.Command("echo", "hello world")
  // cmd.Stdout = os.Stdout
  // cmd.Stderr = os.Stderr
  // cmd.Run()
}


