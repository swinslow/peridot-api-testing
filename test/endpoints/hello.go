// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package endpoints

import (
	"github.com/swinslow/peridot-api-testing/internal/testresult"
	"github.com/swinslow/peridot-api-testing/test/utils"
)

func getHelloTests() []testresult.TestFunc {
	return []testresult.TestFunc{
		helloGet,
	}
}

func helloGet(root string) *testresult.TestResult {
	res := &testresult.TestResult{
		Suite:   "endpoints",
		Element: "hello",
		ID:      "GET",
	}

	res.Wanted = `{"message": "hello"}`
	url := root + "/hello"
	err := utils.GetContent(res, "1", url, 200, "none")
	if err != nil {
		return res
	}

	if !utils.IsMatch(res) {
		utils.FailMatch(res, "2")
		return res
	}

	utils.Pass(res)
	return res
}
