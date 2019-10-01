package issue

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/SonarBeserk/jirag/internal/client"
	"github.com/urfave/cli"
)

// HandleOpenIssue handles opening an issue url in a browser
func HandleOpenIssue(c *cli.Context) error {
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

	_, _, err = jiraClient.Issue.Get(key, nil)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	url := jiraClient.GetBaseURL().Scheme + "://" + jiraClient.GetBaseURL().Host
	url = url + "/browse/" + key

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	return nil
}
