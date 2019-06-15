// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/swinslow/peridot-api-testing/internal/testresult"
	"github.com/swinslow/peridot-api-testing/test/endpoints"
	"github.com/swinslow/peridot-api-testing/test/utils"
)

func resetDB(root string) error {
	resetCommand := `{"command": "resetDB"}`
	client := &http.Client{}
	req, err := http.NewRequest("POST", root+"/admin/db", strings.NewReader(resetCommand))
	if err != nil {
		return fmt.Errorf("got error from resetDB http request creator: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	utils.AddAuthHeader(nil, "", req, "admin")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	bodystr := string(b)

	if bodystr != `{"success": true}` {
		return fmt.Errorf("got error from resetDB command: %s", bodystr)
	}

	return nil
}

func main() {
	root := "http://sut:3005"
	anyFailed := false

	allRs := []*testresult.TestResult{}
	var rs *testresult.TestResult

	// get all test suites
	allTests := endpoints.GetTests()

	// and run them, resetting DB each time
	for _, t := range allTests {
		err := resetDB(root)
		if err != nil {
			fmt.Printf("Error resetting DB before test: %v", err)
			os.Exit(1)
		}
		rs = t(root)
		allRs = append(allRs, rs)
	}

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
