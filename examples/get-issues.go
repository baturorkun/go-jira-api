package main

import (
	"fmt"
	"github.com/baturorkun/go-jira-api/jira"
)

func main() {
	issue := jira.Issue{ Project: "TEST" }

	i := jira.New(issue)

	issues := i.GetIssues()

	fmt.Printf("%+v \n", issues)

}