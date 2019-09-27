package client

import (
	"strings"

	"github.com/andygrunwald/go-jira"
)

// NewJiraClient creates a new instance of a jira client
func NewJiraClient(url string, username string, password string) (*jira.Client, error) {
	transport := jira.BasicAuthTransport{
		Username: strings.TrimSpace(username),
		Password: strings.TrimSpace(password),
	}
	return jira.NewClient(transport.Client(), strings.TrimSpace(url))
}
