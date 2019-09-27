package main

import (
	"log"
	"os"

	"github.com/SonarBeserk/jirag/internal/config"
	"github.com/SonarBeserk/jirag/internal/issue"
	"github.com/SonarBeserk/jirag/internal/login"
	"github.com/urfave/cli"
	"github.com/urfave/cli/altsrc"
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

	flags := []cli.Flag{
		cli.StringFlag{Name: "config"},
		altsrc.NewStringFlag(cli.StringFlag{Name: "host"}),
		altsrc.NewStringFlag(cli.StringFlag{Name: "username"}),
		altsrc.NewStringFlag(cli.StringFlag{Name: "token"}),
	}

	commands := []cli.Command{
		{
			Name:    "login",
			Aliases: []string{"l"},
			Usage:   "Authenticate with jira",
			Action:  login.HandleLogin,
		},
		{
			Name:    "details",
			Aliases: []string{"d"},
			Usage:   "Lists the details of an issue",
			Action:  issue.HandleIssueDetails,
		},
	}

	app.Flags = flags
	app.Commands = commands
	app.Before = loadConfig

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func loadConfig(c *cli.Context) error {
	cfg := c.String("config")
	if cfg == "" {
		path, err := config.GetDefaultConfigPath()
		if err != nil {
			return err
		}

		cfg = path
	}

	if _, err := os.Stat(cfg); os.IsNotExist(err) {
		return nil
	}

	inputSource, err := altsrc.NewTomlSourceFromFile(cfg)
	if err != nil {
		return cli.NewExitError("Unable to load config: "+err.Error(), 1)
	}

	return altsrc.ApplyInputSourceValues(c, inputSource, c.App.Flags)
}
