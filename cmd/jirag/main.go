package main

import (
	"log"
	"os"

	"github.com/SonarBeserk/jirag/internal/assign"
	"github.com/SonarBeserk/jirag/internal/comments"
	"github.com/SonarBeserk/jirag/internal/config"
	"github.com/SonarBeserk/jirag/internal/issue"
	"github.com/SonarBeserk/jirag/internal/login"
	"github.com/SonarBeserk/jirag/internal/track"
	"github.com/SonarBeserk/jirag/internal/transitions"
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
		{
			Name:    "assign",
			Aliases: []string{"a"},
			Usage:   "Assigns a user to an issue",
			Action:  assign.HandleAssignment,
		},
		{
			Name:    "assign-me",
			Aliases: []string{"am"},
			Usage:   "Assigns currently signed in user to issue",
			Action:  assign.HandleSelfAssignment,
		},
		{
			Name:    "unassign",
			Aliases: []string{"u"},
			Usage:   "Unassigns user from issue",
			Action:  assign.HandleUnassignment,
		},
		{
			Name:    "add-comment",
			Aliases: []string{"ac"},
			Usage:   "Adds a comment to an issue",
			Action:  comments.HandleAddComment,
		},
		{
			Name:    "comments",
			Aliases: []string{"co"},
			Usage:   "Lists comments for an issue",
			Action:  comments.HandleListComments,
		},
		{
			Name:    "move",
			Aliases: []string{"mv"},
			Usage:   "Moves an issue between transitions",
			Action:  transitions.HandleTransitionIssue,
		},
		{
			Name:    "add-time",
			Aliases: []string{"at"},
			Usage:   "Adds time to an issue",
			Action:  track.HandleTrackIssueTime,
		},
		{
			Name:    "worklogs",
			Aliases: []string{"w"},
			Usage:   "Lists worklogs for an issue",
			Action:  track.HandleListIssueWorklogs,
		},
		{
			Name:    "open",
			Aliases: []string{"o"},
			Usage:   "Opens the given issue in the default browser",
			Action:  issue.HandleOpenIssue,
		},
		{
			Name:    "issues",
			Aliases: []string{"i"},
			Usage:   "Lists issues assigned to the current user",
			Action:  issue.HandleAssignedIssues,
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
