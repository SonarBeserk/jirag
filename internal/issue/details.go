package issue

import (
	"os"

	"html/template"

	"github.com/SonarBeserk/jirag/internal/client"
	"github.com/urfave/cli"
)

var (
	issueDetailsTmpl *template.Template
)

func init() {
	issueDetailsTmpl = template.Must(template.ParseFiles("tmpl/issue.tmpl"))
}

// HandleIssueDetails handles listing issue details
func HandleIssueDetails(c *cli.Context) error {
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

	err = issueDetailsTmpl.Execute(os.Stdout, issue)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	return nil
}
