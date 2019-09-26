package config

import (
	"os/user"
	"path"

	"github.com/urfave/cli"
)

// Config represents the configuration flags
type Config struct {
	Host     string `toml:"host"`
	Username string `toml:"username"`
	Token    string `toml:"token"`
}

// GetDefaultConfigPath returns the default path for configs
func GetDefaultConfigPath() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", cli.NewExitError("Unable to find current user: "+err.Error(), 1)
	}

	return path.Join(usr.HomeDir, ".jirag"), nil
}
