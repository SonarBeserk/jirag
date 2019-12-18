# jirag

Cli tool to interact with jira instances

## Current Features

* Authentication with Jira and saving of credentials

* Listing the details of an issue

* Assignments to current user or other users

* Unassignment from issues

* Adding comments to issues

* Listing comments of an issue

* Moving issues between different stages

* Adding worklogs to issues

* Listing the worklogs of an issue

## Installing

```bash
go install github.com/SonarBeserk/jirag/cmd/jirag
```

## Generating Assets

```bash
go get -u github.com/acstech/go-bindata/...
go-bindata -o internal/data/bindata.go -pkg data tmpl/...
```
