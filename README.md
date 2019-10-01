# jirag
Cli tool to interact with jira instances

## Generating Assets

```
go get -u github.com/acstech/go-bindata/...
go-bindata -o internal/data/bindata.go -pkg data tmpl/...
```