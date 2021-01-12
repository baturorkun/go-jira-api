package main

import (
	"fmt"
	"github.com/baturorkun/go-jira-api/jira"
)

func main() {
	issue := jira.Issue {
		Project: "TEST",
		Labels: []string{"batur", "orkun"},
		Type: "Task",
		Summary: "Test 1",
		Description: "Test 12345",
	}

	i := jira.New(issue)

	err := i.AddIssue()

	fmt.Printf("%+v \n", err)

	fmt.Printf("%+v \n", i)

}