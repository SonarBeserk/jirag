package comments

import (
	"strings"

	"github.com/SonarBeserk/jirag/internal/client"
	"github.com/andygrunwald/go-jira"
	"github.com/urfave/cli"
)

// HandleAddComment handles adding issue comments
func HandleAddComment(c *cli.Context) error {
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

	commentSlice := c.Args().Tail()
	if len(commentSlice) == 1 {
		return cli.NewExitError("Comment required", 2)
	}

	usr, _, err := jiraClient.User.GetSelf()
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	jiraClient.Issue.AddComment(key, &jira.Comment{
		Author: *usr,
		Body:   strings.Join(commentSlice, " "),
	})

	return nil
}
