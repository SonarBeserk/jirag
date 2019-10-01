package issue

import (
	"fmt"
	"html/template"
	"os"

	"github.com/SonarBeserk/jirag/internal/client"
	"github.com/SonarBeserk/jirag/internal/data"
	"github.com/urfave/cli"
)

var (
	currentUserJql = "assignee = currentUser()"
)

func init() {
	issueAsset := data.MustAsset("tmpl/issue.tmpl")
	issueDetailsTmpl = template.Must(template.New("issue").Parse(string(issueAsset)))
}

// HandleAssignedIssues handles listing issues assigned to the logged in account
func HandleAssignedIssues(c *cli.Context) error {
	jiraClient, err := client.NewJiraClient(
		c.GlobalString("host"),
		c.GlobalString("username"),
		c.GlobalString("token"))
	if err != nil {
		return cli.NewExitError("Failed to create client: "+err.Error(), 1)
	}

	issues, _, err := jiraClient.Issue.Search(currentUserJql, nil)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	for _, issue := range issues {
		err = issueDetailsTmpl.Execute(os.Stdout, issue)
		if err != nil {
			return cli.NewExitError(err.Error(), 1)
		}

		fmt.Println("")
	}

	return nil
}
