package comments

import (
	"fmt"

	"github.com/SonarBeserk/jirag/internal/client"
	"github.com/urfave/cli"
)

// HandleListComments handles listing comments for an issue
func HandleListComments(c *cli.Context) error {
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

	issue, _, err := jiraClient.Issue.Get(key, nil)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	comments := issue.Fields.Comments
	if comments != nil {
		for _, comment := range comments.Comments {
			fmt.Printf("[%v] %v: %v\n", comment.Created, comment.Author.DisplayName, comment.Body)
		}
	}

	return nil
}
