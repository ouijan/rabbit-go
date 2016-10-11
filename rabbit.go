package main

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type Rabbit struct {
  Commands []Command
}

func (rabbit *Rabbit) Load() {
	data, err := ioutil.ReadFile("rabbit.yaml")
	if err != nil { 
		panic(err)
	}
	err = yaml.Unmarshal(data, &rabbit)
	if err != nil {
		panic(err)
	}
}

func (rabbit *Rabbit) Find(args []string) []Command {
	var matches []Command

	for _, command := range rabbit.Commands {
		if (command.Matches(args)) {
			matches = append(matches, command)
		}
	}

	return matches
}

