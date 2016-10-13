package main

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

// Main application class
type Rabbit struct {
  Commands Commands
}

// Load builds the commands from the given file
func (rabbit *Rabbit) Load(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil { 
		panic(err)
	}
	err = yaml.Unmarshal(data, &rabbit)
	if err != nil {
		panic(err)
	}
}

// Find determines the best matches for the given args
func (rabbit *Rabbit) Find(args []string) Commands {
	return rabbit.Commands.SortByMatch(args)
}
