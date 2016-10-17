package main

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

/**
 * Main application class
 */
type Config struct {
  Commands []Command
}

/**
 * Load builds the commands from the given file
 */
func (rabbit *Config) Load(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil { 
		return err
	}
	err = yaml.Unmarshal(data, &rabbit)	
	return err
}

