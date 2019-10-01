package transitions

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/SonarBeserk/jirag/internal/client"
	"github.com/urfave/cli"
)

// HandleTransitionIssue handles moving/transitioning an issue between stages
func HandleTransitionIssue(c *cli.Context) error {
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

	trans, _, err := jiraClient.Issue.GetTransitions(key)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	transitionNames := []string{}

	for _, transition := range trans {
		transitionNames = append(transitionNames, transition.Name)
	}

	transitionName := ""

	prompt := &survey.Select{
		Message: "Please select stage:",
		Options: transitionNames,
	}
	survey.AskOne(prompt, &transitionName)

	transitionID := ""

	for _, transition := range trans {
		if transition.Name == transitionName {
			transitionID = transition.ID
		}
	}

	_, err = jiraClient.Issue.DoTransition(key, transitionID)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	return nil
}
