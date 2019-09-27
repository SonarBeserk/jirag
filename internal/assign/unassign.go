package assign

import (
	"fmt"

	"github.com/SonarBeserk/jirag/internal/client"
	"github.com/urfave/cli"
)

// HandleUnassignment handles unassigning an issue from any user
func HandleUnassignment(c *cli.Context) error {
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

	_, err = jiraClient.Issue.UpdateAssignee(key, nil)
	if err != nil {
		return cli.NewExitError("Failed to update assignee: "+err.Error(), 1)
	}

	fmt.Printf("Unassigned %v\n", key)
	return nil
}
