package main

import (
	"fmt"
)

func main() {
	fmt.Println("***\n")

	fmt.Println("USAGE:\n-----")
	fmt.Printf("sh run <example file>")

	fmt.Println("You can find the files in \"examples\" directory.")

	fmt.Println("EXAMPLE FILES:\n---------------")

	fmt.Println("<add-issue.go> : Adding issue")
	fmt.Println("<del-issue.go> : Deleting issue")

	fmt.Println("EXAMPLE USAGES:\n---------------")

	fmt.Println("sh run.sh add-issue.go")
	fmt.Println("sh run.sh del-issue.go")

	fmt.Println("\n***")
}


