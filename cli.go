package main

import (
	"os"
	"fmt"
  "github.com/urfave/cli"
)

/**
 * Main Function
 */
func main() {
	app := cli.NewApp()
	setDetails(app)

  commands := []cli.Command{}
  commands = append(commands, getStaticCommands()...)
  commands = append(commands, getDynamicCommands()...)
  
  app.Commands = commands
  app.Run(os.Args);
}

func setDetails(app *cli.App) {
	app.Name = "Rabbit"
	app.Version = "0.0.0"
  app.Usage = "Command Line Hopper"
}

func getStaticCommands() []cli.Command {
  commands := []cli.Command{}
	commands = append(commands, getInitCommand())
  return commands;
}

func getInitCommand() cli.Command {
  return cli.Command{
    Name: "init",
    Usage: "Create a rabbit.yaml file in the current directory",
    Action: func(c *cli.Context) error {
      fmt.Println("Init run: create file")
      return nil
    },
  }
}

func getDynamicCommands() []cli.Command {
  config := Config{}
  config.Load("rabbit.yaml")
  return buildCommands(config.Commands)
}

func buildCommands(commandList []Command) []cli.Command {
  commands := []cli.Command{}
  for _, data := range commandList {
    commands = append(commands, buildCommand(data))
  }
  return commands
}

func buildCommand(data Command) cli.Command {
  command := cli.Command{}
  command.Name = data.Hop
  command.Usage = data.Description
  command.Subcommands = buildCommands(data.Commands)
  command.Action = func(c *cli.Context) error {
    args := []string {}
    for _, arg := range c.Args() {
      args = append(args, arg)
    }
    data.Run(args)
    return nil
  }
  return command  
}
