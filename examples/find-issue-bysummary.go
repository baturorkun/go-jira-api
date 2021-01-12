package main

import (
	"fmt"
	"github.com/baturorkun/go-jira-api/jira"
)

func main() {
	issue := jira.Issue {
		Project: "TEST",
		Summary: "Test 1",
	}

	i := jira.New(issue)

	res := i.FindIssueBySummary(nil)

	fmt.Printf("%+v \n", res)

	fmt.Printf("%+v \n", i)

}