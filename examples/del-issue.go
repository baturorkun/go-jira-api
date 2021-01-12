package main

import (
	"fmt"
	"github.com/baturorkun/go-jira-api/jira"
)

func main() {
	issue := jira.Issue {
		Key: "TEST-13",
	}
	i := jira.New(issue)

	err := i.DelIssue()

	fmt.Printf("%+v \n", err)

	fmt.Printf("%+v \n", i)
}