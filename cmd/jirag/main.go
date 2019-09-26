package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "JiraG"
	app.Usage = "Tool for working with Jira"
	app.Version = "1.0.0"
	app.Author = "SonarBeserk"
	app.Action = func(c *cli.Context) error {
		cli.ShowAppHelp(c)
		return cli.NewExitError("No commands provided", 2)
	}

	flags := []cli.Flag{}

	commands := []cli.Command{}

	app.Flags = flags
	app.Commands = commands

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
