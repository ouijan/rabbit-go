package rabbit

import (
	"fmt"
)

func rabbit(input string) {
	fmt.Printf("Rabbit Command Line Hopper\n\n")

  // load Commands
  config := Config{}
  config.Load()
  config.Help()
	
	// Handle Reserved Commands
	// Build Commands
	// Determine Command
	// Run Command
}