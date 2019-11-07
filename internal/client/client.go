package client

import (
	"errors"
	"strings"

	"github.com/andygrunwald/go-jira"
)

// NewJiraClient creates a new instance of a jira client
func NewJiraClient(url string, username string, password string) (*jira.Client, error) {
	if url == "" || username == "" || password == "" {
		return nil, errors.New("Invalid credentials detected, please login to call this command")
	}

	transport := jira.BasicAuthTransport{
		Username: strings.TrimSpace(username),
		Password: strings.TrimSpace(password),
	}
	return jira.NewClient(transport.Client(), strings.TrimSpace(url))
}
