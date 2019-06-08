// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/swinslow/obsidian-api-testing/internal/testresult"
	"github.com/swinslow/obsidian-api-testing/test/endpoints"
)

func main() {
	root := "http://sut:3005"
	anyFailed := false

	allRs := []*testresult.TestResult{}
	var rs []*testresult.TestResult

	// run all test suites
	rs = endpoints.RunTests(root)
	allRs = append(allRs, rs...)

	// set up tabwriter for outputting test result table
	w := tabwriter.NewWriter(os.Stdout, 8, 4, 1, ' ', 0)

	// output results
	for _, r := range allRs {
		var result string
		if r.Success {
			result = "ok"
		} else {
			result = "FAIL"
			anyFailed = true
		}

		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", r.Suite, r.Element, r.ID, result)
	}
	w.Flush()

	if anyFailed {
		// print details of failing tests
		fmt.Printf("\n\n==========\n\n")
		for _, r := range allRs {
			if !r.Success {
				fmt.Printf("%s:%s:%s\n", r.Suite, r.Element, r.ID)
				fmt.Printf("    Status: FAIL\n")
				fmt.Printf("    Step:   %s\n", r.FailStep)
				fmt.Printf("    Errors: %v\n", r.FailError)
				fmt.Printf("    Wanted: %s\n", r.Wanted)
				fmt.Printf("    Got:    %s\n", r.Got)
				fmt.Printf("\n==========\n\n")
			}
		}

		// return failure status code
		os.Exit(1)
	}
}
