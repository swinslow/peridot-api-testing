package main

import (
	"fmt"
	"os"

	"github.com/swinslow/obsidian-api-testing/internal/testresult"
	"github.com/swinslow/obsidian-api-testing/test/endpoints"
)

func main() {
	root := "http://sut"
	anyFailed := false

	allRs := []*testresult.TestResult{}
	var rs []*testresult.TestResult

	// run all test suites
	rs = endpoints.RunTests(root)
	allRs = append(allRs, rs...)

	// output results
	for _, r := range allRs {
		fmt.Printf("%s\t%s\t%s:\t%t\n", r.Suite, r.Element, r.ID, r.Success)
		if !r.Success {
			anyFailed = true
		}
	}

	if anyFailed {
		os.Exit(1)
	}
}
