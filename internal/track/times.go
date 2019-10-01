package track

import (
	"fmt"
	"time"

	"github.com/SonarBeserk/jirag/internal/client"
	"github.com/urfave/cli"
)

// HandleListIssueWorklogs handles listing the worklogs for an issue
func HandleListIssueWorklogs(c *cli.Context) error {
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

	issueWorklog, _, err := jiraClient.Issue.GetWorklogs(key)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	for _, record := range issueWorklog.Worklogs {
		t := time.Time(*record.Created)
		pt := t.Format(trackerFormat)

		fmt.Printf("[%v] %v spent %v: %v\n", pt, record.Author.DisplayName, record.TimeSpent, record.Comment)
	}

	return nil
}
