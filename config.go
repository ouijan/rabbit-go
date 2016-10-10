package rabbit

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type Config struct {
  Commands []Command
}

func (config *Config) Load() {
	data, err := ioutil.ReadFile("rabbit.yaml")
	if err != nil { 
		panic(err)
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}
}

func (config *Config) Help() {
	for _, command := range config.Commands {
		command.Help()
	}
}