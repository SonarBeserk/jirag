package track

import (
	"errors"
	"fmt"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/SonarBeserk/jirag/internal/client"
	"github.com/andygrunwald/go-jira"
	"github.com/urfave/cli"
)

var (
	trackerFormat = "01/02/2006"

	answers = struct {
		Created   string
		Timespent string
		Comment   string
	}{}

	issueTrackQuestions = []*survey.Question{
		{
			Name: "created",
			Prompt: &survey.Input{
				Message: "Date Created:",
				Help:    "This must be a date in the format 01/02/2006",
			},
			Validate: func(val interface{}) error {
				if str, ok := val.(string); !ok || len(str) > 10 {
					_, err := time.Parse(trackerFormat, str)
					if err != nil {
						return errors.New("This must be a date in the format 01/02/2006")
					}
				}
				return nil
			},
		},
		{
			Name: "timespent",
			Prompt: &survey.Input{
				Message: "Time Spent:",
				Help:    "Time spent is in the format 2w 4d 6h 45m",
			},
		},
		{
			Name:   "comment",
			Prompt: &survey.Input{Message: "Comment:"},
		},
	}
)

// HandleTrackIssueTime handles adding time to an issue
func HandleTrackIssueTime(c *cli.Context) error {
	jiraClient, err := client.NewJiraClient(
		c.GlobalString("host"),
		c.GlobalString("username"),
		c.GlobalString("token"))
	if err != nil {
		return cli.NewExitError("Failed to create client: "+err.Error(), 1)
	}

	key := c.Args().First()
	if key == "" {
		return cli.NewExitError("Issue key required", 2)
	}

	usr, _, err := jiraClient.User.GetSelf()
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	err = survey.Ask(issueTrackQuestions, &answers)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	t, err := time.Parse(trackerFormat, answers.Created)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	jt := jira.Time(t)

	worklog := &jira.WorklogRecord{
		Author:    usr,
		Created:   &jt,
		TimeSpent: answers.Timespent,
		Comment:   answers.Comment,
	}

	_, _, err = jiraClient.Issue.AddWorklogRecord(key, worklog)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	fmt.Printf("Added a new worklog to %s for %s on %s with comment %s\n", key, answers.Timespent, answers.Created, answers.Comment)
	return nil
}
