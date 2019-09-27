package login

import (
	"bufio"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/SonarBeserk/jirag/internal/client"
	"github.com/SonarBeserk/jirag/internal/config"
	"github.com/urfave/cli"
)

// HandleLogin handles parsing login data and verifying the given login is valid
func HandleLogin(c *cli.Context) error {
	cfg := &config.Config{}

	if !c.Args().Present() {
		promptCfg, err := promptForConfig()
		if err != nil {
			return err
		}

		cfg = promptCfg
	}

	if cfg.Host == "" {
		stdInCfg, err := readFromStdIn(c)
		if err != nil {
			return err
		}

		cfg = stdInCfg
	}

	jiraClient, err := client.NewJiraClient(
		c.GlobalString("host"),
		c.GlobalString("username"),
		c.GlobalString("token"))
	if err != nil {
		return cli.NewExitError("Failed to create client: "+err.Error(), 1)
	}

	usr, _, err := jiraClient.User.GetSelf()
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	fmt.Printf("Authenticated successfully as user: %v" + usr.EmailAddress)
	return saveConfig(cfg)
}

func promptForConfig() (*config.Config, error) {
	cfg := &config.Config{}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Jira hostname: ")
	host, err := reader.ReadString('\n')
	if err != nil {
		return nil, cli.NewExitError(err, 1)
	}

	if host == "" {
		return nil, cli.NewExitError("Jira hostname is blank", 2)
	}

	fmt.Printf("Jira Username: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		return nil, cli.NewExitError(err, 1)
	}

	if username == "" {
		return nil, cli.NewExitError("Jira username is blank", 2)
	}

	fmt.Print("Jira Api Token: ")
	token, err := reader.ReadString('\n')
	if err != nil {
		return nil, cli.NewExitError(err, 1)
	}

	if token == "" {
		return nil, cli.NewExitError("Jira token is blank", 2)
	}

	cfg.Host = host
	cfg.Username = username
	cfg.Token = token
	return cfg, nil
}

func readFromStdIn(c *cli.Context) (*config.Config, error) {
	cfg := &config.Config{}

	cfg.Host = c.Args().Get(0)
	if cfg.Host == "" {
		return nil, cli.NewExitError("Jira hostname is blank", 2)
	}

	cfg.Username = c.Args().Get(1)
	if cfg.Username == "" {
		return nil, cli.NewExitError("Jira username is blank", 2)
	}

	cfg.Token = c.Args().Get(2)
	if cfg.Token == "" {
		return nil, cli.NewExitError("Jira token is blank", 2)
	}

	return cfg, nil
}

func saveConfig(cfg *config.Config) error {
	path, err := config.GetDefaultConfigPath()
	if err != nil {
		return err
	}

	f, err := os.Create(path)
	if err != nil {
		return cli.NewExitError("Unable to create config file: "+err.Error(), 1)
	}

	fEncoder := toml.NewEncoder(f)
	err = fEncoder.Encode(cfg)
	if err != nil {
		return cli.NewExitError("Unable to encode config file: "+err.Error(), 1)
	}

	return nil
}
