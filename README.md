# go-jira-api #

JIRA REST API Client for Go - Golang

What can you do?

* Adding issue to a Jira board.
    * Assign the user automatically according to "summary" value.
* Deleting issue from a Jira board.
* Gel all issues from a Jira board.
* Find an issue card from a Jira board by the Issue Key.
* Find an issue card by Issue Summary value.
* Update the labels of issue
* Add a new label to the issue.
* Add a new comment to the issue.
* Update the description of an issue.
* Add content to the description of an issue.


## Usage ##

```go
import "github.com/baturorkun/go-jira-api"
```

Set JIRA URL, Username and Password using environment variables

```bash
export JIRA_URL="https://yourcompany.atlassian.net"
export JIRA_USER="username"
export JIRA_PASSWORD="password"
```
Alternatively, you can use the ".env" file.  Rename".env.dist" file to ".env" and change values, then you can use a bash file like "run.sh".

```bash
sh run.sh <example file>
```
## Example Usages ##

Example of Adding Issue.

```go
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
		Summary: "Test Summary",
		Description: "Test Desc 12345",
	}

	i := jira.New(issue)

	err := i.AddIssue()

	fmt.Printf("%+v \n", err)

	fmt.Printf("%+v \n", i)
}
```

```bash
sh run.sh examples/add-issue.go
```

Example of Deleting Issue.

```go
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
```

```bash
sh run.sh examples/add-issue.go
```

**You can find more examples in "examples" directory.**

## Documentation ##

The generated documentation at [GoDoc](http://godoc.org/github.com/baturorkun/go-jira-api/jira).

## License ##

`MIT`, see the `LICENSE` file.