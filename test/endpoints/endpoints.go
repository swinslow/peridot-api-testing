package endpoints

import (
	"github.com/swinslow/obsidian-api-testing/internal/testresult"
)

// RunTests runs all of the endpoints test suites, and accumulates
// the test results.
func RunTests(root string) []*testresult.TestResult {
	allRs := []*testresult.TestResult{}
	var rs []*testresult.TestResult

	rs = runHelloTests(root)
	allRs = append(allRs, rs...)

	return allRs
}