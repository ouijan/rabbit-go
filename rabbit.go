package main

import (
	"fmt"
	"os/exec"
	"os"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

/**
 * Check Function
 */
func check(e error) {
  if e != nil {
    panic(e)
  }
}

/**
 * Config Struct
 */
type Config struct {
  commands []string `yaml:",flow"`
}

/**
 * Main Function
 */
func main() {
	fmt.Printf("Rabbit Command Line Hopper\n")

	// Read Yaml
	config := Config{}
	data, err := ioutil.ReadFile("test.yaml")
  check(err)
  err = yaml.Unmarshal([]byte(data), &config)
  check(err)
  fmt.Printf("--- config:\n%v\n\n", config)
	
	// Handle Reserved Commands
	
	// Build Commands
	
	// Determine Command
	
	// Run Command
  cmd := exec.Command("echo", "hello world")
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  cmd.Run()
}


