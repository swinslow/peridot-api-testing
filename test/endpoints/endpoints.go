// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package endpoints

import (
	"github.com/swinslow/peridot-api-testing/internal/testresult"
)

// GetTests returns all of the endpoints test suites.
func GetTests() []testresult.TestFunc {
	allTests := []testresult.TestFunc{}

	allTests = append(allTests, getHelloTests()...)
	allTests = append(allTests, getLoginTests()...)
	allTests = append(allTests, getUsersTests()...)
	allTests = append(allTests, getProjectsTests()...)
	allTests = append(allTests, getSubprojectsTests()...)
	allTests = append(allTests, getReposTests()...)

	return allTests
}
