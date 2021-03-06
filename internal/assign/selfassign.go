package assign

import (
	"fmt"

	"github.com/SonarBeserk/jirag/internal/client"
	"github.com/urfave/cli"
)

// HandleSelfAssignment handles assigning an issue to the current user
func HandleSelfAssignment(c *cli.Context) error {
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

	assignee, _, err := jiraClient.User.GetSelf()
	if err != nil {
		return cli.NewExitError("Failed to get assignee: "+err.Error(), 1)
	}

	assignee.AccountID = ""

	_, err = jiraClient.Issue.UpdateAssignee(key, assignee)
	if err != nil {
		return cli.NewExitError("Failed to update assignee: "+err.Error(), 1)
	}

	fmt.Printf("Assigned %v to %v\n", key, assignee.Name)
	return nil
}
